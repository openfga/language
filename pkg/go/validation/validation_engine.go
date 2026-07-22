package validation

import (
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
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

// ValidateDSL validates a DSL model with all available validations.
func ValidateDSL(model *openfgav1.AuthorizationModel, dslContent string, options *EngineOptions) *ValidationErrors {
	if options == nil {
		options = DefaultEngineOptions()
	}
	return NewValidationEngine(model, dslContent).RunAllValidations(options)
}

// ValidateJSON validates a JSON model.
func ValidateJSON(model *openfgav1.AuthorizationModel, options *EngineOptions) *ValidationErrors {
	if options == nil {
		options = DefaultEngineOptions()
	}
	return NewValidationEngine(model, "").RunAllValidations(options)
}

// RunAllValidations executes all validation phases in the correct order.
func (ve *ValidationEngine) RunAllValidations(options *EngineOptions) *ValidationErrors {
	if ve.model == nil {
		return NewValidationErrors(nil)
	}

	// Schema and name validation run first and unconditionally.
	ve.runSchemaValidation()
	ve.runNameValidation()

	// Relation-reference validation always runs. The phases that follow are
	// gated on there being no errors yet: a model with bad references or
	// duplicates would otherwise produce a cascade of derived entry-point and
	// complex-operation errors for the same root cause. This mirrors the
	// reference implementation's modelValidation, which skips the later passes
	// once any error has been recorded.
	if !options.SkipSemanticValidation {
		validateRelationReferences(ve.collector, ve.semantic, ve.lines)
	}

	if !ve.collector.HasErrors() {
		ve.runDuplicateDetection()
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
		ve.runMultiFileValidation()
	}
	if !options.SkipConditionValidation {
		ve.runConditionValidation()
	}

	return NewValidationErrors(ve.collector.GetErrors())
}

func (ve *ValidationEngine) runSchemaValidation() {
	ValidateSchemaVersion(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runNameValidation() {
	ValidateNames(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runDuplicateDetection() {
	ValidateDuplicates(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runMultiFileValidation() {
	ValidateMultiFileConsistency(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runConditionValidation() {
	validateConditionReferences(ve.collector, ve.condition, ve.lines)
	ValidateConditionConsistency(ve.collector, ve.model, ve.lines)
	validateUnusedConditions(ve.collector, ve.condition, ve.lines)
}

// ValidateModel is a convenience function that validates a model with default options.
func ValidateModel(model *openfgav1.AuthorizationModel, dslContent string) *ValidationErrors {
	return ValidateDSL(model, dslContent, DefaultEngineOptions())
}

// ValidateModelJSON is a convenience function that validates a JSON model with default options.
func ValidateModelJSON(model *openfgav1.AuthorizationModel) *ValidationErrors {
	return ValidateJSON(model, DefaultEngineOptions())
}

func (ve *ValidationEngine) GetValidationSummary() ValidationSummary {
	errors := ve.collector.GetErrors()
	summary := ValidationSummary{
		TotalErrors:       len(errors),
		ErrorsByType:      make(map[ValidationErrorType]int),
		ErrorsByFile:      make(map[string]int),
		HasCriticalErrors: false,
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
		if ve.isCriticalError(err.Metadata.ErrorType) {
			summary.HasCriticalErrors = true
		}
	}
	return summary
}

// ValidationSummary provides a high-level overview of validation results.
type ValidationSummary struct {
	TotalErrors       int
	ErrorsByType      map[ValidationErrorType]int
	ErrorsByFile      map[string]int
	HasCriticalErrors bool
}

// criticalErrorTypes is the fixed set of error types considered critical. It is
// a package-level lookup table so it isn't rebuilt on every isCriticalError call
// (which runs once per error while summarizing).
var criticalErrorTypes = map[ValidationErrorType]bool{
	RelationNoEntrypoint:  true,
	CyclicRelation:        true,
	UndefinedType:         true,
	UndefinedRelation:     true,
	InvalidRelationType:   true,
	DuplicatedError:       true,
	InvalidSchema:         true,
	InvalidSchemaVersion:  true,
	MultipleModulesInFile: true,
}

func (ve *ValidationEngine) isCriticalError(errorType ValidationErrorType) bool {
	return criticalErrorTypes[errorType]
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

func (vr *ValidationReport) IsValid() bool           { return vr.ValidationErrors.Count() == 0 }
func (vr *ValidationReport) HasCriticalErrors() bool { return vr.Summary.HasCriticalErrors }
func (vr *ValidationReport) GetErrorsByType(errorType ValidationErrorType) []*ValidationError {
	var matchingErrors []*ValidationError
	for _, err := range vr.ValidationErrors.GetErrors() {
		if err.Metadata.ErrorType == errorType {
			matchingErrors = append(matchingErrors, err)
		}
	}
	return matchingErrors
}
