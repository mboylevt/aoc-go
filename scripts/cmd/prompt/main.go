package main

import "github.com/mboylevt/aoc-go/scripts/aoc"

func main() {
	day, year, cookie := aoc.ParseFlags()
	aoc.GetPrompt(day, year, cookie)
}
