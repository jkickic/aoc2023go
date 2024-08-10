package day8

import "testing"

func TestSample(t *testing.T) {
	var count = SolvePart1("test_input.txt")
	if count != 6 {
		t.Fatalf("Expected 6, but got %d", count)
	}
}

func TestSolve(t *testing.T) {
	var count = SolvePart1("input.txt")
	if count != 20221 {
		t.Fatalf("Expected 20221, but got %d", count)
	}
}

func TestSamplePart2(t *testing.T) {
	var count = SolvePart2("test_input2.txt")
	if count != 6 {
		t.Fatalf("Expected 6, but got %d", count)
	}
}

func TestSolvePart2(t *testing.T) {
	var count = SolvePart2("input.txt")
	if count != 14616363770447 {
		t.Fatalf("Expected 14616363770447, but got %d", count)
	}
}
