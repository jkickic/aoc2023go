package day4

import "testing"

func TestSampleInputPart1(t *testing.T) {
	result := CalculatePoints("test_input.txt")
	if result != 13 {
		t.Errorf("%d != 13", result)
	}
}

func TestDay4InputPart1(t *testing.T) {
	result := CalculatePoints("input.txt")
	if result != 22193 {
		t.Errorf("%d != 22193", result)
	}
}

func TestSampleInputPart2(t *testing.T) {
	result := CalculateTotalCardCount("test_input.txt")
	if result != 30 {
		t.Errorf("%d != 30", result)
	}
}

func TestDay4InputPart2(t *testing.T) {
	result := CalculateTotalCardCount("input.txt")
	if result != 5625994 {
		t.Errorf("%d != 5625994", result)
	}
}
