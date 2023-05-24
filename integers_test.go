// integers_test.go
package main

import "testing"

func TestMultiply(t *testing.T) {
	got := Multiply(2, 2)
	want := 6

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
