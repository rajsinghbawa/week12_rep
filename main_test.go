package main

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {

	if !IsLeapYear(2000) {
		t.Errorf("Expected 2000 to be a leap year")
	}

	if IsLeapYear(1900) {
		t.Errorf("Expected 1900 not to be a leap year")
	}

	if IsLeapYear(2024) {
		t.Errorf("Expected 2024 to be a leap year")
	}

	// Test case: Year not divisible by 4
	if IsLeapYear(2019) {
		t.Errorf("Expected 2019 not to be a leap year")
	}
}
