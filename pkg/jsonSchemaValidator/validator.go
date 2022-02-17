package jsonSchemaValidator

import (
	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

type JsonSchemaValidator struct {

}

func New() *JsonSchemaValidator {
	return &JsonSchemaValidator{}
}

type Result = gojsonschema.Result

func (jsv *JsonSchemaValidator) Validate(schemaContent string, yamlContent string) (*Result, error) {

	jsonSchema, _ := yaml.YAMLToJSON([]byte(schemaContent))

	json, _ := yaml.YAMLToJSON([]byte(yamlContent))

	schemaLoader := gojsonschema.NewStringLoader(string(jsonSchema))
	documentLoader := gojsonschema.NewStringLoader(string(json))

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