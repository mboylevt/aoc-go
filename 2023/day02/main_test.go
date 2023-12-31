package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	want := 8

	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	want := 2286

	if got := part2(input); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}

func Test_parser(t *testing.T) {
	input := "Game 2341: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	games := []game{
		game{red: 4, blue: 3, green: 0},
		game{red: 1, blue: 6, green: 2},
		game{red: 0, blue: 0, green: 2},
	}
	want := map[int][]game{2341: games}
	got := gameParser(input)
	for i := 0; i < len(got[2341]); i++ {
		if got[2341][i] != want[2341][i] {
			t.Errorf("gameParser() = %v, want %v", got, want)
		}
	}
}
