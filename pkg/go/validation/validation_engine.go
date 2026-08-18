package validation

import (
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// ValidationEngine is the main entry point for all validation operations.
type ValidationEngine struct {
	model     *openfgav1.AuthorizationModel
	lines     []string
	collector *ErrorCollector
	// semantic and condition index the model once and are shared across every
	// phase that needs them, rather than each phase rebuilding its own.
	semantic  *SemanticValidator
	condition *ConditionValidator
}

// EngineOptions configures validation behavior.
type EngineOptions struct {
	SkipSemanticValidation         bool
	SkipComplexOperationValidation bool
	SkipWildcardValidation         bool
	SkipMultiFileValidation        bool
	SkipConditionValidation        bool
}

func DefaultEngineOptions() *EngineOptions {
	return &EngineOptions{}
}

func NewValidationEngine(model *openfgav1.AuthorizationModel, dslContent string) *ValidationEngine {
	lines := strings.Split(dslContent, "\n")
	collector := NewErrorCollector(lines)
	ve := &ValidationEngine{model: model, lines: lines, collector: collector}
	if model != nil {
		ve.semantic = NewSemanticValidator(model)
		ve.condition = NewConditionValidator(model)
	}
	return ve
}

// ValidateDSL runs every validation over model, using dslContent to resolve each
// finding's position in the source text. The model is the already-parsed proto,
// here and in ValidateJSON; neither parses anything.
//
// Returns nil for a valid model. Otherwise the error is a *ValidationErrors, which
// errors.As recovers to list every finding; see ValidationErrors.ErrorOrNil for why
// a model carrying only warnings is nil here, and CreateValidationReport for reaching
// those findings.
func ValidateDSL(model *openfgav1.AuthorizationModel, dslContent string, options *EngineOptions) error {
	if options == nil {
		options = DefaultEngineOptions()
	}
	return NewValidationEngine(model, dslContent).RunAllValidations(options).ErrorOrNil()
}

// ValidateJSON runs every validation over a model that reached the caller as JSON,
// so without the DSL source text behind it. It takes the same parsed proto as
// ValidateDSL and decodes no JSON itself; the name matches pkg/js's validateJSON and
// pkg/java's ModelValidator.validateJson.
//
// With no source text to resolve positions against, findings carry a nil Line and
// Column. The messages, categories and metadata are what ValidateDSL reports for the
// same model. Returns nil for a valid model, as ValidateDSL does.
func ValidateJSON(model *openfgav1.AuthorizationModel, options *EngineOptions) error {
	if options == nil {
		options = DefaultEngineOptions()
	}
	return NewValidationEngine(model, "").RunAllValidations(options).ErrorOrNil()
}

// RunAllValidations executes all validation phases in the correct order.
func (ve *ValidationEngine) RunAllValidations(options *EngineOptions) *ValidationErrors {
	if ve.model == nil {
		return NewValidationErrors(nil)
	}

	// Schema and name validation run first and unconditionally.
	ValidateSchemaVersion(ve.collector, ve.model, ve.lines)
	ValidateNames(ve.collector, ve.model, ve.lines)

	// Relation-reference validation always runs. The phases that follow are
	// gated on there being no blocking error yet: a model with bad references or
	// duplicates would otherwise produce a cascade of derived entry-point and
	// complex-operation errors for the same root cause. This mirrors the
	// reference implementation's modelValidation, which skips the later passes
	// once any error has been recorded.
	//
	// The gate counts blocking findings only, so a warning or advisory does not stop
	// the later passes from finding an error that would invalidate the model.
	if !options.SkipSemanticValidation {
		validateRelationReferences(ve.collector, ve.semantic, ve.lines)
	}

	if !ve.collector.HasErrors() {
		ValidateDuplicates(ve.collector, ve.model, ve.lines)
	}

	if !ve.collector.HasErrors() {
		if !options.SkipSemanticValidation {
			validateCyclesAndEntryPoints(ve.collector, ve.semantic, ve.lines)
			validateTupleToUsersetRequirements(ve.collector, ve.semantic, ve.lines)
		}
		if !options.SkipComplexOperationValidation {
			validateComplexOperations(ve.collector, ve.semantic, ve.lines)
		}
		if !options.SkipWildcardValidation {
			validateWildcardUsage(ve.collector, ve.semantic, ve.lines)
		}
	}

	// Multi-file and condition checks are independent of the cascade and always
	// run, matching the reference's handling of conditions.
	if !options.SkipMultiFileValidation {
		ValidateMultiFileConsistency(ve.collector, ve.model, ve.lines)
	}
	if !options.SkipConditionValidation {
		validateConditionReferences(ve.collector, ve.condition, ve.lines)
		ValidateConditionConsistency(ve.collector, ve.model, ve.lines)
		validateUnusedConditions(ve.collector, ve.condition, ve.lines)
	}

	return NewValidationErrors(ve.collector.AllFindings())
}

// ValidateModel is ValidateDSL with the default options, which skip no phase.
func ValidateModel(model *openfgav1.AuthorizationModel, dslContent string) error {
	return ValidateDSL(model, dslContent, DefaultEngineOptions())
}

// ValidateModelJSON is ValidateJSON with the default options.
func ValidateModelJSON(model *openfgav1.AuthorizationModel) error {
	return ValidateJSON(model, DefaultEngineOptions())
}

func (ve *ValidationEngine) GetValidationSummary() ValidationSummary {
	errors := ve.collector.AllFindings()
	summary := ValidationSummary{
		TotalErrors:        ve.collector.Count(),
		TotalFindings:      ve.collector.CountAll(),
		ErrorsByType:       make(map[ValidationErrorType]int),
		ErrorsByFile:       make(map[string]int),
		FindingsBySeverity: make(map[fgaerrors.Severity]int),
		HasCriticalErrors:  false,
	}
	for _, err := range errors {
		if err == nil || err.Metadata == nil {
			// Metadata is always set by the collector, but a directly-constructed
			// error (e.g. in a consumer or test) could omit it; don't panic.
			continue
		}
		summary.ErrorsByType[err.Metadata.ErrorType]++
		if err.File != "" {
			summary.ErrorsByFile[err.File]++
		}
		summary.FindingsBySeverity[err.Severity]++
		if isCriticalErrorType(err.Metadata.ErrorType) {
			summary.HasCriticalErrors = true
		}
	}
	return summary
}

// ValidationSummary provides a high-level overview of validation results.
//
// The breakdowns cover every finding, so they sum to TotalFindings, not TotalErrors.
type ValidationSummary struct {
	// TotalErrors counts only the findings that make the model invalid.
	TotalErrors int

	// TotalFindings counts everything reported, including warnings and advisories.
	TotalFindings int

	ErrorsByType map[ValidationErrorType]int
	ErrorsByFile map[string]int

	// FindingsBySeverity counts findings by severity. A finding with no severity set
	// counts under SeverityUnspecified.
	FindingsBySeverity map[fgaerrors.Severity]int

	HasCriticalErrors bool
}

// CreateValidationReport creates a detailed validation report.
func CreateValidationReport(model *openfgav1.AuthorizationModel, dslContent string, options *EngineOptions) ValidationReport {
	engine := NewValidationEngine(model, dslContent)
	validationErrors := engine.RunAllValidations(options)
	summary := engine.GetValidationSummary()
	return ValidationReport{
		Model:            model,
		ValidationErrors: validationErrors,
		Summary:          summary,
		Options:          options,
	}
}

// ValidationReport contains comprehensive validation results.
type ValidationReport struct {
	Model            *openfgav1.AuthorizationModel
	ValidationErrors *ValidationErrors
	Summary          ValidationSummary
	Options          *EngineOptions
}

// IsValid reports whether the model is usable: no finding blocks it. Warnings and
// advisories leave it valid; HasFindings reports whether any were raised.
func (vr *ValidationReport) IsValid() bool           { return !vr.ValidationErrors.HasErrors() }
func (vr *ValidationReport) HasCriticalErrors() bool { return vr.Summary.HasCriticalErrors }

// GetErrorsByType returns findings of a given error type, blocking or not: the
// caller has named the code it wants, so filtering by severity as well would drop
// matches it asked for.
func (vr *ValidationReport) GetErrorsByType(errorType ValidationErrorType) []*ValidationError {
	var matchingErrors []*ValidationError
	for _, err := range vr.ValidationErrors.AllFindings() {
		// The collector always sets metadata, but a directly-constructed finding
		// need not have, and a code is only readable off metadata.
		if err == nil || err.Metadata == nil {
			continue
		}

		if err.Metadata.ErrorType == errorType {
			matchingErrors = append(matchingErrors, err)
		}
	}
	return matchingErrors
}
