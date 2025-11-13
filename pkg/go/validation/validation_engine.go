package validation

import (
	"strings"

	fgaSdk "github.com/openfga/go-sdk"
)

// ValidationEngine is the main entry point for all validation operations
// It coordinates all validation components and provides unified validation functions
type ValidationEngine struct {
	model     *fgaSdk.AuthorizationModel
	lines     []string
	collector *ErrorCollector
}

// EngineOptions configures validation behavior
type EngineOptions struct {
	// SkipSemanticValidation disables semantic validation (cycles, entry points, etc.)
	SkipSemanticValidation bool

	// SkipComplexOperationValidation disables complex operation validation
	SkipComplexOperationValidation bool

	// SkipWildcardValidation disables wildcard usage validation
	SkipWildcardValidation bool

	// SkipMultiFileValidation disables multi-file and module consistency validation
	SkipMultiFileValidation bool

	// SkipConditionValidation disables condition validation (unused conditions, etc.)
	SkipConditionValidation bool
}

// DefaultEngineOptions returns the default validation options with all validations enabled
func DefaultEngineOptions() *EngineOptions {
	return &EngineOptions{
		SkipSemanticValidation:         false,
		SkipComplexOperationValidation: false,
		SkipWildcardValidation:         false,
		SkipMultiFileValidation:        false,
		SkipConditionValidation:        false,
	}
}

// NewValidationEngine creates a new validation engine for the given model
func NewValidationEngine(model *fgaSdk.AuthorizationModel, dslContent string) *ValidationEngine {
	lines := strings.Split(dslContent, "\n")
	collector := NewErrorCollector(lines)

	return &ValidationEngine{
		model:     model,
		lines:     lines,
		collector: collector,
	}
}

// ValidateDSL validates a DSL model with all available validations
// This is the Go equivalent of the JavaScript validateDSL() function
func ValidateDSL(model *fgaSdk.AuthorizationModel, dslContent string, options *EngineOptions) *ValidationErrors {
	if options == nil {
		options = DefaultEngineOptions()
	}

	engine := NewValidationEngine(model, dslContent)
	return engine.RunAllValidations(options)
}

// ValidateJSON validates a JSON model by performing the same validations as ValidateDSL
// This is the Go equivalent of the JavaScript validateJSON() function
func ValidateJSON(model *fgaSdk.AuthorizationModel, options *EngineOptions) *ValidationErrors {
	if options == nil {
		options = DefaultEngineOptions()
	}

	// For JSON validation, we don't have DSL content, so we use an empty line array
	engine := NewValidationEngine(model, "")
	return engine.RunAllValidations(options)
}

// RunAllValidations executes all validation phases in the correct order
func (ve *ValidationEngine) RunAllValidations(options *EngineOptions) *ValidationErrors {
	if ve.model == nil {
		return NewValidationErrors(nil)
	}

	// Phase 1: Schema and Basic Structure Validation
	ve.runSchemaValidation()
	ve.runBasicStructureValidation()

	// Phase 2: Name and Reserved Keyword Validation
	ve.runNameValidation()
	ve.runDuplicateDetection()

	// Phase 3: Semantic Validation
	if !options.SkipSemanticValidation {
		ve.runSemanticValidation()
	}

	// Phase 4: Advanced Validation Features
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

// Phase 1: Schema and Basic Structure Validation

func (ve *ValidationEngine) runSchemaValidation() {
	ValidateSchemaVersion(ve.collector, ve.model, ve.lines)
}

func (ve *ValidationEngine) runBasicStructureValidation() {
	// Basic structure validation is inherently handled by the SDK parsing
	// Additional custom structure validation could be added here if needed
}

// Phase 2: Name and Reserved Keyword Validation

func (ve *ValidationEngine) runNameValidation() {
	// Reserved keyword and naming validation
	// Note: Individual name validation is performed during model construction
	// Additional validation rules can be added here as needed
}

func (ve *ValidationEngine) runDuplicateDetection() {
	// Duplicate detection - implemented in duplicate_detection.go
	ValidateDuplicates(ve.collector, ve.model, ve.lines)
}

// Phase 3: Semantic Validation

func (ve *ValidationEngine) runSemanticValidation() {
	// Core semantic validation
	ValidateRelationReferences(ve.collector, ve.model, ve.lines)
	ValidateCyclesAndEntryPoints(ve.collector, ve.model, ve.lines)

	// Tuple-to-userset validation
	ValidateTupleToUsersetRequirements(ve.collector, ve.model, ve.lines)
}

// Phase 4: Advanced Validation Features

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
	// Note: Condition validation is currently paused due to SDK structure complexities
	// ValidateUnusedConditions(ve.collector, ve.model, ve.lines)
	ValidateConditionReferences(ve.collector, ve.model, ve.lines)
	ValidateConditionConsistency(ve.collector, ve.model, ve.lines)
}

// ValidateModel is a convenience function that validates a model with default options
func ValidateModel(model *fgaSdk.AuthorizationModel, dslContent string) *ValidationErrors {
	return ValidateDSL(model, dslContent, DefaultEngineOptions())
}

// ValidateModelJSON is a convenience function that validates a JSON model with default options
func ValidateModelJSON(model *fgaSdk.AuthorizationModel) *ValidationErrors {
	return ValidateJSON(model, DefaultEngineOptions())
}

// GetValidationSummary returns a summary of validation results
func (ve *ValidationEngine) GetValidationSummary() ValidationSummary {
	errors := ve.collector.GetErrors()

	summary := ValidationSummary{
		TotalErrors:       len(errors),
		ErrorsByType:      make(map[ValidationErrorType]int),
		ErrorsByFile:      make(map[string]int),
		HasCriticalErrors: false,
	}

	// Categorize errors
	for _, err := range errors {
		summary.ErrorsByType[err.Metadata.ErrorType]++

		if err.File != "" {
			summary.ErrorsByFile[err.File]++
		}

		// Check for critical errors
		if ve.isCriticalError(err.Metadata.ErrorType) {
			summary.HasCriticalErrors = true
		}
	}

	return summary
}

// ValidationSummary provides a high-level overview of validation results
type ValidationSummary struct {
	TotalErrors       int
	ErrorsByType      map[ValidationErrorType]int
	ErrorsByFile      map[string]int
	HasCriticalErrors bool
}

// isCriticalError determines if an error type should be considered critical
func (ve *ValidationEngine) isCriticalError(errorType ValidationErrorType) bool {
	criticalErrors := map[ValidationErrorType]bool{
		// Semantic errors that break model functionality
		RelationNoEntrypoint: true,
		CyclicRelation:       true,
		UndefinedType:        true,
		UndefinedRelation:    true,
		InvalidRelationType:  true,

		// Structure errors that break model validity
		DuplicatedError:      true,
		InvalidSchemaVersion: true,

		// Module consistency errors
		MultipleModulesInFile: true,
	}

	return criticalErrors[errorType]
}

// CreateValidationReport creates a detailed validation report
func CreateValidationReport(model *fgaSdk.AuthorizationModel, dslContent string, options *EngineOptions) ValidationReport {
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

// ValidationReport contains comprehensive validation results
type ValidationReport struct {
	Model            *fgaSdk.AuthorizationModel
	ValidationErrors *ValidationErrors
	Summary          ValidationSummary
	Options          *EngineOptions
}

// IsValid returns true if the model has no validation errors
func (vr *ValidationReport) IsValid() bool {
	return vr.ValidationErrors.Count() == 0
}

// HasCriticalErrors returns true if the model has critical validation errors
func (vr *ValidationReport) HasCriticalErrors() bool {
	return vr.Summary.HasCriticalErrors
}

// GetErrorsByType returns all errors of a specific type
func (vr *ValidationReport) GetErrorsByType(errorType ValidationErrorType) []*ValidationError {
	var matchingErrors []*ValidationError
	for _, err := range vr.ValidationErrors.GetErrors() {
		if err.Metadata.ErrorType == errorType {
			matchingErrors = append(matchingErrors, err)
		}
	}
	return matchingErrors
}
