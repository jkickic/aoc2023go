package day6

type TimeDistance struct {
	maxTime  int
	distance int
}

func solvePart1(settings []TimeDistance) int {
	var total = 0
	for _, setting := range settings {
		var winningTimesCount = findWinningDistances(setting)
		if total == 0 {
			total = winningTimesCount
		} else {
			total = total * winningTimesCount
		}
	}
	return total
}

func findWinningDistances(timeDist TimeDistance) int {
	var counter = 0
	for time := 0; time < timeDist.maxTime; time++ {
		var newDistance = (timeDist.maxTime - time) * time
		if newDistance > timeDist.distance {
			counter++
		}
	}

	return counter
}
