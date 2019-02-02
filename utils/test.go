package test

import (
	"testing"
)

// AreEqual compares two values and throws a testing error if values don't match
func AreEqual(t *testing.T, expected interface{}, actual interface{}, message string) {
	if expected != actual {
		t.Errorf("Expected: %v\nActual: %v\n%v", expected, actual, message)
	}
}
