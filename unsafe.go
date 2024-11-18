package testutils

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

var (
	ErrFieldNotFound = errors.New("field not found")
	ErrNonStructPtr  = errors.New("the first parameter must be a pointer to a struct")
)

// https://stackoverflow.com/a/78879092/9021186
func getFieldOffset(instance interface{}, fieldName string) (uintptr, error) {
	instanceValue := reflect.ValueOf(instance)
	// Don't need to check if the type is a pointer, as it is already enforced by the function signature.

	instanceType := instanceValue.Type().Elem()
	if instanceType.Kind() != reflect.Struct {
		return 0, ErrNonStructPtr
	}

	field, found := instanceType.FieldByName(fieldName)
	if !found {
		return 0, fmt.Errorf("field '%s': %w", fieldName, ErrFieldNotFound)
	}

	return field.Offset, nil
}

// AssignPrivateField allows to access and modify private struct fields. Because sometimes just fuck everything.
func AssignPrivateField[T any, V any](src *T, field string, value V) error {
	offset, err := getFieldOffset(src, field)
	if err != nil {
		return fmt.Errorf("get offset: %w", err)
	}

	fieldPtr := (*V)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + offset))

	*fieldPtr = value
	return nil
}

// ReadPrivateField reads the value of a private, inaccessible field.
func ReadPrivateField[T any, V any](src *T, field string) (V, error) {
	var output V

	offset, err := getFieldOffset(src, field)
	if err != nil {
		return output, fmt.Errorf("get offset: %w", err)
	}

	output = *(*V)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + offset))
	return output, nil
}
