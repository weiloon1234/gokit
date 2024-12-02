package utils

import (
	"errors"
	"reflect"
)

// MergeStruct merges fields from src into dest for the same struct type.
// Fields in src override those in dest, and the result retains methods.
func MergeStruct(defaults, overrides interface{}) (interface{}, error) {
	// Ensure both arguments are pointers to the same struct type
	defaultsVal := reflect.ValueOf(defaults)
	overridesVal := reflect.ValueOf(overrides)

	if defaultsVal.Kind() != reflect.Ptr || overridesVal.Kind() != reflect.Ptr {
		return nil, errors.New("both arguments must be pointers to structs")
	}

	defaultsElem := defaultsVal.Elem()
	overridesElem := overridesVal.Elem()

	if defaultsElem.Type() != overridesElem.Type() {
		return nil, errors.New("both arguments must have the same struct type")
	}

	// Create a new instance of the struct
	mergedVal := reflect.New(defaultsElem.Type()).Elem()

	// Copy fields from defaults into mergedVal
	for i := 0; i < defaultsElem.NumField(); i++ {
		mergedVal.Field(i).Set(defaultsElem.Field(i))
	}

	// Override fields with non-zero values from overrides
	for i := 0; i < overridesElem.NumField(); i++ {
		overrideField := overridesElem.Field(i)
		if !isZeroValue(overrideField) {
			mergedVal.Field(i).Set(overrideField)
		}
	}

	// Return the merged struct as a pointer
	return mergedVal.Addr().Interface(), nil
}

func isZeroValue(v reflect.Value) bool {
	zero := reflect.Zero(v.Type()).Interface()
	current := v.Interface()
	return reflect.DeepEqual(current, zero)
}
