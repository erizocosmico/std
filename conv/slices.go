package conv

import (
	"fmt"
	"reflect"
)

// StringsToInterfaces converts a slice of strings into a slice of interfaces.
func StringsToInterfaces(from ...string) []interface{} {
	var result = make([]interface{}, len(from))
	for i, v := range from {
		result[i] = v
	}
	return result
}

// InterfacesToStrings converts a slice of interfaces into a slice of strings.
func InterfacesToStrings(from ...interface{}) ([]string, error) {
	var result = make([]string, len(from))
	for i, v := range from {
		val, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("can't convert \"%#v\" to string", v)
		}
		result[i] = val
	}
	return result, nil
}

// IntsToInterfaces converts a slice of ints into a slice of interfaces.
func IntsToInterfaces(from ...int) []interface{} {
	var result = make([]interface{}, len(from))
	for i, v := range from {
		result[i] = v
	}
	return result
}

// InterfacesToInts converts a slice of interfaces into a slice of ints.
func InterfacesToInts(from ...interface{}) ([]int, error) {
	var result = make([]int, len(from))
	for i, v := range from {
		val, ok := v.(int)
		if !ok {
			return nil, fmt.Errorf("can't convert \"%#v\" to int", v)
		}
		result[i] = val
	}
	return result, nil
}

// ToSliceOf converts the given array or slice to a slice of the given type.
//   result, err := To([]int{1, 2, 3}, int64(0))
//   fmt.Println(result.([]int64))
//   result, err = To([]mystruct{a, b, c}, (*MyInterface)(nil))
//   fmt.Println(result.([]MyInterface))
func ToSliceOf(from interface{}, to interface{}) (interface{}, error) {
	vfrom := reflect.ValueOf(from)
	vto := reflect.ValueOf(to)

	switch vfrom.Kind() {
	case reflect.Array, reflect.Slice:
		underlying := vfrom.Type().Elem()
		if !underlying.ConvertibleTo(vto.Type()) {
			return nil, fmt.Errorf("can't convert from %q to %q", underlying, vto.Type())
		}

		result := reflect.MakeSlice(reflect.SliceOf(vto.Type()), vfrom.Len(), vfrom.Cap())

		for i := 0; i < vfrom.Len(); i++ {
			result.Index(i).Set(vfrom.Index(i).Convert(vto.Type()))
		}

		return result.Interface(), nil
	default:
		return nil, fmt.Errorf("from argument needs to be slice or array")
	}
}
