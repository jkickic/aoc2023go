package day7

import (
	"aoc2023go/utils"
	"slices"
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
		score += (index + 1) * hand.bid
	}
	return score
}

func (hand *Hand) rankHand2() int {
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
	var countCount = make(map[int]int)
	countCount[0] = 0
	for card, value := range cardCount {
		cc, exitst := countCount[value]
		if card != "J" {
			if exitst {
				countCount[value] = cc + 1
			} else {
				countCount[value] = 1
			}
		}
	}

	jokers := cardCount["J"]

	_, fiveOfAKind := countCount[5-jokers]
	if fiveOfAKind {
		return 1
	}
	_, fourOfAKind := countCount[4-jokers]
	if fourOfAKind {
		return 2
	}
	_, threeOfAKind := countCount[3-jokers]
	if threeOfAKind {
		twoOfAKindCount, _ := countCount[2]
		if twoOfAKindCount == 2 || jokers == 0 {
			return 3
		}
		if threeOfAKind {
			return 4
		}
	}
	pairCount, _ := countCount[2]
	twoPairs := pairCount == 2
	if twoPairs {
		return 5
	}
	_, twoOfAKind := countCount[2-jokers]
	if twoOfAKind {
		return 6
	}
	return 7
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
		}
		if p1Card > p2Card {
			return 1
		}
	}
	return 0
}

var SortByRankPart2 = func(p1 Hand, p2 Hand) int {
	if p1.rankHand2() == p2.rankHand2() {
		return SortByHighCardPart2(p1, p2)
	}
	if p1.rankHand2() > p2.rankHand2() {
		return -1
	}
	return 1
}
