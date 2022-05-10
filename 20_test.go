package main

import "testing"

func TestIsValidString(t *testing.T) {
	input1 := "()"
	if !isValid(input1) {
		t.Errorf("output must be true")
	}

	input2 := "()[]{}"
	if !isValid(input2) {
		t.Errorf("output must be true")
	}

	input3 := "(]"
	if isValid(input3) {
		t.Errorf("output must be false")
	}

	input4 := "(])"
	if isValid(input4) {
		t.Errorf("output must be false")
	}

	input5 := "([}}])"
	if isValid(input5) {
		t.Errorf("output must be false")
	}
}
