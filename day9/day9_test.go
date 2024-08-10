package day9

import "testing"

func TestSample(t *testing.T) {
	var count = SolvePart1("test_input.txt")
	if count != 114 {
		t.Fatalf("Expected 114, but got %d", count)
	}
}

func TestSolvePart1(t *testing.T) {
	var count = SolvePart1("input.txt")
	if count != 1772145754 {
		t.Fatalf("Expected 1772145754, but got %d", count)
	}
}

func TestSamplePart2(t *testing.T) {
	var count = SolvePart2("test_input.txt")
	if count != 2 {
		t.Fatalf("Expected 2, but got %d", count)
	}
}

func TestSolvePart2(t *testing.T) {
	var count = SolvePart2("input.txt")
	if count != 867 {
		t.Fatalf("Expected 867, but got %d", count)
	}
}
