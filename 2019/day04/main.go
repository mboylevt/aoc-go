package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/mboylevt/aoc-go/cast"
	"github.com/mboylevt/aoc-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {

	// ans := part1(input)
	// util.CopyToClipboard(fmt.Sprintf("%v", ans))
	// fmt.Println("Output1:", ans)
	ans := part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output2:", ans)
}

func part1(input string) int {
	parsed := parseInput(input)
	start := parsed[0]
	end := parsed[1]
	validPWCount := 0
	for i := start; i <= end; i++ {
		if evaluateEntry(cast.ToString(i), 1) {
			validPWCount++
		}
	}

	return validPWCount
}

func part2(input string) int {
	parsed := parseInput(input)
	start := parsed[0]
	end := parsed[1]
	validPWCount := 0
	for i := start; i <= end; i++ {
		if evaluateEntry(cast.ToString(i), 2) {
			fmt.Printf("Valid: %v\n", i)
			validPWCount++
		}
	}

	return validPWCount
}

func evaluateEntry(entry string, part int) bool {
	adjacent := false
	// fmt.Printf("Evaluating %v\n", entry)
	pairIdx := -1
	for idx, c := range entry[:len(entry)-1] {
		current := cast.ToString(c)
		next := cast.ToString(entry[idx+1])
		if current == next {
			adjacent = true
			if pairIdx == -1 {
				pairIdx = idx
			}

			// fmt.Printf("\tFound adjacent numbers: %v at index %v,%v\n", current, idx, idx+1)
			if part == 2 {
				if idx-1 >= 0 {
					prior := cast.ToString(entry[idx-1])
					if prior == next && pairIdx == idx-1 {
						// fmt.Printf("\tFailure - group too large: Entry %v, %v at index %v\n", entry, prior, idx-1)
						adjacent = false
					}
				}
			}
		}
		if next < current {
			// fmt.Printf("\tFailure: %v descending when compared to %v\n", next, current)
			return false
		}
	}
	// fmt.Printf("\tEvauation complete for %v: result %v\n", entry, adjacent)
	return adjacent
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "-") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}
