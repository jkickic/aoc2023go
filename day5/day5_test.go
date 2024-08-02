package day5

import (
	"testing"
)

func TestFindLocationSampleInput(t *testing.T) {
	result := Mappings{
		seedToSoil: []Mapping{
			{50, 98, 2},
			{52, 50, 48}},
		soilToFertilizer: []Mapping{
			{0, 15, 37},
			{37, 52, 2},
			{39, 0, 15},
		},
		fertilizerToWater: []Mapping{
			{49, 53, 8},
			{0, 11, 42},
			{42, 0, 7},
			{57, 7, 4},
		},
		waterToLight: []Mapping{
			{88, 18, 7},
			{18, 25, 70},
		},
		lightToTemperature: []Mapping{
			{45, 77, 23},
			{81, 45, 19},
			{68, 64, 13},
		},
		temperatureToHumidity: []Mapping{
			{0, 69, 1},
			{1, 0, 69},
		},
		humidityToLocation: []Mapping{
			{60, 56, 37},
			{56, 93, 4},
		},
	}.mapSeedToLocation(79)
	if result != 82 {
		t.Errorf("%d != 82", result)
	}
}
