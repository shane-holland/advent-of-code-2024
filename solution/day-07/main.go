/* -------------------------------------------------------------------------- */
/*                        --- Day 7: Bridge Repair ---                        */
/* -------------------------------------------------------------------------- */
package day07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

// Part 1: Find the sum of all test values that pass the equation.
func part1(input string) string {
	sum := 0

	for _, line := range util.GetLines(input) {
		equation := parseEquation(line, false)
		if equation.Test() {
			sum += equation.TestValue
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Part 2: Find the sum of all test values that pass the equation with concatenation.
func part2(input string) string {
	sum := 0

	for _, line := range util.GetLines(input) {
		equation := parseEquation(line, true)
		if equation.Test() {
			sum += equation.TestValue
		}
	}
	return fmt.Sprintf("%d", sum)
}

/* --------------------- Equation Definition and Methods -------------------- */

// Equation represents possible mathematical equations with a test value and components.
type Equation struct {
	TestValue     int
	Components    []int
	ConcatEnabled bool
}

// Test the equation with all possible combinations of operators against components.
// Returns true if the test value is found in the possible values.
func (eq *Equation) Test() bool {
	possibleValues := make([]int, 0)
	possibleValues = append(possibleValues, eq.Components[0])
	for _, component := range eq.Components[1:] {
		newValues := make([]int, 0)
		for _, value := range possibleValues {
			newValues = append(newValues, value+component)
			newValues = append(newValues, value*component)
			if eq.ConcatEnabled {
				concat := util.AtoI(strconv.Itoa(value) + strconv.Itoa(component))
				newValues = append(newValues, concat)
			}
		}
		possibleValues = newValues
	}
	return slices.Contains(possibleValues, eq.TestValue)
}

/* ----------------------------- Helper Methods ----------------------------- */

// Parse an Equation from a string.
// The equation is formatted as "{TestValue}: {Component1} {Component2} ...".
// If concat is true, then the equation will also test concatenation of components.
func parseEquation(input string, concat bool) Equation {
	testValue := util.AtoI(strings.Split(input, ": ")[0])
	components := make([]int, 0)

	for _, component := range strings.Split(strings.Split(input, ": ")[1], " ") {
		components = append(components, util.AtoI(component))
	}

	return Equation{TestValue: testValue, Components: components, ConcatEnabled: concat}
}