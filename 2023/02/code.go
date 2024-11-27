package main

import (
	"fmt"
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

	if part2 {
		return "not implemented"
	}

	gameSum := 0

	lines := strings.Split(input, "\n")
	fmt.Println("lines:")
	fmt.Println(len(lines))

	for i, line := range lines {
		_ = i

		// get game id, probably could just use the index but hey ho
		if line == "" {
			break
		}

		colon := strings.Split(line, ":")

		id := getInt(colon[0])
		fmt.Println("game id:")
		fmt.Println(id)

		hands := strings.Split(colon[1], ";")

		handPoss := make([]bool, len(hands))

		for h, hand := range hands {
			dies := strings.Split(hand, ",")
			diePoss := make([]bool, len(dies))

			for d, die := range dies {
				count := getInt(die)
				colour := strings.Split(die, " ")[2]

				diePoss[d] = diePossible(count, colour)
				if !diePoss[d] {
					fmt.Println("die not poss")
					fmt.Println(count)
					fmt.Println(colour)
				}
			}

			handPoss[h] = !slices.Contains(diePoss, false)
		}

		fmt.Println(handPoss)

		gamePoss := !slices.Contains(handPoss, false)
		fmt.Println(gamePoss)
		if gamePoss {
			gameSum += id
		}
	}

	// solve part 1 here
	return gameSum
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
	}

	return false
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
