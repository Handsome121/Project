package assert

import (
	"fmt"
	"reflect"
)

func compare(obj1, obj2 interface{}, kind reflect.Kind) (int, bool) {
	switch kind {
	case reflect.Int:
		{
			intobj1 := obj1.(int)
			intobj2 := obj2.(int)
			if intobj1 > intobj2 {
				return -1, true
			}
			if intobj1 == intobj2 {
				return 0, true
			}
			if intobj1 < intobj2 {
				return 1, true
			}
		}
	case reflect.Int8:
		{
			int8obj1 := obj1.(int8)
			int8obj2 := obj2.(int8)
			if int8obj1 > int8obj2 {
				return -1, true
			}
			if int8obj1 == int8obj2 {
				return 0, true
			}
			if int8obj1 < int8obj2 {
				return 1, true
			}
		}
	case reflect.Int16:
		{
			int16obj1 := obj1.(int16)
			int16obj2 := obj2.(int16)
			if int16obj1 > int16obj2 {
				return -1, true
			}
			if int16obj1 == int16obj2 {
				return 0, true
			}
			if int16obj1 < int16obj2 {
				return 1, true
			}
		}
	case reflect.Int32:
		{
			int32obj1 := obj1.(int32)
			int32obj2 := obj2.(int32)
			if int32obj1 > int32obj2 {
				return -1, true
			}
			if int32obj1 == int32obj2 {
				return 0, true
			}
			if int32obj1 < int32obj2 {
				return 1, true
			}
		}
	case reflect.Int64:
		{
			int64obj1 := obj1.(int64)
			int64obj2 := obj2.(int64)
			if int64obj1 > int64obj2 {
				return -1, true
			}
			if int64obj1 == int64obj2 {
				return 0, true
			}
			if int64obj1 < int64obj2 {
				return 1, true
			}
		}
	case reflect.Uint:
		{
			uintobj1 := obj1.(uint)
			uintobj2 := obj2.(uint)
			if uintobj1 > uintobj2 {
				return -1, true
			}
			if uintobj1 == uintobj2 {
				return 0, true
			}
			if uintobj1 < uintobj2 {
				return 1, true
			}
		}
	case reflect.Uint8:
		{
			uint8obj1 := obj1.(uint8)
			uint8obj2 := obj2.(uint8)
			if uint8obj1 > uint8obj2 {
				return -1, true
			}
			if uint8obj1 == uint8obj2 {
				return 0, true
			}
			if uint8obj1 < uint8obj2 {
				return 1, true
			}
		}
	case reflect.Uint16:
		{
			uint16obj1 := obj1.(uint16)
			uint16obj2 := obj2.(uint16)
			if uint16obj1 > uint16obj2 {
				return -1, true
			}
			if uint16obj1 == uint16obj2 {
				return 0, true
			}
			if uint16obj1 < uint16obj2 {
				return 1, true
			}
		}
	case reflect.Uint32:
		{
			uint32obj1 := obj1.(uint32)
			uint32obj2 := obj2.(uint32)
			if uint32obj1 > uint32obj2 {
				return -1, true
			}
			if uint32obj1 == uint32obj2 {
				return 0, true
			}
			if uint32obj1 < uint32obj2 {
				return 1, true
			}
		}
	case reflect.Uint64:
		{
			uint64obj1 := obj1.(uint64)
			uint64obj2 := obj2.(uint64)
			if uint64obj1 > uint64obj2 {
				return -1, true
			}
			if uint64obj1 == uint64obj2 {
				return 0, true
			}
			if uint64obj1 < uint64obj2 {
				return 1, true
			}
		}
	case reflect.Float32:
		{
			float32obj1 := obj1.(float32)
			float32obj2 := obj2.(float32)
			if float32obj1 > float32obj2 {
				return -1, true
			}
			if float32obj1 == float32obj2 {
				return 0, true
			}
			if float32obj1 < float32obj2 {
				return 1, true
			}
		}
	case reflect.Float64:
		{
			float64obj1 := obj1.(float64)
			float64obj2 := obj2.(float64)
			if float64obj1 > float64obj2 {
				return -1, true
			}
			if float64obj1 == float64obj2 {
				return 0, true
			}
			if float64obj1 < float64obj2 {
				return 1, true
			}
		}
	case reflect.String:
		{
			stringobj1 := obj1.(string)
			stringobj2 := obj2.(string)
			if stringobj1 > stringobj2 {
				return -1, true
			}
			if stringobj1 == stringobj2 {
				return 0, true
			}
			if stringobj1 < stringobj2 {
				return 1, true
			}
		}
	}

	return 0, false
}

// Greater asserts that the first element is greater than the second
//
//    assert.Greater(t, control, 1)
//    assert.Greater(t, float64(control), float64(1))
//    assert.Greater(t, "b", "a")
func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	res, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf("Can not compare type \"%s\"", reflect.TypeOf(e1)), msgAndArgs...)
	}

	if res != -1 {
		return Fail(t, fmt.Sprintf("\"%v\" is not greater than \"%v\"", e1, e2), msgAndArgs...)
	}

	return true
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqual(t, control, 1)
//    assert.GreaterOrEqual(t, control, control)
//    assert.GreaterOrEqual(t, "b", "a")
//    assert.GreaterOrEqual(t, "b", "b")
func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	res, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf("Can not compare type \"%s\"", reflect.TypeOf(e1)), msgAndArgs...)
	}

	if res != -1 && res != 0 {
		return Fail(t, fmt.Sprintf("\"%v\" is not greater than or equal to \"%v\"", e1, e2), msgAndArgs...)
	}

	return true
}

// Less asserts that the first element is less than the second
//
//    assert.Less(t, 1, control)
//    assert.Less(t, float64(1), float64(control))
//    assert.Less(t, "a", "b")
func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	res, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf("Can not compare type \"%s\"", reflect.TypeOf(e1)), msgAndArgs...)
	}

	if res != 1 {
		return Fail(t, fmt.Sprintf("\"%v\" is not less than \"%v\"", e1, e2), msgAndArgs...)
	}

	return true
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqual(t, 1, control)
//    assert.LessOrEqual(t, control, control)
//    assert.LessOrEqual(t, "a", "b")
//    assert.LessOrEqual(t, "b", "b")
func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	res, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf("Can not compare type \"%s\"", reflect.TypeOf(e1)), msgAndArgs...)
	}

	if res != 1 && res != 0 {
		return Fail(t, fmt.Sprintf("\"%v\" is not less than or equal to \"%v\"", e1, e2), msgAndArgs...)
	}

	return true
}
