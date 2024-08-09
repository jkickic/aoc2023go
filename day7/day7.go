package day7

import (
	"aoc2023go/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand string
	bid  int
}

func Solve(file string, jokers bool) int {
	lines := utils.LoadFileLines(file)
	var hands []Hand
	for _, line := range lines {
		handParts := strings.Split(line, " ")
		atoi, _ := strconv.Atoi(handParts[1])
		hands = append(hands, Hand{handParts[0], atoi})
	}

	slices.SortFunc(hands, createHandSorter(jokers))

	score := 0
	for index, hand := range hands {
		score += (index + 1) * hand.bid
	}
	return score
}

func createHandSorter(jokers bool) func(Hand, Hand) int {
	return func(p1 Hand, p2 Hand) int {
		strength1 := p1.handStrength(jokers)
		strength2 := p2.handStrength(jokers)
		if strength1 == strength2 {
			return sortByHighCard(jokers, p1, p2)
		}
		return strength2 - strength1
	}
}

func (hand *Hand) handStrength(jokers bool) int {
	var cardCount = make(map[string]int)
	for _, cardInt := range hand.hand {
		cardCount[string(cardInt)]++
	}
	var cardCounts []int
	for card, value := range cardCount {
		if !jokers || (card != "J") {
			cardCounts = append(cardCounts, value)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cardCounts)))

	if len(cardCounts) == 0 {
		cardCounts = append(cardCounts, 0)
	}
	if jokers {
		cardCounts[0] = cardCounts[0] + cardCount["J"]
	}
	return calculateHandStrength(cardCounts)
}

func calculateHandStrength(counts []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	if utils.ArraysEqual(counts, []int{5}) {
		return 1
	}
	if utils.ArraysEqual(counts, []int{4, 1}) {
		return 2
	}
	if utils.ArraysEqual(counts, []int{3, 2}) {
		return 3
	}
	if utils.ArraysEqual(counts, []int{3, 1, 1}) {
		return 4
	}
	if utils.ArraysEqual(counts, []int{2, 2, 1}) {
		return 5
	}
	if utils.ArraysEqual(counts, []int{2, 1, 1, 1}) {
		return 6
	}
	return 7
}

var sortByHighCard = func(jokers bool, p1 Hand, p2 Hand) int {
	for idx, characterInt := range p1.hand {
		p1Card := mapCard(jokers, string(characterInt))
		p2Card := mapCard(jokers, string(p2.hand[idx]))
		if p1Card < p2Card {
			return -1
		} else if p1Card > p2Card {
			return 1
		}
	}
	return 0
}

func mapCard(jokers bool, card string) int {
	mapping := cardStrengthMapping(jokers)
	mappedValue, exists := mapping[card]
	if exists {
		return mappedValue
	}
	mappedInt, err := strconv.Atoi(card)
	if err != nil || mappedInt < 2 || mappedInt > 9 {
		return 0
	}
	return mappedInt
}

func cardStrengthMapping(jokers bool) map[string]int {
	var mapping = map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
	}
	if jokers {
		mapping["J"] = 1
	}
	return mapping
}
