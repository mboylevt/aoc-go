package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/mboylevt/aoc-go/cast"
	"github.com/mboylevt/aoc-go/lib"
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

func getBlocks(input []string) [][][]int {
	var blocks [][][]int
	for _, rule := range input {
		var block [][]int
		for _, bString := range strings.Split(rule, "\n")[1:] {
			var rng []int
			bFields := strings.Fields(bString)
			for _, field := range bFields {
				rng = append(rng, cast.ToInt(field))
			}
			block = append(block, rng)
		}
		blocks = append(blocks, block)
	}

	return blocks
}

func getMappedNumber(block [][]int, seed int) int {
	for _, rng := range block {
		dr := rng[0]
		sr := rng[1]
		rl := rng[2]
		if sr <= seed && seed <= sr+rl {
			return dr + seed - sr
		}
	}
	return seed
}

func identifySeeds(input string) []int {
	var seedSlice []int
	var seeds []int
	seedStr := strings.Split(input, ": ")[1]
	for _, str := range strings.Fields(seedStr) {
		seedSlice = append(seedSlice, cast.ToInt(str))
	}
	for i := 0; i < len(seedSlice); i += 2 {
		start := seedSlice[i]
		size := seedSlice[i+1]
		for j := 0; j < size; j++ {
			seeds = append(seeds, start+j)
		}
	}
	return seeds
}

func part1(input string) int {
	parsed := parseInput(input)
	seeds := getSeeds(parsed[0])
	blocks := getBlocks(parsed[1:])

	for _, block := range blocks {
		for idx, seed := range seeds {
			seeds[idx] = getMappedNumber(block, seed)
		}
	}

	return lib.FindMinimum(seeds)
}

func part2(input string) int {
	parsed := parseInput(input)
	seeds := identifySeeds(parsed[0])
	blocks := getBlocks(parsed[1:])

	for _, block := range blocks {
		for idx, seed := range seeds {
			seeds[idx] = getMappedNumber(block, seed)
		}
	}

	return lib.FindMinimum(seeds)
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n\n") {
		ans = append(ans, line)
	}
	return ans
}
