package day6

import "testing"

func TestFindingDistances(t *testing.T) {
	var res = findWinningDistances(TimeDistance{15, 40})
	if res != 8 {
		t.Fatalf(`Wrong result, expected 8 got %d`, res)
	}
}

func TestSolvePart1(t *testing.T) {
	var total = solvePart1([]TimeDistance{
		{47, 207},
		{84, 1394},
		{74, 1209},
		{67, 1014},
	})
	if total != 741000 {
		t.Fatalf(`Wrong result, expected 741000 got %d`, total)
	}
}

func TestSolvePart2(t *testing.T) {
	var total = findWinningDistances(TimeDistance{47847467, 207139412091014})
	if total != 38220708 {
		t.Fatalf(`Wrong result, expected 38220708 got %d`, total)
	}
}
