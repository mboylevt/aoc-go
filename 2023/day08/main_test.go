package main

import (
	"testing"
)

func Test_part1_first(t *testing.T) {
	input := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	want := 2

	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part1_second(t *testing.T) {
	input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	want := 6

	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	want := 6

	if got := part2(input); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
