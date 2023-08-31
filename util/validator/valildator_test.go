package validator

import "testing"

func TestValidateModel(t *testing.T) {
	// Create a mock model object
	type TestModel struct {
		Field1 *string
		Field2 *string
		Field3 *string
	}
	field1 := "test"
	field2 := "123"
	model := TestModel{
		Field1: &field1,
		Field2: &field2,
		Field3: nil,
	}

	// Test with all required fields present
	fields := []string{"Field1", "Field2"}
	err := ValidateModel(fields, model)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test with one required field missing
	fields = []string{"Field1", "Field2", "Field3"}
	err = ValidateModel(fields, model)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	expectedError := "[Field3]"
	if err.Error() != expectedError {
		t.Errorf("Unexpected error message: got %v want %v", err.Error(), expectedError)
	}
}

func TestFieldIsEmpty(t *testing.T) {
	// Create a mock model object
	type TestModel struct {
		Field1 *string
		Field2 *string
		Field3 *string
	}
	field1 := "test"
	field2 := "123"
	model := TestModel{
		Field1: &field1,
		Field2: &field2,
		Field3: nil,
	}

	// Test with a top-level field that is not empty
	field := "Field1"
	isEmpty := fieldIsEmpty(model, field)
	if isEmpty {
		t.Errorf("Unexpected result: field %v is not empty", field)
	}

	// Test with a top-level field that is empty
	field = "Field3"
	isEmpty = fieldIsEmpty(model, field)
	if !isEmpty {
		t.Errorf("Unexpected result: field %v is empty", field)
	}

	// Test with a nested field that is not empty
	field = "Field2"
	isEmpty = fieldIsEmpty(model, field)
	if isEmpty {
		t.Errorf("Unexpected result: field %v is not empty", field)
	}

	// Test with a nested field that is empty
	field = "Field3.Length"
	isEmpty = fieldIsEmpty(model, field)
	if !isEmpty {
		t.Errorf("Unexpected result: field %v is empty", field)
	}
}
