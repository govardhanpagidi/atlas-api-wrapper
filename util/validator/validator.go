package validator

import (
	"fmt"
	"reflect"
	"strings"
)

func ValidateModel(fields []string, model interface{}) error {
	var requiredFields []string

	for _, field := range fields {
		if fieldIsEmpty(model, field) {
			requiredFields = append(requiredFields, field)
		}
	}
	if len(requiredFields) == 0 {
		return nil
	}
	return fmt.Errorf("the following fields are missing: %s", []string{strings.Join(requiredFields, ", ")})
}

func fieldIsEmpty(model interface{}, field string) bool {
	var f reflect.Value
	if strings.Contains(field, ".") {
		fields := strings.Split(field, ".")
		r := reflect.ValueOf(model)

		for _, f := range fields {
			fmt.Println(f)
			baseProperty := reflect.Indirect(r).FieldByName(f)

			if baseProperty.IsNil() {
				return true
			}

			r = baseProperty
		}
		return false
	}
	r := reflect.ValueOf(model)
	f = reflect.Indirect(r).FieldByName(field)
	print(f.Elem().String())
	return f.IsNil() || (f.IsValid() && f.Elem().String() == "")
}
