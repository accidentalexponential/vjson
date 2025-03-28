package vjson

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ObjectField is the type for validating another JSON object in a JSON
type ObjectField struct {
	name     string
	required bool
	schema   Schema
}

// To Force Implementing Field interface by ObjectField
var _ Field = (*ObjectField)(nil)

// GetName returns name of the field
func (o *ObjectField) GetName() string {
	return o.name
}

// GetType returns the Fields type
func (o *ObjectField) GetType() string {
	return "object"
}

// GetRequired returns true if field is required
func (o *ObjectField) GetRequired() bool {
	return o.required
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (o *ObjectField) Validate(v interface{}) error {
	if v == nil {
		if !o.required {
			return nil
		}
		return errors.Errorf("Value for %s field is required", o.name)
	}

	// The input is either string or an interface{} object
	value, ok := v.(string)

	var err error
	var jsonBytes []byte
	if !ok {
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			return errors.Errorf("Value for %s should be an object", o.name)
		}
	} else {
		return o.schema.ValidateString(value)
	}

	return o.schema.ValidateBytes(jsonBytes)
}

// Required is called to make a field required in a JSON
func (o *ObjectField) Required() *ObjectField {
	o.required = true
	return o
}

// Object is the constructor of an object field
func Object(name string, schema Schema) *ObjectField {
	return &ObjectField{
		name:     name,
		required: false,
		schema:   schema,
	}
}
