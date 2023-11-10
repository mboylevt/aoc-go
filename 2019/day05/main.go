package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/mboylevt/aoc-go/cast"
	intcode "github.com/mboylevt/aoc-go/lib"
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
	// fmt.Println("Running part 1")
	// ans := part1(input)
	// util.CopyToClipboard(fmt.Sprintf("%v", ans))

	fmt.Println("Running part 2")
	part2(input)

}

func part1(input string) int {
	parsed := parseInput(input)
	intcode.RunProgram(parsed)
	return 0
}

func part2(input string) int {
	parsed := parseInput(input)
	intcode.RunProgram(parsed)
	return 0
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, ",") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}
