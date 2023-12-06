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
	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)

	ans = part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func getSeeds(input string) []int {
	seedStr := strings.Split(input, ": ")[1]
	var seedSlice []int
	for _, str := range strings.Fields(seedStr) {
		seedSlice = append(seedSlice, cast.ToInt(str))
	}
	return seedSlice
}

func getBlocks(input []string) [][]int {
	var blocks [][]int
	for _, rule := range input {
		for _, bString := range strings.Split(rule, "\n")[1:] {
			var block []int
			bFields := strings.Fields(bString)
			for _, field := range bFields {
				block = append(block, cast.ToInt(field))
			}
			blocks = append(blocks, block)
		}
	}

	return blocks
}

func getMappedNumber(block []int, seed int) int {
	dr := block[0]
	sr := block[1]
	rl := block[2]
	if sr <= seed && seed <= sr+rl {
		return dr + seed - sr
	}
	return seed
}

func part1(input string) int {
	parsed := parseInput(input)

	seeds := getSeeds(parsed[0])
	blocks := getBlocks(parsed[1:])
	// _ = seeds
	// _ = blocks
	for _, block := range blocks {
		for idx, seed := range seeds {
			seeds[idx] = getMappedNumber(block, seed)
		}
	}

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n\n") {
		ans = append(ans, line)
	}
	return ans
}
