package jsonSchemaValidator

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)


type JsonSchemaValidator struct {

}

func New() *JsonSchemaValidator {
	return &JsonSchemaValidator{}
}

func (jsv *JsonSchemaValidator) Validate(jsonSchema string, json string) error {
	schemaLoader := gojsonschema.NewStringLoader(jsonSchema)
	documentLoader := gojsonschema.NewStringLoader(json)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

	return fmt.Errorf("validation always fails. duh")
}