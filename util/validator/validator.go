package validator

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidateModel validates the given model by checking if the specified fields are empty
func ValidateModel(fields []string, model interface{}) error {
	var requiredFields []string

	// Iterate over the specified fields and check if they are empty
	for _, field := range fields {
		if fieldIsEmpty(model, field) {
			requiredFields = append(requiredFields, field)
		}
	}

	// If all specified fields are not empty, return nil
	if len(requiredFields) == 0 {
		return nil
	}

	// Otherwise, return an error with the list of required fields
	return fmt.Errorf("%s", []string{strings.Join(requiredFields, ", ")})
}

// fieldIsEmpty checks if the specified field in the given model is empty
func fieldIsEmpty(model interface{}, field string) bool {
	var f reflect.Value

	// If the field is a nested field, traverse the model to get the nested field value
	if strings.Contains(field, ".") {
		fields := strings.Split(field, ".")
		r := reflect.ValueOf(model)

		for _, f := range fields {
			baseProperty := reflect.Indirect(r).FieldByName(f)

			// If the nested field is nil, return true
			if baseProperty.IsNil() {
				return true
			}

			r = baseProperty
		}
		return false
	}

	// Otherwise, get the value of the specified field in the model
	r := reflect.ValueOf(model)
	f = reflect.Indirect(r).FieldByName(field)

	// If the field is a pointer, dereference it
	if f.Kind() == reflect.Ptr {
		f = f.Elem()
	}

	// If the field is a slice, check if it is empty or if all its elements are empty strings
	if f.Kind() == reflect.Slice {
		if f.Len() == 0 {
			return true
		}
		for i := 0; i < f.Len(); i++ {
			elemValue := f.Index(i)
			if elemValue.Kind() == reflect.Ptr {
				elemValue = elemValue.Elem()
			}
			if elemValue.Kind() != reflect.String || elemValue.String() != "" {
				return false
			}
		}
		return true
	}

	// Otherwise, check if the field is an empty string
	return f.Kind() == reflect.String && f.String() == ""
}
