package day9

import (
	"aoc2023go/utils"
	"strings"
)

func SolvePart1(file string) int {
	lines := utils.LoadFileLines(file)

	sumRight := 0
	for _, lineString := range lines {
		line := utils.MapStringArrayToInt(strings.Split(lineString, " "))
		_, extrapolateRight := extrapolateLine(line)
		sumRight += extrapolateRight
	}
	return sumRight
}

func SolvePart2(file string) int {
	lines := utils.LoadFileLines(file)

	sumLeft := 0
	for _, lineString := range lines {
		line := utils.MapStringArrayToInt(strings.Split(lineString, " "))
		extrapolatedLeft, _ := extrapolateLine(line)
		sumLeft += extrapolatedLeft
	}
	return sumLeft
}

func extrapolateLine(line []int) (int, int) {
	lines := generateLines(line)
	return calculateDeltas(lines)
}

func calculateDeltas(lines [][]int) (int, int) {
	extrapolateRightDelta := 0
	extrapolateLeftDelta := 0
	for i := len(lines) - 2; i >= 0; i-- {
		extrapolateRightDelta = lines[i][len(lines[i])-1]
		extrapolateLeftDelta = lines[i][0]
		if i > 0 {
			nextLine := lines[i-1]
			extrapolateRight := extrapolateRightDelta + nextLine[len(nextLine)-1]
			extrapolateLeft := nextLine[0] - extrapolateLeftDelta

			addedRight := append(nextLine, extrapolateRight)
			lines[i-1] = append([]int{extrapolateLeft}, addedRight...)
		}
	}
	return extrapolateLeftDelta, extrapolateRightDelta
}

func generateLines(line []int) [][]int {
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
	return lines
}
