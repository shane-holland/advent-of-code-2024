/* -------------------------------------------------------------------------- */
/*                       --- Day 15: Warehouse Woes ---                       */
/* -------------------------------------------------------------------------- */
package day15

import (
	"fmt"
	"slices"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

/* ------------------------------- Main Method ------------------------------ */
type Puzzle struct{}

// The Solve method is called to solve the puzzle.
func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

/* -------------------------------- Solution -------------------------------- */

// Part 1:
func part1(input string) string {
	warehouse := parseWarehouse(input)
	for warehouse.NextInstruction() {
		// Move boxes until out of instructions
	}
	return fmt.Sprintf("%d", warehouse.BoxGpsSum())
}

// Part 2:
func part2(input string) string {
	warehouse := parseWideWarehouse(input)
	for warehouse.NextInstruction() {
		// Move boxes until out of instructions
		// warehouse.Draw()
	}
	// warehouse.Draw()
	return fmt.Sprintf("%d", warehouse.BoxGpsSum())
}

/* -------------------------------- Constants ------------------------------- */

// Part 1
const FREE = 0
const BOX = 1
const WALL = 2

// Part 2
const BOX_LEFT = 3
const BOX_RIGHT = 4

var Directions = map[string]util.Point{
	"^": {X: 0, Y: -1},
	"v": {X: 0, Y: 1},
	"<": {X: -1, Y: 0},
	">": {X: 1, Y: 0},
}

/* -------------------- Warehouse Definition and Methods -------------------- */

type Warehouse struct {
	Map          [][]int
	Robot        util.Point
	Instructions []util.Point
}

func (w Warehouse) Gps(point util.Point) int {
	return point.X + (100 * point.Y)
}

func (w Warehouse) Draw() {
	for y := 0; y < len(w.Map); y++ {
		for x := 0; x < len(w.Map[y]); x++ {
			switch w.Map[y][x] {
			case WALL:
				print("#")
			case FREE:
				print(".")
			case BOX:
				print("O")
			}
		}
		println()
	}
}

func (w *Warehouse) NextInstruction() bool {
	if len(w.Instructions) == 0 {
		return false
	}

	next := w.Instructions[0]
	w.Instructions = w.Instructions[1:]

	if w.Map[w.Robot.Y+next.Y][w.Robot.X+next.X] == WALL {
		// Do nothing
	} else if w.Map[w.Robot.Y+next.Y][w.Robot.X+next.X] == FREE {
		// Free space, the robot can move
		w.Robot.X += next.X
		w.Robot.Y += next.Y
	} else if w.Map[w.Robot.Y+next.Y][w.Robot.X+next.X] == BOX {
		// Box in the way, see if the box can move
		boxStack := []util.Point{{X: w.Robot.X + next.X, Y: w.Robot.Y + next.Y}}
		for {
			test := util.Point{X: boxStack[len(boxStack)-1].X + next.X, Y: boxStack[len(boxStack)-1].Y + next.Y}
			if w.Map[test.Y][test.X] == WALL {
				// We've hit a wall.  No movement is possible
				boxStack = []util.Point{}
				break
			} else if w.Map[test.Y][test.X] == FREE {
				// The robot and box stack can move
				w.Robot.X += next.X
				w.Robot.Y += next.Y
				w.Map[w.Robot.Y][w.Robot.X] = FREE
				break
			} else if w.Map[test.Y][test.X] == BOX {
				// Add the box to the stack
				boxStack = append(boxStack, test)
			}
		}
		for len(boxStack) > 0 {
			w.Map[boxStack[0].Y+next.Y][boxStack[0].X+next.X] = BOX
			boxStack = boxStack[1:]
		}
	}
	return true
}

func (w Warehouse) BoxGpsSum() int {
	sum := 0
	for y := 0; y < len(w.Map); y++ {
		for x := 0; x < len(w.Map[y]); x++ {
			if w.Map[y][x] == BOX {
				sum += w.Gps(util.Point{X: x, Y: y})
			}
		}
	}
	return sum
}

/* ------------------ WideWarehouse Definition and Methods ------------------ */

type WideWarehouse struct {
	Map          [][]int
	Robot        util.Point
	Instructions []util.Point
}

func (w WideWarehouse) Draw() {
	for y := 0; y < len(w.Map); y++ {
		for x := 0; x < len(w.Map[y]); x++ {
			switch w.Map[y][x] {
			case WALL:
				print("#")
			case FREE:
				if w.Robot.X == x && w.Robot.Y == y {
					print("@")
				} else {
					print(".")
				}
			case BOX_LEFT:
				print("[")
			case BOX_RIGHT:
				print("]")
			}
		}
		println()
	}
}

func (w WideWarehouse) Gps(point util.Point) int {
	return point.X + (100 * point.Y)
}

func (w *WideWarehouse) NextInstruction() bool {
	if len(w.Instructions) == 0 {
		return false
	}

	next := w.Instructions[0]
	w.Instructions = w.Instructions[1:]

	if w.Map[w.Robot.Y+next.Y][w.Robot.X+next.X] == WALL {
		// Do nothing
	} else if w.Map[w.Robot.Y+next.Y][w.Robot.X+next.X] == FREE {
		// Free space, the robot can move
		w.Robot.X += next.X
		w.Robot.Y += next.Y
	} else if w.Map[w.Robot.Y+next.Y][w.Robot.X+next.X] > BOX {
		// Box in the way, see if the box can move
		boxes := []util.Point{{X: w.Robot.X + next.X, Y: w.Robot.Y + next.Y}}
		if w.Map[w.Robot.Y+next.Y][w.Robot.X+next.X] == BOX_LEFT {
			boxes = append(boxes, util.Point{X: w.Robot.X + next.X + 1, Y: w.Robot.Y + next.Y})
		} else {
			boxes = append(boxes, util.Point{X: w.Robot.X + next.X - 1, Y: w.Robot.Y + next.Y})
		}
		boxStacks := [][]util.Point{boxes}
		for {
			testStack := boxStacks[len(boxStacks)-1]
			boxesAdded := false
			wall := false
			nextStack := []util.Point{}
			for _, box := range testStack {
				test := util.Point{X: box.X + next.X, Y: box.Y + next.Y}
				if w.Map[test.Y][test.X] == WALL {
					// We've hit a wall.  No movement is possible
					wall = true
					break
				} else if w.Map[test.Y][test.X] > BOX && !slices.Contains(testStack, test) {
					if !slices.Contains(nextStack, util.Point{X: test.X, Y: test.Y}) {
						if w.Map[test.Y][test.X] == BOX_LEFT {
							nextStack = append(nextStack, []util.Point{test, {X: test.X + 1, Y: test.Y}}...)
						} else {
							nextStack = append(nextStack, []util.Point{test, {X: test.X - 1, Y: test.Y}}...)
						}
						boxesAdded = true
					}
				}
			}

			if wall {
				// We've hit a wall.  No movement is possible
				boxStacks = [][]util.Point{}
				break
			}
			if !boxesAdded {
				// The robot and box stack can move
				// w.Robot.X += next.X
				// w.Robot.Y += next.Y
				// w.Map[w.Robot.Y][w.Robot.X] = FREE
				break
			} else {
				boxStacks = append(boxStacks, nextStack)
			}
		}
		if (len(boxStacks) > 0) {
			temp := make([][]int, len(w.Map))
			for i, row := range w.Map {
				temp[i] =slices.Clone(row)
			}
			for len(boxStacks)> 0 {
				boxstack := boxStacks[len(boxStacks)-1]
				boxStacks = boxStacks[:len(boxStacks)-1]
				for len(boxstack) > 0 {
					box := boxstack[len(boxstack)-1]
					boxstack = boxstack[:len(boxstack)-1]
					w.Map[box.Y][box.X] = FREE
					w.Map[box.Y+next.Y][box.X+next.X] = temp[box.Y][box.X]
				}
			}
			w.Robot.X += next.X
			w.Robot.Y += next.Y
			// if next.Y != 0 {
			// 	if w.Map[w.Robot.Y][w.Robot.X+1] == BOX_RIGHT {
			// 		w.Map[w.Robot.Y][w.Robot.X+1] = FREE
			// 	}
			// 	if w.Map[w.Robot.Y][w.Robot.X-1] == BOX_LEFT {
			// 		w.Map[w.Robot.Y][w.Robot.X-1] = FREE
			// 	}
			// }
			// w.Map[w.Robot.Y][w.Robot.X] = FREE
		}
	}
	return true
}

func (w WideWarehouse) BoxGpsSum() int {
	sum := 0
	for y := 0; y < len(w.Map); y++ {
		for x := 0; x < len(w.Map[y]); x++ {
			if w.Map[y][x] == BOX_LEFT {
				sum += w.Gps(util.Point{X: x, Y: y})
			}
		}
	}
	return sum
}

/* ----------------------------- Helper Methods ----------------------------- */

func parseWarehouse(input string) Warehouse {
	robot := util.Point{}

	mapGrid := strings.Split(strings.Split(input, "\n\n")[0], "\n")
	instructionsString := strings.Split(input, "\n\n")[1]
	var warehouseMap [][]int = make([][]int, len(mapGrid))

	for y := 0; y < len(mapGrid); y++ {
		warehouseMap[y] = make([]int, len(mapGrid[y]))
		for x := 0; x < len(mapGrid[y]); x++ {
			switch string(mapGrid[y][x]) {
			case "#":
				warehouseMap[y][x] = WALL
			case ".":
				warehouseMap[y][x] = FREE
			case "O":
				warehouseMap[y][x] = BOX
			case "@":
				robot = util.Point{X: x, Y: y}
			}
		}
	}

	return Warehouse{
		Map:          warehouseMap,
		Robot:        robot,
		Instructions: parseInstructions(instructionsString)}
}

func parseWideWarehouse(input string) WideWarehouse {
	robot := util.Point{}

	mapGrid := strings.Split(strings.Split(input, "\n\n")[0], "\n")
	instructionsString := strings.Split(input, "\n\n")[1]
	var warehouseMap [][]int = make([][]int, len(mapGrid)*2)

	for y := 0; y < len(mapGrid); y++ {
		warehouseMap[y] = make([]int, len(mapGrid[y])*2)
		for x := 0; x < len(mapGrid[y]); x++ {
			switch string(mapGrid[y][x]) {
			case "#":
				warehouseMap[y][x*2] = WALL
				warehouseMap[y][x*2+1] = WALL
			case ".":
				warehouseMap[y][x*2] = FREE
				warehouseMap[y][x*2+1] = FREE
			case "O":
				warehouseMap[y][x*2] = BOX_LEFT
				warehouseMap[y][x*2+1] = BOX_RIGHT
			case "@":
				robot = util.Point{X: x * 2, Y: y}
				warehouseMap[y][x*2+1] = FREE
			}
		}
	}

	return WideWarehouse{
		Map:          warehouseMap,
		Robot:        robot,
		Instructions: parseInstructions(instructionsString)}
}

func parseInstructions(input string) []util.Point {
	instructions := []util.Point{}

	input = strings.ReplaceAll(input, "\n", "")
	for _, direction := range strings.Split(input, "") {
		instructions = append(instructions, Directions[direction])
	}
	return instructions
}
