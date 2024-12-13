// Package solution defines the logic for solving Advent of Code problems.
package solution

// Solution is an interface that defines the contract for solving Advent of Code problems.
type Solution interface {
	// Solve returns the answers to an Advent of Code problem (part1, part2), given the puzzle input as a string.
	Solve(string) (string, string)
}
