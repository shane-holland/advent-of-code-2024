package main

import (
	"flag"
	"fmt"
	"log"

	"shaneholland.dev/aoc-2024/util"
	"shaneholland.dev/aoc-2024/solution"
)

func main() {
	dayPtr := flag.String("day", "01", "The day of the Advent of Code challenge to run.")
	flag.Parse()

	dayString := *dayPtr
	if len(dayString) == 1 {
		dayString = "0" + dayString
	}

	if _, ok := solution.Solutions["day-"+dayString]; !ok {
		log.Fatalf("Invalid day specified. No solution exists for day %s.\n", *dayPtr)
	}
	// Read the input file
	input := util.ReadFile("./data/day-" + dayString + "/data.txt")
	Solver := solution.Solutions["day-"+dayString]

	answer1, answer2 := Solver.Solve(input)
	fmt.Printf("Day %v:\n", *dayPtr)
	fmt.Printf("\tPart 1: %s\n", answer1)
	fmt.Printf("\tPart 2: %s\n\n", answer2)

}
