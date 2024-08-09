package day7

import (
	"testing"
)

func TestSampleHand(t *testing.T) {
	score := SolvePart1("test_input.txt")
	if score != 6440 {
		t.Errorf("Incorrect score %d, expected %d", score, 6440)
	}
}

func TestPart1(t *testing.T) {
	score := SolvePart1("input.txt")
	if score != 247815719 {
		t.Errorf("Incorrect score %d, expected %d", score, 247815719)
	}
}
