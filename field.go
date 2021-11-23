package vjson

// Field is the abstraction on a field in a json.
// different field types can be implemented with implementing this interface.
type Field interface {
	GetName() string
	GetType() string
	GetRequired() bool
	Validate(interface{}) error
}
