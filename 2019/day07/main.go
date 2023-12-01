package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/mboylevt/aoc-go/cast"
	mafs "github.com/mboylevt/aoc-go/lib"
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

	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
	ans = part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)

}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	permutations := mafs.Permutations([]int{0, 1, 2, 3, 4})
	max_thruster := 0
	for _, perm := range(permutations) {
		input := 0
		for _, phase := perm {
			
		}
	}
	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, ",") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}
