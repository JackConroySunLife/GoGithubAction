// integers_test.go
package main

import "testing"

func TestMultiply(t *testing.T) {
	got := Multiply(2, 5)
	want := 10

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
