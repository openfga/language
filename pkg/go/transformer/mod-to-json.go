package transformer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v3"
)

type ModFile struct {
	Schema   string   `json:"schema"`
	Module   string   `json:"module"`
	Contents []string `json:"contents"`
}

type YAMLModFile struct {
	Schema   yaml.Node `yaml:"schema"`
	Module   yaml.Node `yaml:"module"`
	Contents yaml.Node `yaml:"contents"`
}

type ModFileValidationErrorMetadata struct{}

type ModFileValidationError struct {
	Line, Column int
	Msg          string
}

func (e *ModFileValidationError) Error() string {
	return fmt.Sprintf("validation error at line=%d, column=%d: %s", e.Line, e.Column, e.Msg)
}

type ModFileValidationMultipleError multierror.Error

func (e *ModFileValidationMultipleError) Error() string {
	errors := e.Errors

	pluralS := ""
	if len(errors) > 1 {
		pluralS = "s"
	}

	errorsString := []string{}
	for _, item := range errors {
		errorsString = append(errorsString, item.Error())
	}

	return fmt.Sprintf("%d error%s occurred:\n\t* %s\n\n", len(errors), pluralS, strings.Join(errorsString, "\n\t* "))
}

const (
	stringNode = "!!str"
	seqNode    = "!!seq"
)

func TransformModFile(data string) (*ModFile, error) { //nolint:cyclop
	yamlModFile := &YAMLModFile{}

	err := yaml.Unmarshal([]byte(data), yamlModFile)
	if err != nil {
		return nil, err
	}

	modFile := &ModFile{}
	errors := &multierror.Error{}

	switch {
	case yamlModFile.Schema.IsZero():
		errors = multierror.Append(errors, &ModFileValidationError{
			Msg:    "missing schema field",
			Line:   0,
			Column: 0,
		})
	case yamlModFile.Schema.Tag != stringNode:
		errors = multierror.Append(errors, &ModFileValidationError{
			Msg: fmt.Sprintf(
				"unexpected schema type, expected string got value %s",
				yamlModFile.Schema.Value,
			),
			Line:   yamlModFile.Schema.Line - 1,
			Column: yamlModFile.Schema.Column - 1,
		})
	case yamlModFile.Schema.Value != "1.2":
		errors = multierror.Append(errors, &ModFileValidationError{
			Msg:    "unsupported schema version, fga.mod only supported in version `1.2`",
			Line:   yamlModFile.Schema.Line - 1,
			Column: yamlModFile.Schema.Column - 1,
		})
	default:
		modFile.Schema = yamlModFile.Schema.Value
	}

	switch {
	case yamlModFile.Module.IsZero():
		errors = multierror.Append(errors, &ModFileValidationError{
			Msg:    "missing module field",
			Line:   0,
			Column: 0,
		})
	case yamlModFile.Module.Tag != stringNode:
		errors = multierror.Append(errors, &ModFileValidationError{
			Msg: fmt.Sprintf(
				"unexpected module type, expected string got value %s",
				yamlModFile.Module.Value,
			),
			Line:   yamlModFile.Module.Line - 1,
			Column: yamlModFile.Module.Column - 1,
		})
	default:
		modFile.Module = yamlModFile.Module.Value
	}

	switch {
	case yamlModFile.Contents.IsZero():
		errors = multierror.Append(errors, &ModFileValidationError{
			Msg:    "missing contents field",
			Line:   0,
			Column: 0,
		})
	case yamlModFile.Contents.Tag != seqNode:
		errors = multierror.Append(errors, &ModFileValidationError{
			Msg: fmt.Sprintf(
				"unexpected contents type, expected list of strings got value %s",
				yamlModFile.Contents.Value,
			),
			Line:   yamlModFile.Contents.Line - 1,
			Column: yamlModFile.Contents.Column - 1,
		})
	default:
		contents := []string{}
		for _, file := range yamlModFile.Contents.Content {
			contents = append(contents, file.Value)

			if file.Tag != stringNode {
				errors = multierror.Append(errors, &ModFileValidationError{
					Msg: fmt.Sprintf(
						"unexpected contents item type, expected string got value %s",
						file.Value,
					),
					Line:   file.Line - 1,
					Column: file.Column - 1,
				})
			} else if !strings.HasSuffix(file.Value, ".fga") {
				errors = multierror.Append(errors, &ModFileValidationError{
					Msg:    "contents items should use fga file extension, got " + file.Value,
					Line:   file.Line - 1,
					Column: file.Column - 1,
				})
			}
		}

		modFile.Contents = contents
	}

	if len(errors.Errors) != 0 {
		return nil, &ModFileValidationMultipleError{
			Errors: errors.Errors,
		}
	}

	return modFile, nil
}
