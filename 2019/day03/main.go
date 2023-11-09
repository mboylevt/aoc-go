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

type move struct {
	dir  string
	dist int
}

type point struct {
	x int
	y int
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

	ans2 := part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans2))
	fmt.Println("Output:", ans2)

}

func part1(input string) int {
	var wire1Moves, wire2Moves []move = parseInput(input)
	wire1 := populateSet(wire1Moves)
	wire2 := populateSet(wire2Moves)
	closest_distance := 1000000000
	var intersections []point

	for k, _ := range wire1 {
		if (point{0, 0} == k) {
			_ = 0
		} else {
			if _, ok := wire2[k]; ok {
				intersections = append(intersections, k)
				if mafs.AbsInt(k.x)+mafs.AbsInt(k.y) < closest_distance {
					closest_distance = mafs.AbsInt(k.x) + mafs.AbsInt(k.y)
				}
			}
		}
	}
	return closest_distance
}

func part2(input string) int {
	var wire1Moves, wire2Moves []move = parseInput(input)
	wire1 := populateSet(wire1Moves)
	wire2 := populateSet(wire2Moves)
	stepcount := 10000000000
	for k, _ := range wire1 {
		if (point{0, 0} == k) {
			_ = 0
		} else {
			if _, ok := wire2[k]; ok {
				intersection_steps := wire1[k] + wire2[k]
				if intersection_steps < stepcount {
					stepcount = intersection_steps
				}
			}
		}
	}
	return stepcount
}

func populateSet(wire []move) map[point]int {
	m := make(map[point]int)
	var x, y, steps int = 0, 0, 0
	m[point{x, y}] = steps
	for _, mv := range wire {
		switch mv.dir {
		case "R":
			for i := x + 1; i <= x+mv.dist; i++ {
				steps++
				m[point{i, y}] = steps
			}
			x = x + mv.dist
		case "L":
			for i := x - 1; i >= x-mv.dist; i-- {
				steps++
				m[point{i, y}] = steps
			}
			x = x - mv.dist
		case "U":
			for i := y + 1; i <= y+mv.dist; i++ {
				steps++
				m[point{x, i}] = steps
			}
			y = y + mv.dist
		case "D":
			for i := y - 1; i >= y-mv.dist; i-- {
				steps++
				m[point{x, i}] = steps
			}
			y = y - mv.dist
		}
		_ = 1
	}
	return m
}

func parseInput(input string) (wire1 []move, wire2 []move) {
	var list1, list2 string = strings.Split(input, "\n")[0], strings.Split(input, "\n")[1]
	_ = list2
	for _, mv := range strings.Split(list1, ",") {
		var dir = cast.ToString(mv[0])
		var dist = cast.ToInt(mv[1:])
		wire1 = append(wire1, move{dir: dir, dist: dist})
	}
	for _, mv := range strings.Split(list2, ",") {
		var dir = cast.ToString(mv[0])
		var dist = cast.ToInt(mv[1:])
		wire2 = append(wire2, move{dir: dir, dist: dist})
	}
	return wire1, wire2
}
