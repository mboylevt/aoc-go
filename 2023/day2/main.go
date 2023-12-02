package main

import (
	_ "embed"
	"fmt"
	"strings"

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

func gameParser(line string) string {
	gameSplit := strings.Split(line, ":")

	// find the game id
	gameId := gameSplit[0][5:]
	for _, game := range strings.Split(gameSplit[1], ";") {
		fmt.Printf(game)
	}
	_ = gameId

	return "matt"
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	sum := 0
	// gameRx, _ := regexp.Compile("^Game (\\d+):")
	for _, line := range parsed {
		gameParser(line)

	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	sum := 0

	return sum
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)

	return ans
}
