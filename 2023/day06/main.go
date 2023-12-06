package main

import (
	_ "embed"
	"flag"
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
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	times := cast.StringSliceToIntSlice(strings.Fields(strings.Split(parsed[0], ":")[1]))
	records := cast.StringSliceToIntSlice(strings.Fields(strings.Split(parsed[1], ":")[1]))
	var newRecordCounts []int

	for idx, time := range times {
		record := records[idx]
		max := time / 2
		result := record + 1
		breaks := 0
		//start in the middle, go backwards until we lose
		for i := max; result > record; i-- {
			result = (time - i) * (time - (time - i))
			if result > record {
				breaks++
			}
		}
		breaks *= 2
		if time%2 == 0 {
			breaks -= 1
		}
		newRecordCounts = append(newRecordCounts, breaks)
	}

	mult := 1
	for _, count := range newRecordCounts {
		mult *= count
	}
	return mult
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
