package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"shaneholland.dev/aoc-2024/util"
)

const PART_1_EXPECTED = "4,6,3,5,6,3,5,2,1,0"
const PART_2_EXPECTED = "117440"

const PUZZLE_INPUT_PATH_PART_1 = "./test-data.txt"
const PUZZLE_INPUT_PATH_PART_2 = "./test-data-part-2.txt"

func TestPart1(t *testing.T) {
	testInput := util.ReadFile(PUZZLE_INPUT_PATH_PART_1)
	solver := Puzzle{}
	answer1, _ := solver.Solve(testInput)

	assert.Equal(t, PART_1_EXPECTED, answer1)
}

func TestPart2(t *testing.T) {
	testInput := util.ReadFile(PUZZLE_INPUT_PATH_PART_2)
	solver := Puzzle{}
	_, answer2 := solver.Solve(testInput)

	assert.Equal(t, PART_2_EXPECTED, answer2)
}
