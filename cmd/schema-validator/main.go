package schema_validator

import (
	"fmt"
	"github.com/datreeio/datree/pkg/extractor"
	"github.com/spf13/cobra"
)

type JsonSchemaValidator interface {
	Validate(jsonSchema string, json string) error
}

type JsonSchemaValidationPrinter interface {
	PrintJsonSchemaResults(results error) error
}

type SchemaValidatorCommandContext struct {
	JsonSchemaValidator JsonSchemaValidator
	Printer             JsonSchemaValidationPrinter
}

func New(ctx *SchemaValidatorCommandContext) *cobra.Command {
	benCommand := &cobra.Command{
		Use:   "schema-validator",
		Short: "Execute static analysis for given <pattern>",
		Long:  "Execute static analysis for given <pattern>. Input should be glob or `-` for stdin",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				errMessage := "Requires at least 1 arg\n"
				return fmt.Errorf(errMessage)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			jsonPath, _ := extractor.ReadFileContent(args[0])
			json, _ := extractor.ReadFileContent(args[1])

			err := ctx.JsonSchemaValidator.Validate(jsonPath, json)

			err = ctx.Printer.PrintJsonSchemaResults(err)

			if err != nil {
				return err
			}

			return nil
		},
	}
	return benCommand
}

func schemaValidator(arg string) {
	fmt.Println("schemaValidator", arg)
}
