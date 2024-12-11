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

/**
 * Function to find the minimum of two integers.
 */
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

/**
 * Function to find the minimum of two integers.
 */
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

func parseEquation(input string, concat bool) Equation {
	testValue := util.AtoI(strings.Split(input, ": ")[0])
	components := make([]int, 0)

	for _, component := range strings.Split(strings.Split(input, ": ")[1], " ") {
		components = append(components, util.AtoI(component))
	}

	return Equation{TestValue: testValue, Components: components, ConcatEnabled: concat}
}

type Equation struct {
	TestValue     int
	Components    []int
	ConcatEnabled bool
}

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
