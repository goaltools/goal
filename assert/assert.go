// Package assert provides a set of tools for testing
// your handler functions and controllers.
package assert

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

// Type provides methods for getting GET, POST, etc. requests
// and assertions for testing of handler functions.
type Type struct {
	Body           *Assertion
	Request        *http.Request
	ResponseWriter *httptest.ResponseRecorder
	url            string

	t *testing.T
}

// Assertion is a type that represents an actual result.
// It has a number of methods to compare it with expected value.
type Assertion struct {
	actual interface{}
	t      *testing.T
}

// New allocates and returns a new Type.
func New(t *testing.T) *Type {
	return &Type{
		t: t,
	}
}

// NewAssertion allocates and returns a new Assertion.
func (t *Type) NewAssertion(actual interface{}) *Assertion {
	return &Assertion{
		actual: actual,
		t:      t.t,
	}
}

// Args is a shortcut for myHandlerFn(t.ResponseWriter, t.Request).
// It makes possible to use myHandlerFn(t.Args()) instead.
func (t *Type) Args() (*httptest.ResponseRecorder, *http.Request) {
	return t.ResponseWriter, t.Request
}

// String returns string representation of assertion.
func (a Assertion) String() string {
	return fmt.Sprintf("%v", a.actual)
}

// Equal makes sure two values are equal to each other using reflect.DeepEqual.
// It panics if they are not equal.
func (a *Assertion) Equal(expected interface{}) {
	if !reflect.DeepEqual(expected, a.actual) {
		a.t.Errorf("(expected) %v != %v (actual)", expected, a.actual)
	}
}

// NotEqual is an opposite of Equal.
func (a *Assertion) NotEqual(expected interface{}) {
	if reflect.DeepEqual(expected, a.actual) {
		a.t.Errorf("(expected) %v == %v (actual)", expected, a.actual)
	}
}

// True panics if expression is not true.
func (t *Type) True(exp bool) {
	if !exp {
		t.t.Errorf("(expected) true != false (actual)")
	}
}

// Contains panics if fragment is not a part of actual string.
func (a *Assertion) Contains(fragment string) {
	if !strings.Contains(a.String(), fragment) {
		a.t.Errorf(`(actual) does not contain "%s".`, fragment)
	}
}

// NotContains is an opposite of Contains.
func (a *Assertion) NotContains(source, fragment string) {
	if strings.Contains(a.String(), fragment) {
		a.t.Errorf(`(actual) contains "%s".`, fragment)
	}
}

// MatchesRegex makes sure the actual string matches the given regular expression.
// It panics otherwise.
func (a *Assertion) MatchesRegex(regex string) {
	r := regexp.MustCompile(regex)
	if !r.MatchString(a.String()) {
		a.t.Errorf(`(actual) does not match regexp "%s".`, regex)
	}
}

// Status panics panics if ResponseWriter's code is not equal
// to the given status code.
func (t *Type) Status(expected int) {
	if t.ResponseWriter.Code != expected {
		t.t.Errorf("(expected) %v != %v (actual) status code", expected, t.ResponseWriter.Code)
	}
}

// StatusOK makes sure status code is equal to 200.
func (t *Type) StatusOK() {
	t.Status(200)
}

// StatusNotFound makes sure status code is equal to 404.
func (t *Type) StatusNotFound() {
	t.Status(404)
}

// Header returns a header with the given name in a form
// of assertion, so assert methods can be called on it as follows:
// t.Header("Content-Type").Equal("plain/text")
func (t *Type) Header(name string) *Assertion {
	return t.NewAssertion(t.ResponseWriter.Header().Get(name))
}

// ContentType is a shortcut for Header("Content-Type").Equal("value").
func (t *Type) ContentType(expected string) {
	if ct := t.Header("Content-Type"); ct.actual != expected {
		t.t.Errorf("(expected) %v != %v (actual) content type", expected, ct.actual)
	}
}
