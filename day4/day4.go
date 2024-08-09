package day4

import (
	"aoc2023go/utils"
	"math"
	"strings"
)

func CalculatePoints(cardPath string) int {
	lines := utils.LoadFileLines(cardPath)
	points := 0
	for _, line := range lines {
		scratchcard := createScratchcard(line)

		hitCount := hitCount(scratchcard)
		if hitCount > 0 {
			points += int(math.Pow(2, float64(hitCount-1)))
		}
	}
	return points
}

func CalculateTotalCardCount(cardPath string) int {
	lines := utils.LoadFileLines(cardPath)
	scratchcards := make([]scratchcard, 0)

	for _, line := range lines {
		scratchcards = append(scratchcards, createScratchcard(line))
	}

	for idx, scratchcard := range scratchcards {
		hitCount := hitCount(scratchcard)

		if hitCount > 0 {
			for i := idx + 1; i <= idx+hitCount && idx < len(scratchcards); i++ {
				scratchcards[i].cardCount += scratchcard.cardCount
			}
		}
	}

	totalCardCount := 0
	for _, scratchcard := range scratchcards {
		totalCardCount += scratchcard.cardCount
	}
	return totalCardCount
}

func createScratchcard(line string) scratchcard {
	splitLine := strings.Split(line, ":")
	cardNumbers := splitLine[1]
	cardNumbersSplit := strings.Split(cardNumbers, "|")
	winningNumbers := utils.MapStringArrayToInt(utils.Filter(strings.Split(cardNumbersSplit[0], " "), utils.NotEmpty))
	chosenNumbers := utils.MapStringArrayToInt(utils.Filter(strings.Split(cardNumbersSplit[1], " "), utils.NotEmpty))

	scratchcard := scratchcard{winningNumbers, chosenNumbers, 1}
	return scratchcard
}

type scratchcard struct {
	winningNumbers []int
	chosenNumbers  []int
	cardCount      int
}

func hitCount(scratchcard scratchcard) int {
	array2Map := make(map[int]bool)
	for _, num := range scratchcard.chosenNumbers {
		array2Map[num] = true
	}
	count := 0
	for _, num := range scratchcard.winningNumbers {
		if array2Map[num] {
			count++
		}
	}
	return count
}
