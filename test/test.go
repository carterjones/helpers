// Package test provides helper functions that assist testers reduce boilerplate
// code when writing Golang tests.
package test

import (
	"reflect"
	"strings"
	"testing"
)

// UnexpectedErrCheck marks the test as an error if the supplied error is
// non-nil.
func UnexpectedErrCheck(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// ExpectedErrCheck first checks to see if the supplied error is nil. If it is,
// it marks the test as failed, noting that it expeceted an error but did not
// receive one. If the error is non-nil, but does not match the expected
// message, then it also marks the test as failed, noting the reason in that
// case as well.
func ExpectedErrCheck(t *testing.T, err error, msg string, v interface{}) {
	if err == nil {
		t.Errorf("expected error, but received none: %v", v)
	} else if !strings.Contains(err.Error(), msg) {
		t.Errorf("wrong error\nexpected: %v\ngot:      %v", msg, err)
	}
}

// ExpectedValueCheck determines if the expected value and the actual value
// match. If they do not, it prints both the expected and actual values and
// marks the test as failed.
func ExpectedValueCheck(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("not equal\nexpected: %v\nactual:   %v", expected, actual)
	}
}
