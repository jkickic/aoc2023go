package day8

import (
	"aoc2023go/utils"
	"regexp"
)

type LR struct {
	left  string
	right string
}

func Solve(file string) int {
	lines := utils.LoadFileLines(file)

	directions := lines[0]
	directionsMap := make(map[string]LR)

	println(directions)
	re := regexp.MustCompile("[^a-zA-Z]+")
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		res := re.Split(line, -1)
		directionsMap[res[0]] = LR{res[1], res[2]}
	}

	counter := 0
	key := "AAA"
	for key != "ZZZ" {
		directionIndex := counter % len(directions)
		direction := string(directions[directionIndex])
		lr := directionsMap[key]
		if direction == "L" {
			key = lr.left
		} else {
			key = lr.right
		}
		counter++
	}
	return counter
}
