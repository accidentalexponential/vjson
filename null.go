package vjson

import "github.com/pkg/errors"

// NullField is the type for validating floats in a JSON
type NullField struct {
	name string
}

// To Force Implementing Field interface by NullField
var _ Field = (*NullField)(nil)

// GetName returns name of the field
func (n *NullField) GetName() string {
	return n.name
}

// GetType returns the Fields type
func (n *NullField) GetType() string {
	return "null"
}

// GetRequired returns true if field is required
func (n *NullField) GetRequired() bool {
	return false
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (n *NullField) Validate(input interface{}) error {
	if input == nil {
		return nil
	}
	return errors.Errorf("Value for %s should be null", n.name)
}

// Null is the constructor of a null field in a JSON.
func Null(name string) *NullField {
	return &NullField{
		name: name,
	}
}
