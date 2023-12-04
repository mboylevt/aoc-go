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

type game struct {
	red   int
	blue  int
	green int
}

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

func gameParser(line string) map[int][]game {
	gameSplit := strings.Split(line, ": ")

	// find the game id
	gameId := cast.ToInt(gameSplit[0][5:])
	var games []game
	for _, currentGame := range strings.Split(gameSplit[1], "; ") {
		currentGameStruct := game{red: 0, blue: 0, green: 0}
		for _, color := range strings.Split(currentGame, ", ") {
			qc := strings.Split(color, " ")
			switch qc[1] {
			case "red":
				currentGameStruct.red = cast.ToInt(qc[0])
			case "green":
				currentGameStruct.green = cast.ToInt(qc[0])
			case "blue":
				currentGameStruct.blue = cast.ToInt(qc[0])
			}
		}
		games = append(games, currentGameStruct)
	}
	_ = gameId

	return map[int][]game{gameId: games}
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	sum := 0

	// set max conditions
	redMax := 12
	greenMax := 13
	blueMax := 14
	var games []map[int][]game
	for _, line := range parsed {
		games = append(games, gameParser(line))
	}
	for _, g := range games {
		possible := true
		var gameId int
		for k, _ := range g {
			gameId = k
		}
		cubeSets := g[gameId]
		for _, cubeSet := range cubeSets {
			if cubeSet.red > redMax || cubeSet.blue > blueMax || cubeSet.green > greenMax {
				possible = false
				break
			}
		}
		if possible == true {
			sum += gameId
		}
	}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	sum := 0

	// set max conditions

	var games []map[int][]game
	for _, line := range parsed {
		games = append(games, gameParser(line))
	}
	for _, g := range games {
		redMax := 0
		greenMax := 0
		blueMax := 0
		var gameId int
		for k, _ := range g {
			gameId = k
		}
		cubeSets := g[gameId]
		for _, cubeSet := range cubeSets {
			if cubeSet.red > redMax {
				redMax = cubeSet.red
			}
			if cubeSet.green > greenMax {
				greenMax = cubeSet.green
			}
			if cubeSet.blue > blueMax {
				blueMax = cubeSet.blue
			}
		}
		sum += redMax * blueMax * greenMax
	}
	return sum
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)

	return ans
}
