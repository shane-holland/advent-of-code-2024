package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"shaneholland.dev/aoc-2024/util"
)

const PART_1_EXPECTED = "1928"
const PART_2_EXPECTED = "2858"
const PUZZLE_INPUT_PATH = "../../data/day-09/test.txt"

func TestPart1(t *testing.T) {
	testInput := util.ReadFile(PUZZLE_INPUT_PATH)
	solver := Puzzle{}
	answer1, _ := solver.Solve(testInput)

	assert.Equal(t, PART_1_EXPECTED, answer1)
}

func TestPart2(t *testing.T) {
	testInput := util.ReadFile(PUZZLE_INPUT_PATH)
	solver := Puzzle{}
	_, answer2 := solver.Solve(testInput)

	assert.Equal(t, PART_2_EXPECTED, answer2)
}
