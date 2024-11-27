package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block

	lines := strings.Split(input, "\n")

	part1Sum := 0
	part2Sum := 0

	for _, line := range lines {

		// get game id, probably could just use the index but hey ho
		if line == "" {
			break
		}

		colon := strings.Split(line, ":")
		id := getInt(colon[0])

		hands := strings.Split(colon[1], ";")
		handPoss := make([]bool, len(hands))

		gameR := 0
		gameG := 0
		gameB := 0

		for h, hand := range hands {
			dies := strings.Split(hand, ",")
			diePoss := make([]bool, len(dies))

			for d, die := range dies {
				count := getInt(die)
				colour := strings.Split(die, " ")[2]

				if strings.Contains(colour, "red") {
					if count > gameR {
						gameR = count
					}
				} else if strings.Contains(colour, "green") {
					if count > gameG {
						gameG = count
					}
				} else if strings.Contains(colour, "blue") {
					if count > gameB {
						gameB = count
					}
				} else {
					panic("shouldn't be possible!")
				}

				diePoss[d] = diePossible(count, colour)
			}

			handPoss[h] = !slices.Contains(diePoss, false)
		}

		gamePoss := !slices.Contains(handPoss, false)
		if gamePoss {
			part1Sum += id
		}

		power := gameR * gameG * gameB
		part2Sum += power
	}

	if part2 {
		return part2Sum
	}
	return part1Sum
}

func diePossible(count int, colour string) bool {

	var (
		r = 12
		g = 13
		b = 14
	)

	if strings.Contains(colour, "red") {
		return count <= r
	} else if strings.Contains(colour, "green") {
		return count <= g
	} else if strings.Contains(colour, "blue") {
		return count <= b
	} else {
		panic("shouldn't be possible!")
	}

}

func getInt(input string) int {
	re := regexp.MustCompile("[0-9]+")
	idString := re.FindString(input)
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	return id
}
