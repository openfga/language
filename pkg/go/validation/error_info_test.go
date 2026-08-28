package validation

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// unemittedErrorTypes are declared ValidationErrorType values that no validation
// produces. The other side is read out of the source by emittedErrorTypes below, and
// TestEveryErrorTypeIsClassified requires every declared code to be in one or the
// other.
//
// They are kept rather than deleted because each has a published documentation
// page, and because SelfError and InvalidSyntax are equally unemitted in
// pkg/js/errors.ts. A cycle with no entrypoint surfaces as RelationNoEntrypoint,
// leaving CyclicError and CyclicRelation nothing to report. InvalidSchemaVersion is
// unreachable because newInvalidSchemaVersionError emits InvalidSchema, which is what
// the shared corpus expects.
//
// None get an errorInfoByType entry, so lookupErrorInfo treats them as blocking
// with no cause. Anything that starts emitting one must add it to that table in the
// same change.
var unemittedErrorTypes = map[ValidationErrorType]struct{}{
	SelfError:            {},
	InvalidSyntax:        {},
	CyclicError:          {},
	CyclicRelation:       {},
	InvalidSchemaVersion: {},
}

// allErrorTypes lists every declared ValidationErrorType. A Go const block of a
// string type cannot be enumerated at runtime, so exhaustiveness checks need it
// written out.
//
// Keep in sync with the const block in errors.go. TestAllErrorTypesIsComplete reads
// that block and fails if the two disagree.
var allErrorTypes = []ValidationErrorType{
	SchemaVersionRequired,
	SchemaVersionUnsupported,
	ReservedTypeKeywords,
	ReservedRelationKeywords,
	SelfError,
	InvalidName,
	MissingDefinition,
	InvalidRelationType,
	InvalidRelationOnTupleset,
	InvalidType,
	RelationNoEntrypoint,
	TuplesetNotDirect,
	DuplicatedError,
	UndefinedType,
	UndefinedRelation,
	CyclicError,
	InvalidWildcardError,
	AssignableRelationsMustHaveType,
	InvalidSchema,
	InvalidSyntax,
	TypeRestrictionCannotHaveWildcardAndRelation,
	ConditionNotDefined,
	ConditionNotUsed,
	DifferentNestedConditionName,
	MultipleModulesInFile,
	CyclicRelation,
	InvalidSchemaVersion,
}

// emittedErrorTypes parses this package's non-test sources and returns the name of
// every ValidationErrorType passed as the errorType argument of a newValidationError
// call. It reads the source rather than a hand-written list, which would go stale in
// the same edit that leaves a code out of the table.
func emittedErrorTypes(t *testing.T) map[string]string {
	t.Helper()

	entries, err := os.ReadDir(".")
	require.NoError(t, err)

	emitted := make(map[string]string)
	fileSet := token.NewFileSet()

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(fileSet, name, nil, parser.SkipObjectResolution)
		require.NoError(t, err, "parsing %s", name)

		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Every constructor names the finding's code as the second argument to
			// newValidationError(message, <ErrorType>, ...). That is the one place a
			// code reaches a finding, so reading it here reports exactly the set a
			// raise site can produce.
			identFun, ok := call.Fun.(*ast.Ident)
			if !ok || identFun.Name != "newValidationError" || len(call.Args) < 2 {
				return true
			}

			identifier, ok := call.Args[1].(*ast.Ident)
			if !ok {
				// A non-identifier errorType means the emitted set can't be
				// determined statically, and this test would silently under-report.
				t.Errorf("%s: newValidationError called with a non-constant errorType at %s; "+
					"emittedErrorTypes can no longer see what this emits",
					name, fileSet.Position(call.Args[1].Pos()))

				return true
			}

			emitted[identifier.Name] = fileSet.Position(call.Pos()).String()

			return true
		})
	}

	return emitted
}

// TestErrorInfoCoversEveryEmittedErrorType checks every code a constructor can
// emit has a table entry, so any finding that reaches a caller has a cause to
// match. It fails when a new constructor is added without one.
func TestErrorInfoCoversEveryEmittedErrorType(t *testing.T) {
	t.Parallel()

	emitted := emittedErrorTypes(t)
	require.NotEmpty(t, emitted, "found no newValidationError calls — the AST walk is broken, not the errorInfoByType")

	// Names, because the AST gives us identifiers and the table is keyed by value.
	classifiedNames := make(map[string]struct{}, len(errorInfoByType))
	for errorType := range errorInfoByType {
		classifiedNames[errorTypeConstantName(t, errorType)] = struct{}{}
	}

	for name, position := range emitted {
		if _, ok := classifiedNames[name]; !ok {
			t.Errorf("%s is emitted at %s but has no errorInfoByType entry: "+
				"callers cannot match its cause with errors.Is", name, position)
		}
	}
}

// TestErrorInfoHasNoUnemittedEntries checks the other direction: an entry for a code
// nothing raises hides that the code is dead.
func TestErrorInfoHasNoUnemittedEntries(t *testing.T) {
	t.Parallel()

	emitted := emittedErrorTypes(t)

	for errorType := range errorInfoByType {
		name := errorTypeConstantName(t, errorType)
		if _, ok := emitted[name]; !ok {
			t.Errorf("errorInfoByType has an entry for %s (%q) but nothing emits it — "+
				"either wire up the constructor or move it to unemittedErrorTypes",
				name, errorType)
		}
	}
}

// TestEveryErrorTypeIsClassified checks a newly declared error type cannot sit in
// neither map. Adding a constant and forgetting the table passes every other test
// in this file.
func TestEveryErrorTypeIsClassified(t *testing.T) {
	t.Parallel()

	for _, errorType := range allErrorTypes {
		_, inErrorInfo := errorInfoByType[errorType]
		_, unemitted := unemittedErrorTypes[errorType]

		assert.Truef(t, inErrorInfo || unemitted,
			"%q appears in neither errorInfoByType nor unemittedErrorTypes; "+
				"classify it as one or the other", errorType)
		assert.Falsef(t, inErrorInfo && unemitted,
			"%q is in both errorInfoByType and unemittedErrorTypes", errorType)
	}
}

// TestAllErrorTypesIsComplete checks the hand-written allErrorTypes list the two
// tests above depend on, by reading the const block it mirrors.
func TestAllErrorTypesIsComplete(t *testing.T) {
	t.Parallel()

	declared := declaredErrorTypeValues(t)

	listed := make(map[ValidationErrorType]struct{}, len(allErrorTypes))
	for _, errorType := range allErrorTypes {
		_, duplicate := listed[errorType]
		assert.Falsef(t, duplicate, "%q is listed twice in allErrorTypes", errorType)
		listed[errorType] = struct{}{}
	}

	for name, value := range declared {
		_, ok := listed[value]
		assert.Truef(t, ok, "%s (%q) is declared in errors.go but missing from allErrorTypes", name, value)
	}

	assert.Len(t, allErrorTypes, len(declared),
		"allErrorTypes has %d entries but %d ValidationErrorType constants are declared",
		len(allErrorTypes), len(declared))
}

// TestErrorInfoEntriesAreWellFormed checks each entry says something usable: a
// severity that exists and a non-nil cause. Which part of the model a code is about
// is not in the table, so it is checked on the findings themselves, in
// TestEverySemanticFindingCarriesErrorInfo.
func TestErrorInfoEntriesAreWellFormed(t *testing.T) {
	t.Parallel()

	validSeverities := map[fgaerrors.Severity]struct{}{
		fgaerrors.SeverityError:    {},
		fgaerrors.SeverityWarning:  {},
		fgaerrors.SeverityAdvisory: {},
	}

	for errorType, entry := range errorInfoByType {
		t.Run(string(errorType), func(t *testing.T) {
			t.Parallel()

			_, ok := validSeverities[entry.Severity]
			assert.Truef(t, ok, "severity %q is not one of error/warning/advisory", entry.Severity)

			assert.Error(t, entry.Cause, "no cause: errors.Is has nothing to match against")
		})
	}
}

// TestEveryEntryBlocks pins what the validation entry points currently rely on:
// every classified code blocks, so Count equals CountAll and a model with any
// finding at all returns non-nil from ValidateDSL.
//
// This is a tripwire rather than a rule. The first non-blocking entry is a
// deliberate change, and it needs ValidationErrors.ErrorOrNil settled in the same
// edit: a collection holding only warnings answers nil there, so that finding never
// reaches a caller through the entry points at all. Either route it through
// CreateValidationReport or change what the entry points return, then update this
// test.
func TestEveryEntryBlocks(t *testing.T) {
	t.Parallel()

	for errorType, entry := range errorInfoByType {
		assert.Truef(t, entry.Severity.Blocks(),
			"%q is classified %s, the first non-blocking code in the table: a model whose "+
				"only finding is this one is valid as far as ValidateDSL is concerned",
			errorType, entry.Severity)
	}
}

// TestLookupErrorInfoFallsBackToBlocking checks an unclassified finding still fails
// validation. Downgrading it to advisory would let an invalid model through.
func TestLookupErrorInfoFallsBackToBlocking(t *testing.T) {
	t.Parallel()

	entry := lookupErrorInfo(ValidationErrorType("no-such-error-type"))

	assert.Equal(t, fgaerrors.SeverityError, entry.Severity)
	assert.True(t, entry.Severity.Blocks(), "an unknown error type must still block validation")
	assert.NoError(t, entry.Cause, "an unknown error type has no cause to report")
}

// TestEveryErrorTypeHasDocumentation keeps the slugs and docs/validation/model in
// step. The slug is what a user sees, so one with no page is a dead end.
func TestEveryErrorTypeHasDocumentation(t *testing.T) {
	t.Parallel()

	docsDir := filepath.Join("..", "..", "..", "docs", "validation", "model")
	if _, err := os.Stat(docsDir); os.IsNotExist(err) {
		t.Skipf("docs directory not present at %s", docsDir)
	}

	for _, errorType := range allErrorTypes {
		page := filepath.Join(docsDir, string(errorType)+".md")
		_, err := os.Stat(page)
		assert.NoErrorf(t, err, "%q has no documentation page at %s", errorType, page)
	}
}

// errorTypeConstantName maps a slug back to its Go constant name, so failures name
// the identifier to edit rather than the string.
func errorTypeConstantName(t *testing.T, errorType ValidationErrorType) string {
	t.Helper()

	for name, value := range declaredErrorTypeValues(t) {
		if value == errorType {
			return name
		}
	}

	t.Fatalf("%q is not a declared ValidationErrorType constant", errorType)

	return ""
}

// declaredErrorTypeValues parses errors.go and returns every declared
// ValidationErrorType constant as name → value.
func declaredErrorTypeValues(t *testing.T) map[string]ValidationErrorType {
	t.Helper()

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "errors.go", nil, parser.SkipObjectResolution)
	require.NoError(t, err)

	declared := make(map[string]ValidationErrorType)

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			typeIdent, ok := valueSpec.Type.(*ast.Ident)
			if !ok || typeIdent.Name != "ValidationErrorType" {
				continue
			}

			for i, name := range valueSpec.Names {
				require.Lessf(t, i, len(valueSpec.Values), "%s has no value", name.Name)

				literal, ok := valueSpec.Values[i].(*ast.BasicLit)
				require.Truef(t, ok, "%s is not assigned a string literal", name.Name)

				value, err := strconv.Unquote(literal.Value)
				require.NoError(t, err)

				declared[name.Name] = ValidationErrorType(value)
			}
		}
	}

	require.NotEmpty(t, declared, "parsed no ValidationErrorType constants from errors.go")

	return declared
}
