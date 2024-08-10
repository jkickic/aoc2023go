package day8

import (
	"aoc2023go/utils"
	"regexp"
	"strings"
)

type LR struct {
	left  string
	right string
}

func SolvePart1(file string) int {
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

func SolvePart2(file string) int {
	lines := utils.LoadFileLines(file)

	directions := lines[0]
	directionsMap := make(map[string]LR)
	var startingKeys []string

	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		res := re.Split(line, -1)
		directionsMap[res[0]] = LR{res[1], res[2]}
		if strings.HasSuffix(res[0], "A") {
			startingKeys = append(startingKeys, res[0])
		}
	}

	results := []int{}

	for _, key := range startingKeys {
		counter := 0
		for !strings.HasSuffix(key, "Z") {
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
		results = append(results, counter)
	}

	return lcmArray(results)
}

// COPILOT CHEATS:
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmArray(arr []int) int {
	result := arr[0]
	for _, num := range arr[1:] {
		result = lcm(result, num)
	}
	return result
}
