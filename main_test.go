package main

import "testing"

func TestAdd2(t *testing.T) {
	if Add2(2) != 4 {
		t.Error("Expected 2 Add 2 is 4")
	}
}

func TestTableAdd2(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{0, 2},
		{-5, -3},
		{1000, 1002},
	}
	// Test for
	for _, test := range tests {
		if output := Add2(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
