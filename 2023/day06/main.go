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
	fmt.Println("Running part 1")
	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
	fmt.Println("Running part 2")
	ans = part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func getRecordBreaks(time int, record int) int {
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
	return breaks
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	times := cast.StringSliceToIntSlice(strings.Fields(strings.Split(parsed[0], ":")[1]))
	records := cast.StringSliceToIntSlice(strings.Fields(strings.Split(parsed[1], ":")[1]))
	var newRecordCounts []int
	for idx, time := range times {
		newRecordCounts = append(newRecordCounts, getRecordBreaks(time, records[idx]))
	}
	mult := 1
	for _, count := range newRecordCounts {
		mult *= count
	}
	return mult
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	time := cast.ToInt(strings.Join(strings.Fields(strings.Split(parsed[0], ":")[1]), ""))
	record := cast.ToInt(strings.Join(strings.Fields(strings.Split(parsed[1], ":")[1]), ""))
	return getRecordBreaks(time, record)
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
