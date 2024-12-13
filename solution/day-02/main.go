/* -------------------------------------------------------------------------- */
/*                      --- Day 2: Red-Nosed Reports ---                      */
/* -------------------------------------------------------------------------- */
package day02

import (
	"strconv"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

// Part 1: Find the number of safe reports.
func part1(input string) string {
	safeReports := 0

	for _, line := range util.GetLines(input) {
		report := parseLine(line)

		if isSafe(report) {
			safeReports++
		}
	}
	return strconv.Itoa(safeReports)
}

// Part 2: Find the number of safe reports with a dampener.
func part2(input string) string {
	safeReports := 0

	for _, line := range util.GetLines(input) {
		report := parseLine(line)

		if isSafeWithDampener(report) {
			safeReports++
		}
	}
	return strconv.Itoa(safeReports)
}

/* ----------------------------- Helper Methods ----------------------------- */

// Parse the line into a slice of integers.
func parseLine(line string) []int {
	var levels []int
	for _, level := range strings.Split(line, " ") {
		levelNum := util.AtoI(level)
		levels = append(levels, levelNum)
	}
	return levels
}

// Check if the levels are safe.
// A level is safe if the difference between it and the next level is between 1 and 3.
func isSafe(levels []int) bool {
	ascending := levels[0] < levels[1]
	for i, level := range levels[1:] {

		if !areSafe(levels[i], level, ascending) {
			return false
		}

	}
	return true
}

// Check if the levels are safe with a dampener.
// The dampener can remove exactly one level to make the levels safe.
func isSafeWithDampener(levels []int) bool {
	if !isSafe(levels) {
		for i := 0; i < len(levels); i++ {
			levelCopy := make([]int, len(levels))
			copy(levelCopy, levels)
			revised := append(levelCopy[:i], levelCopy[i+1:]...)
			if isSafe(revised) {
				return true
			}
		}
		return false
	}
	return true
}

// Check if the levels are safe.
func areSafe(a, b int, ascending bool) bool {
	difference := util.AbsInt(a - b)
	return !((ascending && a > b) || (!ascending && a < b) || difference < 1 || difference > 3)
}
