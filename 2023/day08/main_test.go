package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	want := 6440

	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	want := 5905

	if got := part2(input); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
