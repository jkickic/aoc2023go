package day7

import (
	"aoc2023go/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func SolvePart2(file string) int {
	lines := utils.LoadFileLines(file)
	var hands []Hand
	for _, line := range lines {
		handParts := strings.Split(line, " ")
		atoi, _ := strconv.Atoi(handParts[1])
		hands = append(hands, Hand{handParts[0], atoi})
	}

	slices.SortFunc(hands, SortByRankPart2)

	score := 0
	for index, hand := range hands {
		fmt.Printf("%s %d\n", hand.hand, hand.rankHand2())
		score += (index + 1) * hand.bid
	}
	return score
}

var SortByRankPart2 = func(p1 Hand, p2 Hand) int {
	rank1 := p1.rankHand2()
	rank2 := p2.rankHand2()
	if rank1 == rank2 {
		return SortByHighCardPart2(p1, p2)
	}
	return rank2 - rank1
}

func (hand *Hand) rankHand2() int {
	cardCount := CountCards(hand)
	var counts []int
	for card, value := range cardCount {
		if card != "J" {
			counts = append(counts, value)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	if len(counts) == 0 {
		counts = append(counts, 0)
	}
	counts[0] = counts[0] + cardCount["J"]
	return CalculateHandStrength(counts)
}

var part2Mapping = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"J": 1,
}

func mapCard2(card string) int {
	mappedValue, exists := part2Mapping[card]
	if exists {
		return mappedValue
	}
	mappedInt, err := strconv.Atoi(card)
	if err != nil || mappedInt < 2 || mappedInt > 9 {
		return 0
	}
	return mappedInt
}

var SortByHighCardPart2 = func(p1 Hand, p2 Hand) int {
	for idx, characterInt := range p1.hand {
		p1Card := mapCard2(string(characterInt))
		p2Card := mapCard2(string(p2.hand[idx]))
		if p1Card < p2Card {
			return -1
		} else if p1Card > p2Card {
			return 1
		}
	}
	return 0
}
