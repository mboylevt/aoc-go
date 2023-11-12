package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/mboylevt/aoc-go/util"
)

//go:embed input.txt
var input string

type orb struct {
	name   string
	orbits string
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

func part1(input string) int {
	parsed := parseInput(input)
	bodies := make(map[string][]string)
	for _, entry := range parsed {
		o := strings.Split(entry, ")")
		bodies[o[0]] = append(bodies[o[0]], o[1])
	}

	totalOrbits := countOrbits(bodies, "COM", 0)
	return totalOrbits
}

func part2(input string) int {
	parsed := parseInput(input)
	bodies := make(map[string][]string)
	for _, entry := range parsed {
		o := strings.Split(entry, ")")
		bodies[o[0]] = append(bodies[o[0]], o[1])
	}
	var santaPath []string
	var youPath []string
	santaPath = hasPath(bodies, santaPath, "COM", "SAN")
	youPath = hasPath(bodies, youPath, "COM", "YOU")
	commonAncestor := -1
	for idx := 0; idx < len(santaPath); idx++ {
		if santaPath[idx] != youPath[idx] {
			commonAncestor = idx - 1
			break
		}
	}
	santaToAncestor := len(santaPath) - (commonAncestor + 2)
	youToAncestor := len(youPath) - (commonAncestor + 2)

	fmt.Printf("Santa -> Ancestor: %v\n", santaToAncestor)
	fmt.Printf("You -> Ancestor: %v\n", youToAncestor)
	fmt.Printf("Orbital transfers you->santa: %v\n", santaToAncestor+youToAncestor)
	return 0
}

func countOrbits(bodies map[string][]string, currentBody string, depth int) int {
	myChildrenDepth := 0
	for _, child := range bodies[currentBody] {
		myChildrenDepth += countOrbits(bodies, child, depth+1)
	}
	return depth + myChildrenDepth
}

func hasPath(bodies map[string][]string, path []string, currentBody string, target string) []string {
	path = append(path, currentBody)
	if currentBody == target {
		return path
	}

	for _, child := range bodies[currentBody] {
		newPath := hasPath(bodies, path, child, target)
		if len(newPath) > len(path) {
			return newPath
		}
	}

	// remove this node from the path
	path = path[:len(path)-1]
	return path
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
