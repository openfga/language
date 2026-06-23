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
	return &ValidationEngine{model: model, lines: lines, collector: collector}
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

	ve.runSchemaValidation()
	ve.runDuplicateDetection()

	if !options.SkipSemanticValidation {
		ve.runSemanticValidation()
	}
	if !options.SkipComplexOperationValidation {
		ve.runComplexOperationValidation()
	}
	if !options.SkipWildcardValidation {
		ve.runWildcardValidation()
	}
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

func (ve *ValidationEngine) runDuplicateDetection() {
	ValidateDuplicates(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runSemanticValidation() {
	ValidateRelationReferences(ve.collector, ve.model, ve.lines)
	ValidateCyclesAndEntryPoints(ve.collector, ve.model, ve.lines)
	ValidateTupleToUsersetRequirements(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runComplexOperationValidation() {
	ValidateComplexOperations(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runWildcardValidation() {
	ValidateWildcardUsage(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runMultiFileValidation() {
	ValidateMultiFileConsistency(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runConditionValidation() {
	ValidateConditionReferences(ve.collector, ve.model, ve.lines)
	ValidateConditionConsistency(ve.collector, ve.model, ve.lines)
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

func (ve *ValidationEngine) isCriticalError(errorType ValidationErrorType) bool {
	criticalErrors := map[ValidationErrorType]bool{
		RelationNoEntrypoint:  true,
		CyclicRelation:        true,
		UndefinedType:         true,
		UndefinedRelation:     true,
		InvalidRelationType:   true,
		DuplicatedError:       true,
		InvalidSchemaVersion:  true,
		MultipleModulesInFile: true,
	}
	return criticalErrors[errorType]
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
