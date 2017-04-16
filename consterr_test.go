package consterr

import "testing"

const (
	err1 Error = "Test error 1"
	err2 Error = "Test error 2"
)

var errs = []error{err1, err2}
var expected = []string{"Test error 1", "Test error 2"}

func TestString(t *testing.T) {
	for i, err := range errs {
		expect := expected[i]
		str := err.Error()
		if str != expect {
			t.Errorf("Expected error message \"%s\" but got \"%s\"", expect, str)
		} else {
			t.Logf("Error message \"%s\" correct", str)
		}
	}
}
