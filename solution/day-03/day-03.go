package day03

import (
	"fmt"
	"regexp"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

/**
 * Return the sum of the products where mul(a, b) is the product of a and b.
 */
func part1(input string) string {
	result := 0

	puzzleInput := strings.ReplaceAll(input, "\n", "")

	result += sumProducts(puzzleInput)
	return fmt.Sprintf("%d", result)
}

/**
 * Return the sum of the products where mul(a, b) is the product of a and b.
 * Only do this when the string when the instruction "do()" was last given, rather than "don't()".
 */
func part2(input string) string {
	result := 0
	puzzleInput := strings.ReplaceAll(input, "\n", "")

	for {
		index := strings.Index(puzzleInput, "don't()")
		if index == -1 {
			index = len(puzzleInput)
		}

		// Get the substring before the "don't()" instruction
		substring := puzzleInput[:index]
		// Update the line to be the substring after the "don't()" instruction
		puzzleInput = puzzleInput[index:]
		result += sumProducts(substring)

		index = strings.Index(puzzleInput, "do()")
		if index != -1 {
			// Skip to the "do()" instruction
			puzzleInput = puzzleInput[index:]
		} else {
			break
		}
	}

	return fmt.Sprintf("%d", result)
}

func sumProducts(input string) int {
	sum := 0
	pattern := `(mul\((\d+),(\d+)\))`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			num1 := util.AtoI(match[2])
			num2 := util.AtoI(match[3])
			sum += num1 * num2
		}
	}

	return sum
}
