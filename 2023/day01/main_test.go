package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := "1abc2\npr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	want := 142

	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	want := 281

	if got := part2(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_regex(t *testing.T) {
	input := "eightwothree"
	want := "eight8eightwo2twothree3three"

	if got := changeWordsToDigits(input); got != want {
		t.Errorf("Regex = %v, want %v", got, want)
	}
}
