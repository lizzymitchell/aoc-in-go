package main

import (
	"slices"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/lizzymitchell/aoc-in-go/utils"
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
	// parse input
	rows := strings.Split(input, "\n")

	left := make([]int, len(rows))
	right := make([]int, len(rows))

	for r, row := range rows {
		if row == "" {
			break
		}
		pair := strings.Split(row, "  ")

		left[r] = utils.ParseInt(pair[0])
		right[r] = utils.ParseInt(pair[1])
	}

	sortedLeft := make([]int, len(left))
	copy(sortedLeft, left)
	slices.Sort(sortedLeft)

	sortedRight := make([]int, len(right))
	copy(sortedRight, right)
	slices.Sort(sortedRight)

	// adding up each number in the left list
	// after multiplying it by the number of times
	// that number appears in the right list.

	part1Result := 0

	for i, _ := range sortedLeft {
		diff := sortedLeft[i] - sortedRight[i]
		absDiff := max(diff, -diff)
		part1Result += absDiff
	}

	part2Result := 0

	for _, lefty := range left {
		count := 0
		for _, righty := range right {
			if righty == lefty {
				count++
			}
		}

		similarity := lefty * count
		part2Result += similarity
	}

	if part2 {
		return part2Result
	}

	// solve part 1 here
	return part1Result
}
