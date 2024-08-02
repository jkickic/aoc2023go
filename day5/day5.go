package day5

type Mappings struct {
	seedToSoil            []Mapping
	soilToFertilizer      []Mapping
	fertilizerToWater     []Mapping
	waterToLight          []Mapping
	lightToTemperature    []Mapping
	temperatureToHumidity []Mapping
	humidityToLocation    []Mapping
}

func (receiver Mappings) mapSeedToLocation(seed int) int {
	soil := mapValue(receiver.seedToSoil, seed)
	fertilizer := mapValue(receiver.soilToFertilizer, soil)
	water := mapValue(receiver.fertilizerToWater, fertilizer)
	light := mapValue(receiver.waterToLight, water)
	temperature := mapValue(receiver.lightToTemperature, light)
	humidity := mapValue(receiver.temperatureToHumidity, temperature)
	return mapValue(receiver.humidityToLocation, humidity)
}

func mapValue(mappings []Mapping, value int) int {
	for _, mapping := range mappings {
		if value >= mapping.sourceRangeStart && value < mapping.sourceRangeStart+mapping.rangeLength {
			delta := value - mapping.sourceRangeStart
			return mapping.destinationRangeStart + delta
		}
	}
	return value
}

type Mapping struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type Almanac struct {
	seeds    []int
	mappings Mappings
}
