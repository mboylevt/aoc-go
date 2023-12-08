package main

import (
	"container/ring"
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/mboylevt/aoc-go/lib"
	"github.com/mboylevt/aoc-go/util"
)

//go:embed input.txt
var input string

const (
	left int = iota
	right
)

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

type node struct {
	id    string
	left  string
	right string
}

func populateRingBuffer(input string) *ring.Ring {
	r := ring.New(len(input))
	for _, c := range input {
		if c == 'L' {
			r.Value = left
		} else {
			r.Value = right
		}
		r = r.Next()
	}
	return r
}

func populateNodes(nodeStrings []string) map[string]node {
	nodeMap := make(map[string]node)

	for _, str := range nodeStrings {
		nodeFields := strings.Fields(str)
		nodeId := nodeFields[0]
		left := nodeFields[2][1:4]
		right := nodeFields[3][0:3]
		nodeMap[nodeId] = node{nodeId, left, right}
	}

	return nodeMap
}

func getNextNode(r *ring.Ring, n node, nodeMap map[string]node) node {
	switch r.Value {
	case left:
		return nodeMap[n.left]
	case right:
		return nodeMap[n.right]
	}
	return n //this should never happen
}

func getStartingNodes(nodeMap map[string]node) []node {
	var startingNodes []node
	for key, n := range nodeMap {
		if key[2] == 'A' {
			startingNodes = append(startingNodes, n)
		}
	}
	return startingNodes
}

func part1(input string) int {
	parsed := parseInput(input)
	r := populateRingBuffer(parsed[0])
	nodeMap := populateNodes(parsed[2:])
	thisNode := nodeMap["AAA"]
	stepCount := 0
	for {
		if thisNode.id == "ZZZ" {
			break
		}
		thisNode = getNextNode(r, thisNode, nodeMap)
		stepCount++
		r = r.Next()
	}
	return stepCount
}

func part2(input string) int {
	parsed := parseInput(input)
	nodeMap := populateNodes(parsed[2:])
	nodesToCheck := getStartingNodes(nodeMap)
	fmt.Printf("Checking %v nodes\n", len(nodesToCheck))
	var distances []int
	// Check each node to find their steps to Z
	for _, n := range nodesToCheck {
		nodeStepCount := 0
		r := populateRingBuffer(parsed[0])
		for {
			if n.id[2] == 'Z' {
				break
			}
			n = getNextNode(r, n, nodeMap)
			r = r.Next()
			nodeStepCount++
		}
		distances = append(distances, nodeStepCount)
	}

	// run LCM checks on all nodes
	return lib.LCM(distances[0], distances[1], distances[2:]...)
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
