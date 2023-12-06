package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	want := 288

	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
