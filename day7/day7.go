package day7

import (
	"aoc2023go/utils"
	"fmt"
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
		cardCounts[0] += cardCount["J"]
	}
	return calculateHandStrength(cardCounts)
}

func calculateHandStrength(counts []int) int {
	handStrengths := map[string]int{
		"[5]":         1,
		"[4 1]":       2,
		"[3 2]":       3,
		"[3 1 1]":     4,
		"[2 2 1]":     5,
		"[2 1 1 1]":   6,
		"[1 1 1 1 1]": 7,
	}
	strength, _ := handStrengths[fmt.Sprint(counts)]
	return strength
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
	mappedInt, _ := strconv.Atoi(card)
	return utils.GetOr(mapping, card, mappedInt)
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
