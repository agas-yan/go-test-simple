package main

import "testing"

func TestAdd2(t *testing.T) {
	if Add2(2) != 4 {
		t.Error("Expected 2 Add 2 is 4")
	}
}
