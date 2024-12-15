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

// Part 1: Return the sum of the GPS coordinates of all boxes in the warehouse.
func part1(input string) string {
	warehouse := parseWarehouse(input)
	for warehouse.NextInstruction() {
		// Move boxes until out of instructions
	}
	return fmt.Sprintf("%d", warehouse.BoxGpsSum())
}

// Part 2:Return the sum of the GPS coordinates of all boxes in the wide warehouse.
func part2(input string) string {
	warehouse := parseWarehouse(input)
	warehouse.Expand()
	for warehouse.NextInstruction() {
		// Move boxes until out of instructions
	}

	return fmt.Sprintf("%d", warehouse.BoxGpsSum())
}

/* -------------------------------- Constants ------------------------------- */

// Part 1
const FREE = 0
const WALL = 1
const BOX = 2

// Part 2
const BOX_LEFT = 3
const BOX_RIGHT = 4

// Directions Map
var Directions = map[string]util.Point{
	"^": {X: 0, Y: -1},
	"v": {X: 0, Y: 1},
	"<": {X: -1, Y: 0},
	">": {X: 1, Y: 0},
}

/* -------------------- Warehouse Definition and Methods -------------------- */

// Represents a warehouse with wall and box positions, robot position, instructions, and expanded status.
type Warehouse struct {
	Map          [][]int
	Robot        util.Point
	Instructions []util.Point
	Expanded     bool
}

// Returns the GPS value of a Box based on it's location in the warehouse.
func (w Warehouse) Gps(point util.Point) int {
	return point.X + (100 * point.Y)
}

// Returns the sum of the GPS coordinates of all boxes in the warehouse.
func (w Warehouse) BoxGpsSum() int {
	comparator := BOX
	if w.Expanded {
		comparator = BOX_LEFT
	}

	sum := 0
	for y := 0; y < len(w.Map); y++ {
		for x := 0; x < len(w.Map[y]); x++ {
			if w.Map[y][x] == comparator {
				sum += w.Gps(util.Point{X: x, Y: y})
			}
		}
	}
	return sum
}

// Expands the warehouse by doubling the width of each row.
func (w *Warehouse) Expand() {
	expandedMap := make([][]int, len(w.Map))
	for y := 0; y < len(w.Map); y++ {
		expandedMap[y] = make([]int, len(w.Map[y])*2)
		for x := 0; x < len(w.Map[y]); x++ {
			if w.Map[y][x] == BOX {
				expandedMap[y][x*2] = BOX_LEFT
				expandedMap[y][x*2+1] = BOX_RIGHT
			} else {
				expandedMap[y][x*2] = w.Map[y][x]
				expandedMap[y][x*2+1] = w.Map[y][x]
			}
		}
	}
	w.Map = expandedMap

	// Adjust the robot's position
	w.Robot.X *= 2
	w.Map[w.Robot.Y][w.Robot.X*2+1] = FREE

	w.Expanded = true
}

// Runs the next robot instruction in the warehouse.
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
	} else {
		// Box in the way, see if the box can move
		boxStacks := w.getBoxesToPush(next)

		// If there are boxes to push, push them
		if len(boxStacks) > 0 {
			w.pushBoxes(boxStacks, next)
			// Update the robot's position
			w.Robot.X += next.X
			w.Robot.Y += next.Y
		}
	}
	return true
}

// Returns a stack of boxes that can be pushed in the direction of the next instruction.
func (w Warehouse) getBoxesToPush(next util.Point) BoxStackStack {
	boxStacks := BoxStackStack{w.getBox(util.Point{X: w.Robot.X + next.X, Y: w.Robot.Y + next.Y})}

	for {
		boxesAdded := false
		testStack := boxStacks.Peek()
		nextStack := make(BoxStack, 0)

		for len(testStack) > 0 {
			box := testStack.Pop()
			test := util.Point{X: box.X + next.X, Y: box.Y + next.Y}
			value := w.Map[test.Y][test.X]

			if value == WALL {
				// We've hit a wall.  No movement is possible
				return make(BoxStackStack, 0)
			} else if value >= BOX && !slices.Contains(testStack, test) {
				if !slices.Contains(nextStack, test) {
					nextStack = append(nextStack, w.getBox(test)...)
					boxesAdded = true
				}
			}
		}

		if boxesAdded {
			boxStacks.Push(nextStack)
		} else {
			break
		}
	}

	return boxStacks
}

// Returns a stack with all components of a box at a given position
func (w Warehouse) getBox(position util.Point) BoxStack {
	box := BoxStack{position}
	// If this box is a double box, add the other side to the stack
	if w.Map[position.Y][position.X] == BOX_LEFT {
		box.Push(util.Point{X: position.X + 1, Y: position.Y})
	} else if w.Map[position.Y][position.X] == BOX_RIGHT {
		box.Push(util.Point{X: position.X - 1, Y: position.Y})
	}

	return box
}

// Pushes the boxes in the boxStacks in the direction of the next instruction.
func (w *Warehouse) pushBoxes(boxStacks BoxStackStack, next util.Point) {
	temp := make([][]int, len(w.Map))
	for i, row := range w.Map {
		temp[i] = slices.Clone(row)
	}
	for len(boxStacks) > 0 {
		boxstack := boxStacks.Pop()
		for len(boxstack) > 0 {
			box := boxstack.Pop()
			w.Map[box.Y][box.X] = FREE
			w.Map[box.Y+next.Y][box.X+next.X] = temp[box.Y][box.X]
		}
	}
}

// Draws the warehouse to the console.
func (w Warehouse) Draw() {
	for y := 0; y < len(w.Map); y++ {
		for x := 0; x < len(w.Map[y]); x++ {
			switch w.Map[y][x] {
			case WALL:
				print("#")
			case BOX:
				print("O")
			case BOX_LEFT:
				print("[")
			case BOX_RIGHT:
				print("]")
			case FREE:
				if w.Robot.X == x && w.Robot.Y == y {
					print("@")
				} else {
					print(".")
				}
			}
		}
		println()
	}
}

/* ----------------------------- Helper Methods ----------------------------- */

// Parses the input string into a Warehouse struct.
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
				warehouseMap[y][x] = FREE
			}
		}
	}

	return Warehouse{
		Map:          warehouseMap,
		Robot:        robot,
		Instructions: parseInstructions(instructionsString)}
}

// Parses the instructions string into an array of position deltas.
func parseInstructions(input string) []util.Point {
	instructions := []util.Point{}

	input = strings.ReplaceAll(input, "\n", "")
	for _, direction := range strings.Split(input, "") {
		instructions = append(instructions, Directions[direction])
	}
	return instructions
}

/* --------------------- BoxStacks Definition and Methods -------------------- */

// BoxStack is a stack of points representing boxes.
type BoxStack []util.Point

// Push a point onto the stack.
func (s *BoxStack) Push(v util.Point) {
	*s = append(*s, v)
}

// Pop a point from the stack.
func (s *BoxStack) Pop() util.Point {
	l := len(*s)
	val := (*s)[l-1]
	*s = (*s)[:l-1]
	return val
}

// Peek at the top of the stack.
func (s BoxStack) Peek() util.Point {
	l := len(s)
	return s[l-1]
}

// BoxStackStack is a stack of BoxStacks.
type BoxStackStack []BoxStack

// Push a BoxStack onto the stack.
func (s *BoxStackStack) Push(v BoxStack) {
	*s = append(*s, v)
}

// Pop a BoxStack from the stack.
func (s *BoxStackStack) Pop() BoxStack {
	l := len(*s)
	val := (*s)[l-1]
	*s = (*s)[:l-1]
	return val
}

// Peek at the top of the stack.
func (s BoxStackStack) Peek() BoxStack {
	l := len(s)
	return s[l-1]
}
