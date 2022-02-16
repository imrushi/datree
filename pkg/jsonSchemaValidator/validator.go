package jsonSchemaValidator

import (
	"github.com/xeipuuv/gojsonschema"
)

type JsonSchemaValidator struct {

}

func New() *JsonSchemaValidator {
	return &JsonSchemaValidator{}
}

type Result = gojsonschema.Result

func (jsv *JsonSchemaValidator) Validate(jsonSchema string, json string) (*Result, error) {
	schemaLoader := gojsonschema.NewStringLoader(jsonSchema)
	documentLoader := gojsonschema.NewStringLoader(json)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return result, err
	}

	if result.Valid() {
		return result, err
	} else {
		return result, err
	}

	return result, err
}