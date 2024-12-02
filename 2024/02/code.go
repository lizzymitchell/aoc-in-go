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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	part1Result := 0
Report:
	for r, report := range strings.Split(input, "\n") {
		_ = r

		if report == "" {
			continue Report
		}

		levels := strings.Split(report, " ")

		diff := make([]int, len(levels)-1)
		positive := make([]bool, len(levels)-1)
		for i := 0; i < len(diff); i++ {
			diff[i] = utils.ParseInt(levels[i]) - utils.ParseInt(levels[i+1])
			if diff[i] == 0 || diff[i] > 3 || diff[i] < -3 {
				continue Report // not safe
			}
			positive[i] = diff[i] > 0
		}

		ascending := slices.Contains(positive, true)
		descending := slices.Contains(positive, false)

		if ascending && descending {
			continue Report // not safe
		}

		part1Result++
	}
	// solve part 1 here
	return part1Result
}
