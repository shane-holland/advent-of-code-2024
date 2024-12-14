/* -------------------------------------------------------------------------- */
/*                                Solution Map                                */
/* -------------------------------------------------------------------------- */
package solution

import (
	day01 "shaneholland.dev/aoc-2024/solution/day-01"
	day02 "shaneholland.dev/aoc-2024/solution/day-02"
	day03 "shaneholland.dev/aoc-2024/solution/day-03"
	day04 "shaneholland.dev/aoc-2024/solution/day-04"
	day05 "shaneholland.dev/aoc-2024/solution/day-05"
	day06 "shaneholland.dev/aoc-2024/solution/day-06"
	day07 "shaneholland.dev/aoc-2024/solution/day-07"
	day08 "shaneholland.dev/aoc-2024/solution/day-08"
	day09 "shaneholland.dev/aoc-2024/solution/day-09"
	day10 "shaneholland.dev/aoc-2024/solution/day-10"
	day11 "shaneholland.dev/aoc-2024/solution/day-11"
	day12 "shaneholland.dev/aoc-2024/solution/day-12"
	day13 "shaneholland.dev/aoc-2024/solution/day-13"
	day14 "shaneholland.dev/aoc-2024/solution/day-14"
)

// Solver is a struct that contains the Solution and an icon for the Advent of Code problem.
type Solver struct {
	Solution Solution
	Icon     string
}

// Solutions is a map of Solvers to the Advent of Code problems.
var Solutions = map[string]Solver{
	"day-01": {day01.Puzzle{}, "🕵"},
	"day-02": {day02.Puzzle{}, "🦌"},
	"day-03": {day03.Puzzle{}, "🧮"},
	"day-04": {day04.Puzzle{}, "🔎"},
	"day-05": {day05.Puzzle{}, "🖨️"},
	"day-06": {day06.Puzzle{}, "💂"},
	"day-07": {day07.Puzzle{}, "🌉"},
	"day-08": {day08.Puzzle{}, "📡"},
	"day-09": {day09.Puzzle{}, "💾"},
	"day-10": {day10.Puzzle{}, "🥾"},
	"day-11": {day11.Puzzle{}, "🪨"},
	"day-12": {day12.Puzzle{}, "🪴"},
	"day-13": {day13.Puzzle{}, "🕹️"},
	"day-14": {day14.Puzzle{}, "🚽"},
}
