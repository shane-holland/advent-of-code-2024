/* -------------------------------------------------------------------------- */
/*                         --- Day 4: Ceres Search ---                        */
/* -------------------------------------------------------------------------- */
package day04

import (
	"fmt"
	"math"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

// Part 1: Count the number of instances of the string "XMAS" (forward or reversed) in the word puzzle
func part1(input string) string {
	instances := 0
	// Horizontal
	instances += countInstancesXmas(util.GetLines(input))
	// Vertical
	instances += countInstancesXmas(util.GetColumns(input))
	// Diagonal SW-NE
	instances += countInstancesXmas(getDiagonalsSwNe(input))
	// Diagonal NW-SE
	instances += countInstancesXmas(getDiagonalsNwSe(input))

	return fmt.Sprintf("%d", instances)
}

// Part 2: Count the number of instances where MAS appears (forward or reversed) in an X pattern in the word puzzle
func part2(input string) string {
	matrix := util.GetLines(input)
	instances := 0

	for y := 0; y < len(matrix)-2; y++ {
		for x := 0; x < len(matrix[y])-2; x++ {
			word1 := string(matrix[y][x]) + string(matrix[y+1][x+1]) + string(matrix[y+2][x+2])
			word2 := string(matrix[y][x+2]) + string(matrix[y+1][x+1]) + string(matrix[y+2][x])

			if (word1 == "SAM" || word1 == "MAS") && (word2 == "SAM" || word2 == "MAS") {
				instances++
			}
		}
	}

	return fmt.Sprintf("%d", instances)
}

/* ----------------------------- Helper Methods ----------------------------- */

// Returns the number of instances of the string "XMAS" (forward or reversed) in the input string array
func countInstancesXmas(lines []string) int {
	instances := 0

	for _, line := range lines {
		substr := line
		for {

			xmasIndex := strings.Index(substr, "XMAS")
			samxIndex := strings.Index(substr, "SAMX")

			if xmasIndex == -1 {
				if samxIndex == -1 {
					break
				}
				substr = substr[samxIndex+1:]
			} else if samxIndex == -1 {
				substr = substr[xmasIndex+1:]
			} else {
				substr = substr[int(math.Min(float64(xmasIndex), float64(samxIndex)))+1:]
			}
			instances++
		}
	}
	return instances
}

// Returns the diagonals of the input string array from the SW to NE
func getDiagonalsSwNe(input string) []string {
	var diagonalLines []string
	lines := util.GetLines(input)

	for i := 0; i < len(lines); i++ {
		var diagonalLine1 string
		var diagonalLine2 string
		for j := 0; j < len(lines); j++ {
			if i+j < len(lines) {
				diagonalLine1 += string(lines[i+j][len(lines[i+j])-(j+1)])
			}
			if i > 0 && j+i < len(lines) {
				diagonalLine2 += string(lines[j][len(lines[i+j])-(j+i+1)])
			}
		}
		diagonalLines = append(diagonalLines, []string{diagonalLine1, diagonalLine2}...)
	}

	return diagonalLines
}

// Returns the diagonals of the input string array from the NW to SE
func getDiagonalsNwSe(input string) []string {
	var diagonalLines []string
	lines := util.GetLines(input)

	for i := 0; i < len(lines); i++ {
		var diagonalLine1 string
		var diagonalLine2 string
		for j := 0; j < len(lines); j++ {
			if i+j < len(lines) {
				diagonalLine1 += string(lines[i+j][j])
			}
			if i > 0 && j+i < len(lines) {
				diagonalLine2 += string(lines[j][j+i])
			}
		}
		diagonalLines = append(diagonalLines, []string{diagonalLine1, diagonalLine2}...)
	}

	return diagonalLines
}
