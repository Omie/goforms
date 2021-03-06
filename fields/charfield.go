package fields

import (
	"errors"
	"fmt"
)

type CharField struct {
	BaseField
	Max int
	Min int
}

// Check whether the given string value is valid for this field
// and return the cleaned value or a relevant error.
func (f CharField) Clean(value string) (interface{}, ValidationError) {
	// Ensure value is between max and min length,
	if f.Max != 0 && len(value) > f.Max {
		return nil, errors.New(fmt.Sprint(
			"The value must have a maximum length of ",
			f.Max, " characters."))
	}
	if len(value) < f.Min {
		return nil, errors.New(fmt.Sprint(
			"The value must have a minimum length of ",
			f.Min, " characters."))
	}

	return value, nil
}

// Create and initialise the new fields with the given defaults.
func NewCharField(defaults Defaults) CharField {
	// This is a helper for creating fields, so that users of goforms/fields
	// don't need to do things like:
	// "purchaseCount": fields.IntegerField{fields.BaseField{Required: true}},
	// More details here:
	// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/FS_H0SiEioA
	// If that ever changes, we should be able to use simple
	// struct literals.
	field := CharField{}
	for fieldName, value := range defaults {
		switch fieldName {
		case "Required":
			if v, ok := value.(bool); ok {
				field.Required = v
			}
		case "Min":
			if v, ok := value.(int); ok {
				field.Min = v
			}
		case "Max":
			if v, ok := value.(int); ok {
				field.Max = v
			}
		}
	}
	return field
}
