/* -------------------------------------------------------------------------- */
/*                             Advent of Code 2024                            */
/* -------------------------------------------------------------------------- */
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"shaneholland.dev/aoc-2024/solution"
	"shaneholland.dev/aoc-2024/util"
)

/* ----------------------------- Command Handler ---------------------------- */

// Main function to run the Advent of Code 2024 solutions.
// The day flag is used to specify which day to run the solution for.
func main() {
	args := getArgs()
	
	day := args["day"]
	path := fmt.Sprintf("day-%s", formatDay(day))

	if Solver, ok := solution.Solutions[path]; ok {
		start := time.Now()
		done := make(chan struct{})

		// Read the input file
		input := util.ReadFile("./data/" + path + "/data.txt")
		fmt.Printf("🎄 Advent of Code [2024] - Day %v %v:\n", day, Solver.Icon)

		// Run the solution
		go indicator(done)
		Solve(Solver.Solution, input, done)

		fmt.Printf("🕒 Execution Time: %v\n", time.Since(start))

		
	} else {
		log.Fatalf("Invalid day specified. No solution exists for day %s.\n", day)
	}
}

/* ----------------------------- Helper Methods ----------------------------- */

// Get the command line arguments and return them as a map.
func getArgs() map[string]string {
	args := make(map[string]string)

	// Flags Definitions
	day := flag.String("day", "01", "The day of the Advent of Code challenge to run.")
	
	// Parse Flags
	flag.Parse()

	// Populate the args Map
	args["day"] = *day

	return args
}

// Ensure the day string is formatted correctly.
func formatDay(day string) string {
	if len(day) == 1 {
		return "0" + day
	}
	return day
}

func Solve(Solver solution.Solution, input string, done chan struct{}) {
	answer1, answer2 := Solver.Solve(input)

	// Clear the "Solving" indicator
	fmt.Print("\033[2K")
	// show the cursor
	fmt.Print("\x1B[?25h")
	fmt.Println()
		
	fmt.Printf("\t✅ Part 1 Solution: %s\n", answer1)
	fmt.Printf("\t✅ Part 2 Solution: %s\n\n", answer2)

	close(done)
}

func indicator(done chan struct{}) {
	ticker := time.NewTicker(500 * time.Millisecond)
	fmt.Print("\t⏳ Solving: ")
	// Save the cursor position
	fmt.Print("\x1B7")
	// Hide the cursor
	fmt.Print("\x1B[?25l")

	animation := []string{"|", "/", "-", "\\"}

	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// Restore the cursor position
			fmt.Print("\x1B8")
			// Save the cursor position
			fmt.Print("\x1B7")
			fmt.Print(animation[0])
			animation = append(animation[1:], animation[0])
		case <-done:
			return
		}
	}
}
