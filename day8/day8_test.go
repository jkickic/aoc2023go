package day8

import "testing"

func TestSample(t *testing.T) {
	var count = Solve("test_input.txt")
	if count != 6 {
		t.Fatalf("Expected 6, but got %d", count)
	}
}

func TestSolve(t *testing.T) {
	var count = Solve("input.txt")
	if count != 20221 {
		t.Fatalf("Expected 20221, but got %d", count)
	}
}
