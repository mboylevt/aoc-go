package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/mboylevt/aoc-go/cast"
	intcode "github.com/mboylevt/aoc-go/lib"
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

	// if part == 1 {
	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
	// } else {
	ans2 := part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans2))
	fmt.Println("Output:", ans2)
	// }
}

func part1(input string) int {
	parsed := parseInput(input)
	return intcode.RunProgram(parsed)
}

func part2(input string) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			parsed := parseInput(input)
			parsed[1] = noun
			parsed[2] = verb
			var final_val = intcode.RunProgram(parsed)
			fmt.Printf("N: %v V: %v Result: %v\n", noun, verb, final_val)
			if final_val == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, ",") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}
