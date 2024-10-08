package day5

import (
	"testing"
)

func TestFindLocationSampleInput(t *testing.T) {
	result := Mappings{
		[][]Mapping{{
			{50, 98, 2},
			{52, 50, 48}},
			{
				{0, 15, 37},
				{37, 52, 2},
				{39, 0, 15},
			},
			{
				{49, 53, 8},
				{0, 11, 42},
				{42, 0, 7},
				{57, 7, 4},
			},
			{
				{88, 18, 7},
				{18, 25, 70},
			},
			{
				{45, 77, 23},
				{81, 45, 19},
				{68, 64, 13},
			},
			{
				{0, 69, 1},
				{1, 0, 69},
			},
			{
				{60, 56, 37},
				{56, 93, 4},
			},
		}}.mapSeedToLocation(79)
	if result != 82 {
		t.Errorf("%d != 82", result)
	}
}

func TestParseAlmanac(t *testing.T) {
	res := ParseAlmanac("test_input.txt")
	if len(res.seedsPart1) != 4 {
		t.Errorf("%d != 4", res.seedsPart1)
	}
	if len(res.mappings.sortedMappings[0]) != 2 {
		t.Errorf("%d != 2", res.mappings.sortedMappings[0])
	}
	if len(res.mappings.sortedMappings[1]) != 3 {
		t.Errorf("%d != 2", res.mappings.sortedMappings[1])
	}
	if len(res.mappings.sortedMappings[2]) != 4 {
		t.Errorf("%d != 2", res.mappings.sortedMappings[2])
	}
	if len(res.mappings.sortedMappings[3]) != 2 {
		t.Errorf("%d != 2", res.mappings.sortedMappings[3])
	}
	if len(res.mappings.sortedMappings[4]) != 3 {
		t.Errorf("%d != 2", res.mappings.sortedMappings[4])
	}
	if len(res.mappings.sortedMappings[5]) != 2 {
		t.Errorf("%d != 2", res.mappings.sortedMappings[5])
	}
	if len(res.mappings.sortedMappings[6]) != 2 {
		t.Errorf("%d != 2", res.mappings.sortedMappings[6])
	}
}

func TestFindLowestLocation(t *testing.T) {
	res := FindLowestLocation("test_input.txt")
	if res != 35 {
		t.Errorf("%d != 35", res)
	}
}

func TestSolvePart1(t *testing.T) {
	res := FindLowestLocation("input.txt")
	if res != 265018614 {
		t.Errorf("%d != 265018614", res)
	}
}

func TestSeedRangeMapping(t *testing.T) {
	res := MapSeedRanges("79 14 55 13")
	if res[79] != 14 {
		t.Errorf("%d != 14", res[79])
	}
	if res[55] != 13 {
		t.Errorf("%d != 13", res[13])
	}
}

func TestFindLowestLocationPart2(t *testing.T) {
	res := FindLowestLocationPart2("test_input.txt")
	if res != 46 {
		t.Errorf("%d != 46", res)
	}
}

func TestSolvePart2(t *testing.T) {
	res := FindLowestLocationPart2("input.txt")
	if res != 63179500 {
		t.Errorf("%d != 63179500", res)
	}
}
