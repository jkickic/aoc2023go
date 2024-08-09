package day7

import (
	"fmt"
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

func TestJokerRanks(t *testing.T) {
	expectedHandResults := map[Hand]int{
		{"JJJJJ", 0}: 1,
		{"JJJJK", 0}: 1,
		{"JJJKK", 0}: 1,
		{"JJKKK", 0}: 1,
		{"JJJKQ", 0}: 2,
		{"JJKKQ", 0}: 2,
		{"JKKKQ", 0}: 2,
		{"KKKQQ", 0}: 3,
		{"JJKTQ", 0}: 4,
		{"J8477", 0}: 4,
		{"KQKTQ", 0}: 5,
		{"1QKTQ", 0}: 6,
		{"J1234", 0}: 6,
	}
	for hand, expectedScore := range expectedHandResults {
		t.Run(fmt.Sprintf("%s should score %d", hand.hand, expectedScore), func(t *testing.T) {
			score := hand.rankHand2()
			if score != expectedScore {
				t.Errorf("%s expected %d not %d", hand.hand, expectedScore, score)
			}
		})
	}

}

func TestSampleHandPart2(t *testing.T) {
	score := SolvePart2("test_input.txt")
	if score != 5905 {
		t.Errorf("Incorrect score %d, expected %d", score, 5905)
	}
}

func TestPart2(t *testing.T) {
	score := SolvePart2("input.txt")
	if score != 248747492 {
		t.Errorf("Incorrect score %d, expected %d", score, 248747492)
	}
}
