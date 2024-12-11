// Package solution defines the logic for solving Advent of Code problems.
//
// Each solution in this package should implement the Solution interface.
package solution


type Solution interface {
	Solve(string) (string, string)
}