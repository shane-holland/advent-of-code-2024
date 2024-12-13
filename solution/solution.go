// Package solution defines the logic for solving Advent of Code problems.
package solution

// Each solution in this package should implement the Solution interface.
type Solution interface {
	Solve(string) (string, string)
}