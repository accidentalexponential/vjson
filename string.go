package vjson

import (
	"regexp"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// StringField is the type for validating strings in a JSON
type StringField struct {
	name     string
	required bool

	validateMinLength bool
	minLength         int

	validateMaxLength bool
	maxLength         int

	validateFormat bool
	format         string

	validateChoices bool
	choices         []string
}

// To Force Implementing Field interface by StringField
var _ Field = (*StringField)(nil)

// GetName returns name of the field
func (s *StringField) GetName() string {
	return s.name
}

// GetType returns the Fields type
func (s *StringField) GetType() string {
	return "string"
}

// GetRequired returns true if field is required
func (s *StringField) GetRequired() bool {
	return s.required
}

// Required is called to make a field required in a JSON
func (s *StringField) Required() *StringField {
	s.required = true
	return s
}

// MinLength is called to set a minimum length to a string field
func (s *StringField) MinLength(length int) *StringField {
	if length < 0 {
		return s
	}
	s.minLength = length
	s.validateMinLength = true
	return s
}

// MaxLength is called to set a maximum length to a string field
func (s *StringField) MaxLength(length int) *StringField {
	if length < 0 {
		return s
	}
	s.maxLength = length
	s.validateMaxLength = true
	return s
}

// Format is called to set a regex format for validation of a string field
func (s *StringField) Format(format string) *StringField {
	s.format = format
	s.validateFormat = true
	return s
}

// Choices is called to set valid choices of a string field in validation
func (s *StringField) Choices(choices ...string) *StringField {
	s.choices = choices
	s.validateChoices = true
	return s
}

// Validate is used for validating a value. it returns an error if the value is invalid.
func (s *StringField) Validate(value interface{}) error {
	if value == nil {
		if !s.required {
			return nil
		}
		return errors.Errorf("Value for %s field is required", s.name)
	}

	stringValue, ok := value.(string)

	if !ok {
		return errors.Errorf("Value for %s should be a string", s.name)
	}

	var result error

	if s.validateMinLength {
		if len(stringValue) < s.minLength {
			result = multierror.Append(result, errors.Errorf("Value for %s field should have at least %d characters", s.name, s.minLength))
		}
	}

	if s.validateMaxLength {
		if len(stringValue) > s.maxLength {
			result = multierror.Append(result, errors.Errorf("Value for %s field should have at most %d characters", s.name, s.maxLength))
		}
	}

	if s.validateChoices {
		for _, choice := range s.choices {
			if stringValue == choice {
				return nil
			}
		}
		result = multierror.Append(result, errors.Errorf("Value for %s field should be one of: [%s] values", s.name, strings.Join(s.choices, ",")))
	}

	if s.validateFormat {
		r, err := regexp.Compile(s.format)
		if err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "Invalid StringField format string for field %s", s.name))
			return result
		}

		isValidFormat := r.MatchString(stringValue)

		if !isValidFormat {
			result = multierror.Append(result, errors.Wrapf(err, "Invalid StringField format string for field %s", s.name))
		}
	}

	return result
}

// String is the constructor of a string field
func String(name string) *StringField {
	return &StringField{
		name:     name,
		required: false,
		choices:  []string{},
	}
}
