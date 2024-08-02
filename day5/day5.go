package day5

import (
	"aoc2023go/utils"
	"strconv"
	"strings"
)

func FindLowestLocation(almanacPath string) int64 {
	almanac := ParseAlmanac(almanacPath)
	lowestLocation := int64(-1)
	for _, seed := range almanac.seeds {
		location := almanac.mappings.mapSeedToLocation(seed)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func ParseAlmanac(almanacPath string) Almanac {
	lines := utils.LoadFileLines(almanacPath)

	currentKey := ""
	seeds := []int64{}
	almanacMap := make(map[string][]Mapping)
	for _, line := range lines {
		text, isText, numbers := utils.CaptureNonDigits(line)
		if isText {
			currentKey = strings.ReplaceAll(strings.Split(text, ":")[0], "map", "")
		}
		if currentKey == "seeds" && numbers != "" {
			seeds = utils.MapStringArrayToInt64(strings.Split(numbers, " "))
		} else {
			if numbers != "" {
				if almanacMap[currentKey] == nil {
					almanacMap[currentKey] = []Mapping{}
				}
				almanacMap[currentKey] = append(almanacMap[currentKey], convertNumbersToMapping(numbers))
			}
		}
	}
	return Almanac{
		seeds,
		Mappings{
			almanacMap["seed-to-soil"],
			almanacMap["soil-to-fertilizer"],
			almanacMap["fertilizer-to-water"],
			almanacMap["water-to-light"],
			almanacMap["light-to-temperature"],
			almanacMap["temperature-to-humidity"],
			almanacMap["humidity-to-location"],
		},
	}
}

func convertNumbersToMapping(numbers string) Mapping {
	splitNumbers := strings.Split(numbers, " ")
	drs, _ := strconv.ParseInt(splitNumbers[0], 10, 64)
	srs, _ := strconv.ParseInt(splitNumbers[1], 10, 64)
	rl, _ := strconv.ParseInt(splitNumbers[2], 10, 64)
	return Mapping{
		drs,
		srs,
		rl,
	}
}

type Mappings struct {
	seedToSoil            []Mapping
	soilToFertilizer      []Mapping
	fertilizerToWater     []Mapping
	waterToLight          []Mapping
	lightToTemperature    []Mapping
	temperatureToHumidity []Mapping
	humidityToLocation    []Mapping
}

func (receiver Mappings) mapSeedToLocation(seed int64) int64 {
	soil := mapValue(receiver.seedToSoil, seed)
	fertilizer := mapValue(receiver.soilToFertilizer, soil)
	water := mapValue(receiver.fertilizerToWater, fertilizer)
	light := mapValue(receiver.waterToLight, water)
	temperature := mapValue(receiver.lightToTemperature, light)
	humidity := mapValue(receiver.temperatureToHumidity, temperature)
	return mapValue(receiver.humidityToLocation, humidity)
}

func mapValue(mappings []Mapping, value int64) int64 {
	for _, mapping := range mappings {
		if value >= mapping.sourceRangeStart && value < mapping.sourceRangeStart+mapping.rangeLength {
			delta := value - mapping.sourceRangeStart
			return mapping.destinationRangeStart + delta
		}
	}
	return value
}

type Mapping struct {
	destinationRangeStart int64
	sourceRangeStart      int64
	rangeLength           int64
}

type Almanac struct {
	seeds    []int64
	mappings Mappings
}
