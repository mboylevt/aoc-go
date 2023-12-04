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

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	totalScore := 0
	for _, line := range parsed {
		cardScore := 0
		winnerMap := make(map[int]struct{})
		// var winnerMap map[int]bool
		lineSplit := strings.Split(line, ": ")
		// find the card id
		cardId := cast.ToInt(strings.Fields(lineSplit[0])[1])
		_ = cardId

		// do more input processing
		lineSplit = strings.Split(lineSplit[1], " | ")
		winnerStr := lineSplit[0]
		numbersStr := lineSplit[1]

		// populate winners
		for _, w := range strings.Fields(winnerStr) {
			winnerMap[cast.ToInt(w)] = struct{}{}
		}

		// check for winners in numbers
		for _, n := range strings.Fields(numbersStr) {
			if _, ok := winnerMap[cast.ToInt(n)]; ok {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore = cardScore + cardScore
				}
			}
		}
		totalScore += cardScore
	}
	return totalScore
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	cards := make(map[int]string)
	cardCount := make(map[int]int)

	// parse cards + get
	for _, line := range parsed {
		lineSplit := strings.Split(line, ": ")
		cardId := cast.ToInt(strings.Fields(lineSplit[0])[1])
		cards[cardId] = lineSplit[1]
		cardCount[cardId] = 1
	}

	// do actual loop
	for idx := 1; idx < len(cards)+1; idx++ {
		cardId := idx
		details := cards[cardId]
		winnerMap := make(map[int]struct{})
		lineSplit := strings.Split(details, " | ")
		winnerStr := lineSplit[0]
		numbersStr := lineSplit[1]
		winnerCount := 0

		// populate winners
		for _, w := range strings.Fields(winnerStr) {
			winnerMap[cast.ToInt(w)] = struct{}{}
		}

		// check for winners in numbers
		for _, n := range strings.Fields(numbersStr) {
			if _, ok := winnerMap[cast.ToInt(n)]; ok {
				winnerCount += 1
			}
		}

		// Add card copies based on winner count
		for i := cardId + 1; i <= cardId+winnerCount; i++ {
			cardCount[i] += cardCount[cardId]
		}
	}

	// calcualte total
	total := 0
	for _, v := range cardCount {
		total += v
	}
	return total
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
