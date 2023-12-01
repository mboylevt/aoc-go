package main

import (
	_ "embed"
	"fmt"
	"regexp"
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

	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
	ans = part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func isDigit(i byte) bool {
	if i >= 48 && i <= 57 {
		return true
	}
	return false
}

func getFirstDigit(calibration string) string {
	for i := 0; i < len(calibration); i++ {
		if isDigit(calibration[i]) {
			return string(calibration[i])
		}
	}
	return ""
}

func getLastDigit(calibration string) string {
	for i := len(calibration) - 1; i >= 0; i-- {
		if isDigit(calibration[i]) {
			return string(calibration[i])
		}
	}
	return ""
}

func _doDigitSwap(calibration string, flip string) string {
	one, _ := regexp.Compile("one")
	two, _ := regexp.Compile("two")
	three, _ := regexp.Compile("three")
	four, _ := regexp.Compile("four")
	five, _ := regexp.Compile("five")
	six, _ := regexp.Compile("six")
	seven, _ := regexp.Compile("seven")
	eight, _ := regexp.Compile("eight")
	nine, _ := regexp.Compile("nine")
	switch flip {
	case "one":
		calibration = one.ReplaceAllString(calibration, "one1one")
	case "two":
		calibration = two.ReplaceAllString(calibration, "two2two")
	case "three":
		calibration = three.ReplaceAllString(calibration, "three3three")
	case "four":
		calibration = four.ReplaceAllString(calibration, "four4four")
	case "five":
		calibration = five.ReplaceAllString(calibration, "five5five")
	case "six":
		calibration = six.ReplaceAllString(calibration, "six6six")
	case "seven":
		calibration = seven.ReplaceAllString(calibration, "seven7seven")
	case "eight":
		calibration = eight.ReplaceAllString(calibration, "eight8eight")
	case "nine":
		calibration = nine.ReplaceAllString(calibration, "nine9nine")
	}

	return calibration
}

func changeWordsToDigits(calibration string) string {

	one, _ := regexp.Compile("one")
	two, _ := regexp.Compile("two")
	three, _ := regexp.Compile("three")
	four, _ := regexp.Compile("four")
	five, _ := regexp.Compile("five")
	six, _ := regexp.Compile("six")
	seven, _ := regexp.Compile("seven")
	eight, _ := regexp.Compile("eight")
	nine, _ := regexp.Compile("nine")
	calibration = one.ReplaceAllString(calibration, "one1one")
	calibration = two.ReplaceAllString(calibration, "two2two")
	calibration = three.ReplaceAllString(calibration, "three3three")
	calibration = four.ReplaceAllString(calibration, "four4four")
	calibration = five.ReplaceAllString(calibration, "five5five")
	calibration = six.ReplaceAllString(calibration, "six6six")
	calibration = seven.ReplaceAllString(calibration, "seven7seven")
	calibration = eight.ReplaceAllString(calibration, "eight8eight")
	calibration = nine.ReplaceAllString(calibration, "nine9nine")
	return calibration
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	sum := 0

	for _, calibration := range parsed {
		first := getFirstDigit(calibration)
		last := getLastDigit(calibration)
		digit := first + last
		sum += cast.ToInt(digit)
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed
	sum := 0

	for _, calibration := range parsed {
		calibration = changeWordsToDigits(calibration)
		first := getFirstDigit(calibration)
		last := getLastDigit(calibration)
		digit := first + last
		sum += cast.ToInt(digit)
	}

	return sum
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
