package schema_validator

import (
	"fmt"
	"github.com/spf13/cobra"
)

type SchemaValidatorCommandContext struct {
}

func New(ctx *SchemaValidatorCommandContext) *cobra.Command {
	benCommand := &cobra.Command{
		Use:   "schema-validator",
		Short: "Execute static analysis for given <pattern>",
		Long:	"Execute static analysis for given <pattern>. Input should be glob or `-` for stdin",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				errMessage := "Requires at least 1 arg\n"
				return fmt.Errorf(errMessage)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error{
			var i = 0
			for i < len(args) {
				schemaValidator(args[i])
				i++
			}
			return nil
		},
	}
	return benCommand
}

func schemaValidator(arg string) {
	fmt.Println("schemaValidator", arg)
}