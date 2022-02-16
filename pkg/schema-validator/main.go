package schema_validator

import (
	"fmt"
	"github.com/spf13/cobra"
)

type SchemaValidatorCommandContext struct {
	SchemaValidator string
}

func New(ctx *SchemaValidatorCommandContext) *cobra.Command {
	benCommand := &cobra.Command{
		Use:   "schema-validator",
		Short: "Execute static analysis for given <pattern>",
		Long:	"Execute static analysis for given <pattern>. Input should be glob or `-` for stdin",
		RunE: func(cmd *cobra.Command, args []string) error{
			schemaValidator(ctx)
			return nil
		},
	}
	return benCommand
}

func schemaValidator(ctx *SchemaValidatorCommandContext) {
	fmt.Println("ben test OUTSIDE", ctx)
}