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
	flag.IntVar(&part, "part", 2, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	// input := "14\n1969\n100756"
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
	var sumFuel int = 0
	for _, v := range parsed {
		sumFuel += calcSingleFuel(v)
	}
	return sumFuel
}

func part2(input string) int {
	parsed := parseInput(input)
	var sumFuel int = 0
	for _, v := range parsed {
		sumFuel += recursiveFuel(v)
	}
	return sumFuel
}

func recursiveFuel(input int) int {
	var fuelReq = calcSingleFuel(input)
	var runningTotal = fuelReq
	for fuelReq > 6 { //this ensures that the final result will be greater than 0
		var newFuelReq = calcSingleFuel(fuelReq)
		runningTotal += newFuelReq
		fuelReq = newFuelReq
	}
	return runningTotal
}

func calcSingleFuel(input int) int {
	return (input / 3) - 2
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}
