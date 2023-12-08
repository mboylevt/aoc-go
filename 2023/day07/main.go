package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

type (
	hand struct {
		cards string
		bid   int
		ht    handType
	}

	handType int

	typeCompute func(string) handType
)

var values = map[byte]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

const (
	highCard handType = iota
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

func typeFromCards(in string) (out handType) {
	sets := map[int][]byte{}

	for len(in) > 0 {
		card := in[0]
		count := 0
		for i := range in {
			if in[i] == card {
				count++
			}
		}
		sets[count] = append(sets[count], card)
		in = strings.ReplaceAll(in, in[0:1], "")
	}

	switch {
	case sets[5] != nil:
		out = fiveKind
	case sets[4] != nil:
		out = fourKind
	case sets[3] != nil:
		if sets[2] != nil {
			out = fullHouse
		} else {
			out = threeKind
		}
	case sets[2] != nil:
		if len(sets[2]) == 2 {
			out = twoPair
		} else {
			out = onePair
		}
	default:
		out = highCard
	}
	return out
}

func (h *hand) less(other hand, valResolutions map[byte]int) bool {
	if h.ht == other.ht {
		for i := range h.cards {
			if h.cards[i] != other.cards[i] {
				return valResolutions[h.cards[i]] < valResolutions[other.cards[i]]
			}
		}
		return true
	}
	return h.ht < other.ht
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	var hands []hand
	for _, line := range parsed {
		fields := strings.Fields(line)
		cards := fields[0]
		bid := cast.ToInt(fields[1])
		hands = append(hands, hand{cards, bid, typeFromCards(cards)})
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].less(hands[j], values)
	})
	fmt.Println("hi")
	winnings := 0
	for i, h := range hands {
		winnings += (i + 1) * h.bid
	}
	return winnings
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
