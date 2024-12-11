package day01

import (
	"log"
	"regexp"
	"sort"
	"strconv"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	left, right := parseInput(input)
	return part1(left, right), part2(left, right)
}

/**
 * Function to find the minimum of two integers.
 */
func part1(left, right []int) string {
	max := min(len(left), len(right))

	distance := 0
	for i := 0; i < max; i++ {
		distance += util.AbsInt(left[i] - right[i])
	}

	return strconv.Itoa(distance)
}

/**
 * Function to find the minimum of two integers.
 */
func part2(left, right []int) string {
	similarity := 0

	index := 0
	for _, num := range left {
		matches := 0
		for index < len(right) && right[index] <= num {
			if right[index] == num {
				matches++
			}
			index++
		}
		// fmt.Printf("Num: %d, Matches: %d\n", num, matches)
		similarity += matches * num
		index -= matches
	}

	return strconv.Itoa(similarity)
}

// Function to parse the input into two arrays of integers
func parseInput(input string) (left []int, right []int) {
	lines := util.GetLines(input)

	for _, line := range lines {
		a, b := parseLine(line)
		left = append(left, a)
		right = append(right, b)
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

// Function to parse a line of input into two integers
func parseLine(line string) (a, b int) {
	pattern := `(\d+)\s+(\d+)`

	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)

	// Check if matches were found
	if len(matches) > 0 {
		// matches[0] is the entire match
		// matches[1:] are the captured groups
		a = util.AtoI(matches[1])
		b = util.AtoI(matches[2])
	} else {
		log.Fatalf("No match found for line: %s", line)
	}

	return a, b
}
