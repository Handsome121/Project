/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package require

import (
	assert "github.com/stretchr/testify/assert"
	http "net/http"
	url "net/url"
	time "time"
)

// Condition uses a Comparison to assert a complex condition.
func Condition(t TestingT, comp assert.Comparison, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Condition(t, comp, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Conditionf uses a Comparison to assert a complex condition.
func Conditionf(t TestingT, comp assert.Comparison, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Conditionf(t, comp, msg, args...) {
		return
	}
	t.FailNow()
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Contains(t, "Hello World", "World")
//    assert.Contains(t, ["Hello", "World"], "World")
//    assert.Contains(t, {"Hello": "World"}, "Hello")
func Contains(t TestingT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Contains(t, s, contains, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Containsf(t, "Hello World", "World", "error message %s", "formatted")
//    assert.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
//    assert.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
func Containsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Containsf(t, s, contains, msg, args...) {
		return
	}
	t.FailNow()
}

// DirExists checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.
func DirExists(t TestingT, path string, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.DirExists(t, path, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// DirExistsf checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.
func DirExistsf(t TestingT, path string, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.DirExistsf(t, path, msg, args...) {
		return
	}
	t.FailNow()
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, control, 3], [1, 3, 3, control])
func ElementsMatch(t TestingT, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.ElementsMatch(t, listA, listB, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatchf(t, [1, 3, control, 3], [1, 3, 3, control], "error message %s", "formatted")
func ElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.ElementsMatchf(t, listA, listB, msg, args...) {
		return
	}
	t.FailNow()
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Empty(t, obj)
func Empty(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Empty(t, object, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Emptyf(t, obj, "error message %s", "formatted")
func Emptyf(t TestingT, object interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Emptyf(t, object, msg, args...) {
		return
	}
	t.FailNow()
}

// Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equal(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Equal(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualError(t, err,  expectedErrorString)
func EqualError(t TestingT, theError error, errString string, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.EqualError(t, theError, errString, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
func EqualErrorf(t TestingT, theError error, errString string, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.EqualErrorf(t, theError, errString, msg, args...) {
		return
	}
	t.FailNow()
}

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValues(t, uint32(123), int32(123))
func EqualValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.EqualValues(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValuesf(t, uint32(123, "error message %s", "formatted"), int32(123))
func EqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.EqualValuesf(t, expected, actual, msg, args...) {
		return
	}
	t.FailNow()
}

// Equalf asserts that two objects are equal.
//
//    assert.Equalf(t, 123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Equalf(t, expected, actual, msg, args...) {
		return
	}
	t.FailNow()
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err) {
// 	   assert.Equal(t, expectedError, err)
//   }
func Error(t TestingT, err error, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Error(t, err, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Errorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
func Errorf(t TestingT, err error, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Errorf(t, err, msg, args...) {
		return
	}
	t.FailNow()
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
func Eventually(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	if assert.Eventually(t, condition, waitFor, tick, msgAndArgs...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Eventuallyf asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventuallyf(t, func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
func Eventuallyf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	if assert.Eventuallyf(t, condition, waitFor, tick, msg, args...) {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.FailNow()
}

// Exactly asserts that two objects are equal in value and type.
//
//    assert.Exactly(t, int32(123), int64(123))
func Exactly(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Exactly(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Exactlyf asserts that two objects are equal in value and type.
//
//    assert.Exactlyf(t, int32(123, "error message %s", "formatted"), int64(123))
func Exactlyf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Exactlyf(t, expected, actual, msg, args...) {
		return
	}
	t.FailNow()
}

// Fail reports a failure through
func Fail(t TestingT, failureMessage string, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Fail(t, failureMessage, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// FailNow fails test
func FailNow(t TestingT, failureMessage string, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.FailNow(t, failureMessage, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// FailNowf fails test
func FailNowf(t TestingT, failureMessage string, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.FailNowf(t, failureMessage, msg, args...) {
		return
	}
	t.FailNow()
}

// Failf reports a failure through
func Failf(t TestingT, failureMessage string, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Failf(t, failureMessage, msg, args...) {
		return
	}
	t.FailNow()
}

// False asserts that the specified value is false.
//
//    assert.False(t, myBool)
func False(t TestingT, value bool, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.False(t, value, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Falsef asserts that the specified value is false.
//
//    assert.Falsef(t, myBool, "error message %s", "formatted")
func Falsef(t TestingT, value bool, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Falsef(t, value, msg, args...) {
		return
	}
	t.FailNow()
}

// FileExists checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.
func FileExists(t TestingT, path string, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.FileExists(t, path, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// FileExistsf checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.
func FileExistsf(t TestingT, path string, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.FileExistsf(t, path, msg, args...) {
		return
	}
	t.FailNow()
}

// Greater asserts that the first element is greater than the second
//
//    assert.Greater(t, control, 1)
//    assert.Greater(t, float64(control), float64(1))
//    assert.Greater(t, "b", "a")
func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Greater(t, e1, e2, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqual(t, control, 1)
//    assert.GreaterOrEqual(t, control, control)
//    assert.GreaterOrEqual(t, "b", "a")
//    assert.GreaterOrEqual(t, "b", "b")
func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.GreaterOrEqual(t, e1, e2, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// GreaterOrEqualf asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqualf(t, control, 1, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, control, control, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "a", "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "b", "error message %s", "formatted")
func GreaterOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.GreaterOrEqualf(t, e1, e2, msg, args...) {
		return
	}
	t.FailNow()
}

// Greaterf asserts that the first element is greater than the second
//
//    assert.Greaterf(t, control, 1, "error message %s", "formatted")
//    assert.Greaterf(t, float64(control, "error message %s", "formatted"), float64(1))
//    assert.Greaterf(t, "b", "a", "error message %s", "formatted")
func Greaterf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Greaterf(t, e1, e2, msg, args...) {
		return
	}
	t.FailNow()
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPBodyContains(t, handler, method, url, values, str, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPBodyContainsf(t, handler, method, url, values, str, msg, args...) {
		return
	}
	t.FailNow()
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPBodyNotContains(t, handler, method, url, values, str, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPBodyNotContainsf(t, handler, method, url, values, str, msg, args...) {
		return
	}
	t.FailNow()
}

// HTTPError asserts that a specified handler returns an error status code.
//
//  assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPError(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPError(t, handler, method, url, values, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// HTTPErrorf asserts that a specified handler returns an error status code.
//
//  assert.HTTPErrorf(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func HTTPErrorf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPErrorf(t, handler, method, url, values, msg, args...) {
		return
	}
	t.FailNow()
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirect(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPRedirect(t, handler, method, url, values, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirectf(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true, "error message %s", "formatted") or not (false).
func HTTPRedirectf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPRedirectf(t, handler, method, url, values, msg, args...) {
		return
	}
	t.FailNow()
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccess(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPSuccess(t, handler, method, url, values, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// HTTPSuccessf asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccessf(t, myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccessf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.HTTPSuccessf(t, handler, method, url, values, msg, args...) {
		return
	}
	t.FailNow()
}

// Implements asserts that an object is implemented by the specified interface.
//
//    assert.Implements(t, (*MyInterface)(nil), new(MyObject))
func Implements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Implements(t, interfaceObject, object, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Implementsf asserts that an object is implemented by the specified interface.
//
//    assert.Implementsf(t, (*MyInterface, "error message %s", "formatted")(nil), new(MyObject))
func Implementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Implementsf(t, interfaceObject, object, msg, args...) {
		return
	}
	t.FailNow()
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 assert.InDelta(t, math.Pi, (22 / 7.0), 0.01)
func InDelta(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InDelta(t, expected, actual, delta, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func InDeltaMapValues(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func InDeltaMapValuesf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InDeltaMapValuesf(t, expected, actual, delta, msg, args...) {
		return
	}
	t.FailNow()
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func InDeltaSlice(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InDeltaSlice(t, expected, actual, delta, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// InDeltaSlicef is the same as InDelta, except it compares two slices.
func InDeltaSlicef(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InDeltaSlicef(t, expected, actual, delta, msg, args...) {
		return
	}
	t.FailNow()
}

// InDeltaf asserts that the two numerals are within delta of each other.
//
// 	 assert.InDeltaf(t, math.Pi, (22 / 7.0, "error message %s", "formatted"), 0.01)
func InDeltaf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InDeltaf(t, expected, actual, delta, msg, args...) {
		return
	}
	t.FailNow()
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
func InEpsilon(t TestingT, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InEpsilon(t, expected, actual, epsilon, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func InEpsilonSlice(t TestingT, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InEpsilonSlice(t, expected, actual, epsilon, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.
func InEpsilonSlicef(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InEpsilonSlicef(t, expected, actual, epsilon, msg, args...) {
		return
	}
	t.FailNow()
}

// InEpsilonf asserts that expected and actual have a relative error less than epsilon
func InEpsilonf(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.InEpsilonf(t, expected, actual, epsilon, msg, args...) {
		return
	}
	t.FailNow()
}

// IsType asserts that the specified objects are of the same type.
func IsType(t TestingT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.IsType(t, expectedType, object, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// IsTypef asserts that the specified objects are of the same type.
func IsTypef(t TestingT, expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.IsTypef(t, expectedType, object, msg, args...) {
		return
	}
	t.FailNow()
}

// JSONEq asserts that two JSON strings are equivalent.
//
//  assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
func JSONEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.JSONEq(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// JSONEqf asserts that two JSON strings are equivalent.
//
//  assert.JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
func JSONEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.JSONEqf(t, expected, actual, msg, args...) {
		return
	}
	t.FailNow()
}

// YAMLEq asserts that two YAML strings are equivalent.
func YAMLEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.YAMLEq(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// YAMLEqf asserts that two YAML strings are equivalent.
func YAMLEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.YAMLEqf(t, expected, actual, msg, args...) {
		return
	}
	t.FailNow()
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    assert.Len(t, mySlice, 3)
func Len(t TestingT, object interface{}, length int, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Len(t, object, length, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Lenf asserts that the specified object has specific length.
// Lenf also fails if the object has a type that len() not accept.
//
//    assert.Lenf(t, mySlice, 3, "error message %s", "formatted")
func Lenf(t TestingT, object interface{}, length int, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Lenf(t, object, length, msg, args...) {
		return
	}
	t.FailNow()
}

// Less asserts that the first element is less than the second
//
//    assert.Less(t, 1, control)
//    assert.Less(t, float64(1), float64(control))
//    assert.Less(t, "a", "b")
func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Less(t, e1, e2, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqual(t, 1, control)
//    assert.LessOrEqual(t, control, control)
//    assert.LessOrEqual(t, "a", "b")
//    assert.LessOrEqual(t, "b", "b")
func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.LessOrEqual(t, e1, e2, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// LessOrEqualf asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqualf(t, 1, control, "error message %s", "formatted")
//    assert.LessOrEqualf(t, control, control, "error message %s", "formatted")
//    assert.LessOrEqualf(t, "a", "b", "error message %s", "formatted")
//    assert.LessOrEqualf(t, "b", "b", "error message %s", "formatted")
func LessOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.LessOrEqualf(t, e1, e2, msg, args...) {
		return
	}
	t.FailNow()
}

// Lessf asserts that the first element is less than the second
//
//    assert.Lessf(t, 1, control, "error message %s", "formatted")
//    assert.Lessf(t, float64(1, "error message %s", "formatted"), float64(control))
//    assert.Lessf(t, "a", "b", "error message %s", "formatted")
func Lessf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Lessf(t, e1, e2, msg, args...) {
		return
	}
	t.FailNow()
}

// Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
func Nil(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Nil(t, object, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Nilf asserts that the specified object is nil.
//
//    assert.Nilf(t, err, "error message %s", "formatted")
func Nilf(t TestingT, object interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Nilf(t, object, msg, args...) {
		return
	}
	t.FailNow()
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoError(t, err) {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
func NoError(t TestingT, err error, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NoError(t, err, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NoErrorf asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoErrorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
func NoErrorf(t TestingT, err error, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NoErrorf(t, err, msg, args...) {
		return
	}
	t.FailNow()
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContains(t, "Hello World", "Earth")
//    assert.NotContains(t, ["Hello", "World"], "Earth")
//    assert.NotContains(t, {"Hello": "World"}, "Earth")
func NotContains(t TestingT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotContains(t, s, contains, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContainsf(t, "Hello World", "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, ["Hello", "World"], "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, {"Hello": "World"}, "Earth", "error message %s", "formatted")
func NotContainsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotContainsf(t, s, contains, msg, args...) {
		return
	}
	t.FailNow()
}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmpty(t, obj) {
//    assert.Equal(t, "two", obj[1])
//  }
func NotEmpty(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotEmpty(t, object, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotEmptyf asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmptyf(t, obj, "error message %s", "formatted") {
//    assert.Equal(t, "two", obj[1])
//  }
func NotEmptyf(t TestingT, object interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotEmptyf(t, object, msg, args...) {
		return
	}
	t.FailNow()
}

// NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqual(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotEqual(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotEqualf asserts that the specified values are NOT equal.
//
//    assert.NotEqualf(t, obj1, obj2, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqualf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotEqualf(t, expected, actual, msg, args...) {
		return
	}
	t.FailNow()
}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
func NotNil(t TestingT, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotNil(t, object, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotNilf asserts that the specified object is not nil.
//
//    assert.NotNilf(t, err, "error message %s", "formatted")
func NotNilf(t TestingT, object interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotNilf(t, object, msg, args...) {
		return
	}
	t.FailNow()
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanics(t, func(){ RemainCalm() })
func NotPanics(t TestingT, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotPanics(t, f, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanicsf(t, func(){ RemainCalm() }, "error message %s", "formatted")
func NotPanicsf(t TestingT, f assert.PanicTestFunc, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotPanicsf(t, f, msg, args...) {
		return
	}
	t.FailNow()
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//  assert.NotRegexp(t, "^start", "it's not starting")
func NotRegexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotRegexp(t, rx, str, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotRegexpf asserts that a specified regexp does not match a string.
//
//  assert.NotRegexpf(t, regexp.MustCompile("starts", "error message %s", "formatted"), "it's starting")
//  assert.NotRegexpf(t, "^start", "it's not starting", "error message %s", "formatted")
func NotRegexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotRegexpf(t, rx, str, msg, args...) {
		return
	}
	t.FailNow()
}

// NotSubset asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubset(t, [1, 3, memory], [1, control], "But [1, 3, memory] does not contain [1, control]")
func NotSubset(t TestingT, list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotSubset(t, list, subset, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotSubsetf asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubsetf(t, [1, 3, memory], [1, control], "But [1, 3, memory] does not contain [1, control]", "error message %s", "formatted")
func NotSubsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotSubsetf(t, list, subset, msg, args...) {
		return
	}
	t.FailNow()
}

// NotZero asserts that i is not the zero value for its type.
func NotZero(t TestingT, i interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotZero(t, i, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// NotZerof asserts that i is not the zero value for its type.
func NotZerof(t TestingT, i interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.NotZerof(t, i, msg, args...) {
		return
	}
	t.FailNow()
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panics(t, func(){ GoCrazy() })
func Panics(t TestingT, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Panics(t, f, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
func PanicsWithValue(t TestingT, expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.PanicsWithValue(t, expected, f, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValuef(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
func PanicsWithValuef(t TestingT, expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.PanicsWithValuef(t, expected, f, msg, args...) {
		return
	}
	t.FailNow()
}

// Panicsf asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panicsf(t, func(){ GoCrazy() }, "error message %s", "formatted")
func Panicsf(t TestingT, f assert.PanicTestFunc, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Panicsf(t, f, msg, args...) {
		return
	}
	t.FailNow()
}

// Regexp asserts that a specified regexp matches a string.
//
//  assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
//  assert.Regexp(t, "start...$", "it's not starting")
func Regexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Regexp(t, rx, str, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Regexpf asserts that a specified regexp matches a string.
//
//  assert.Regexpf(t, regexp.MustCompile("start", "error message %s", "formatted"), "it's starting")
//  assert.Regexpf(t, "start...$", "it's not starting", "error message %s", "formatted")
func Regexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Regexpf(t, rx, str, msg, args...) {
		return
	}
	t.FailNow()
}

// Same asserts that two pointers reference the same object.
//
//    assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Same(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Same(t, expected, actual, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Samef asserts that two pointers reference the same object.
//
//    assert.Samef(t, ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Samef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Samef(t, expected, actual, msg, args...) {
		return
	}
	t.FailNow()
}

// Subset asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subset(t, [1, control, 3], [1, control], "But [1, control, 3] does contain [1, control]")
func Subset(t TestingT, list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Subset(t, list, subset, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Subsetf asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subsetf(t, [1, control, 3], [1, control], "But [1, control, 3] does contain [1, control]", "error message %s", "formatted")
func Subsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Subsetf(t, list, subset, msg, args...) {
		return
	}
	t.FailNow()
}

// True asserts that the specified value is true.
//
//    assert.True(t, myBool)
func True(t TestingT, value bool, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.True(t, value, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Truef asserts that the specified value is true.
//
//    assert.Truef(t, myBool, "error message %s", "formatted")
func Truef(t TestingT, value bool, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Truef(t, value, msg, args...) {
		return
	}
	t.FailNow()
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
func WithinDuration(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.WithinDuration(t, expected, actual, delta, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// WithinDurationf asserts that the two times are within duration delta of each other.
//
//   assert.WithinDurationf(t, time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
func WithinDurationf(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.WithinDurationf(t, expected, actual, delta, msg, args...) {
		return
	}
	t.FailNow()
}

// Zero asserts that i is the zero value for its type.
func Zero(t TestingT, i interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Zero(t, i, msgAndArgs...) {
		return
	}
	t.FailNow()
}

// Zerof asserts that i is the zero value for its type.
func Zerof(t TestingT, i interface{}, msg string, args ...interface{}) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if assert.Zerof(t, i, msg, args...) {
		return
	}
	t.FailNow()
}
