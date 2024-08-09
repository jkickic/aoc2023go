package day5

import (
	"aoc2023go/utils"
	"math"
	"strconv"
	"strings"
	"sync"
)

func FindLowestLocation(almanacPath string) int64 {
	almanac := ParseAlmanac(almanacPath)
	lowestLocation := int64(-1)
	for _, seed := range almanac.seedsPart1 {
		location := almanac.mappings.mapSeedToLocation(seed)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func FindLowestLocationPart2(almanacPath string) int64 {
	almanac := ParseAlmanac(almanacPath)
	locCh := make(chan int64)
	var wg sync.WaitGroup

	for seedRangeStart, seedRangeLength := range almanac.seedsPart2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			findLowestLocationForSeedRange(seedRangeStart, seedRangeLength, almanac, locCh)
		}()
	}

	go func() {
		wg.Wait()
		close(locCh)
	}()

	return findLowestLocationInTheChannel(locCh)
}

func findLowestLocationForSeedRange(seedRangeStart int64, seedRangeLength int64, almanac Almanac, locCh chan int64) {
	localLowest := int64(math.MaxInt64)
	for seed := seedRangeStart; seed < seedRangeStart+seedRangeLength; seed++ {
		location := almanac.mappings.mapSeedToLocation(seed)
		if location < localLowest {
			localLowest = location
		}
	}
	locCh <- localLowest
}

func findLowestLocationInTheChannel(locCh chan int64) int64 {
	lowestLocation := int64(math.MaxInt64)
	for loc := range locCh {
		if loc < lowestLocation {
			lowestLocation = loc
		}
	}
	return lowestLocation
}

func ParseAlmanac(almanacPath string) Almanac {
	lines := utils.LoadFileLines(almanacPath)

	currentKey := ""
	var seedsPart1 []int64
	seedsPart2 := make(map[int64]int64)
	almanacMap := make(map[string][]Mapping)
	for _, line := range lines {
		text, isText, numbers := utils.CaptureNonDigits(line)
		if isText {
			currentKey = strings.ReplaceAll(strings.Split(text, ":")[0], "map", "")
		}
		if currentKey == "seeds" && numbers != "" {
			seedsPart1 = utils.MapStringArrayToInt64(strings.Split(numbers, " "))
			seedsPart2 = MapSeedRanges(numbers)
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
		seedsPart1,
		seedsPart2,
		Mappings{
			[][]Mapping{
				almanacMap["seed-to-soil"],
				almanacMap["soil-to-fertilizer"],
				almanacMap["fertilizer-to-water"],
				almanacMap["water-to-light"],
				almanacMap["light-to-temperature"],
				almanacMap["temperature-to-humidity"],
				almanacMap["humidity-to-location"],
			},
		},
	}
}

func MapSeedRanges(numbers string) map[int64]int64 {
	seeds := utils.MapStringArrayToInt64(strings.Fields(numbers))
	var result = make(map[int64]int64)
	for i := 0; i < len(seeds); i += 2 {
		result[seeds[i]] = seeds[i+1]
	}
	return result
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
	sortedMappings [][]Mapping
}

func (receiver Mappings) mapSeedToLocation(seed int64) int64 {
	for _, mapping := range receiver.sortedMappings {
		seed = mapValue(mapping, seed)
	}
	return seed
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
	seedsPart1 []int64
	seedsPart2 map[int64]int64
	mappings   Mappings
}
