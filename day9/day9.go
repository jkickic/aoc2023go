package day9

import (
	"aoc2023go/utils"
	"strings"
)

func SolvePart1(file string) int {
	lines := utils.LoadFileLines(file)

	sum := 0
	for _, lineString := range lines {
		line := utils.MapStringArrayToInt(strings.Split(lineString, " "))
		sum += extrapolateLine(line)
	}
	return sum
}

func extrapolateLine(line []int) int {
	allZeros := false
	lines := [][]int{line}
	currentLine := line
	for !allZeros {
		var nextLine []int
		for index, value := range currentLine[1:] {
			nextLine = append(nextLine, value-currentLine[index])
		}
		lines = append(lines, nextLine)
		currentLine = nextLine
		allZeros = utils.AllMatch(nextLine, func(value int) bool {
			return value == 0
		})
	}
	latestValue := 0
	for i := len(lines) - 2; i >= 0; i-- {
		latestValue = lines[i][len(lines[i])-1]
		if i > 0 {
			nextLine := lines[i-1]
			lines[i-1] = append(nextLine, latestValue+nextLine[len(nextLine)-1])
		}
	}
	return latestValue
}
