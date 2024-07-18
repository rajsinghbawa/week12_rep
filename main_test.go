package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned an error: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Divide(10, 2) = %f; want %f", result, expected)
	}

}
