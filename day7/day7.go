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

func SolvePart1(file string) int {
	lines := utils.LoadFileLines(file)
	var hands []Hand
	for _, line := range lines {
		handParts := strings.Split(string(line), " ")
		atoi, _ := strconv.Atoi(handParts[1])
		hands = append(hands, Hand{handParts[0], atoi})
	}

	slices.SortFunc(hands, SortByRank)

	score := 0
	for index, hand := range hands {
		score += (index + 1) * hand.bid
	}
	return score
}

func (hand *Hand) rankHand() int {
	var cardCount = CountCards(hand)
	var counts []int
	for _, value := range cardCount {
		counts = append(counts, value)
	}

	return CalculateHandStrength(counts)
}

func CountCards(hand *Hand) map[string]int {
	var cardCount = make(map[string]int)
	for _, cardInt := range hand.hand {
		var card = string(cardInt)
		val, exists := cardCount[card]
		if exists {
			cardCount[card] = val + 1
		} else {
			cardCount[card] = 1
		}
	}
	return cardCount
}

func CalculateHandStrength(counts []int) int {
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

var part1Mapping = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
}

func mapCard(card string) int {
	mappedValue, exists := part1Mapping[card]
	if exists {
		return mappedValue
	}
	mappedInt, err := strconv.Atoi(card)
	if err != nil || mappedInt < 2 || mappedInt > 9 {
		return 0
	}
	return mappedInt
}

var SortByHighCard = func(p1 Hand, p2 Hand) int {
	for idx, characterInt := range p1.hand {
		p1Card := mapCard(string(characterInt))
		p2Card := mapCard(string(p2.hand[idx]))
		if p1Card < p2Card {
			return -1
		}
		if p1Card > p2Card {
			return 1
		}
	}
	return 0
}

var SortByRank = func(p1 Hand, p2 Hand) int {
	if p1.rankHand() == p2.rankHand() {
		return SortByHighCard(p1, p2)
	}
	if p1.rankHand() > p2.rankHand() {
		return -1
	}
	return 1
}
