/* -------------------------------------------------------------------------- */
/*                       --- Day 6: Guard Gallivant ---                       */
/* -------------------------------------------------------------------------- */
package day06

import (
	"fmt"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

// Part 1: Find the number of points visited before the guard leaves the area
func part1(input string) string {
	patrolMap := parsePatrolMap(input)
	return fmt.Sprintf("%d", patrolMap.CountPointsVisited())
}

// Part 2: Find the number of obstacles that can cause the guard to loop
func part2(input string) string {
	patrolMap := parsePatrolMap(input)
	return fmt.Sprintf("%d", patrolMap.CountPositionsWhichCauseALoop())
}

/* -------------------- PatrolMap Definition and Methods -------------------- */

// PatrolMap represents a map of a guard's patrol path.
// It contains the guard's position, obstacles, direction, and bounds.
type PatrolMap struct {
	GuardPosition util.Point
	Grid          [][]bool
	Direction     int
	Bounds        util.Point
}

// PointsVisited returns the number of points visited before the guard leaves the area.
func (pm *PatrolMap) CountPointsVisited() int {
	visited := pm.PointsVisited()
	if len(visited) > 0 {
		return len(visited)
	}
	// Loop encountered
	return -1
}

// Returns the unique set of points visisted by the guard
func (pm *PatrolMap) PointsVisited() []util.Point {
	visited := map[util.Point]int{pm.GuardPosition: pm.Direction}

	for {
		lastPos := pm.GuardPosition
		pm.Move()

		if pm.GuardPosition.X < 0 || pm.GuardPosition.Y < 0 || pm.GuardPosition.X >= pm.Bounds.X || pm.GuardPosition.Y >= pm.Bounds.Y {
			points := make([]util.Point, 0)
			for p := range visited {
				points = append(points, p)
			}
			return points
		}

		if pm.Grid[pm.GuardPosition.Y][pm.GuardPosition.X] {
			pm.Direction = (pm.Direction + 1) % 4
			pm.GuardPosition = lastPos
			continue
		}

		// Loop check
		if dir, ok := visited[pm.GuardPosition]; ok && dir == pm.Direction {
			return make([]util.Point, 0)
		}

		visited[pm.GuardPosition] = pm.Direction
	}
}

// Returns the unique set of points visisted by the guard
func (pm *PatrolMap) LoopCheck() bool {
	visited := map[util.Point]int{pm.GuardPosition: pm.Direction}

	for {
		lastPos := pm.GuardPosition
		pm.Move()

		if pm.GuardPosition.X < 0 || pm.GuardPosition.Y < 0 || pm.GuardPosition.X >= pm.Bounds.X || pm.GuardPosition.Y >= pm.Bounds.Y {
			return false
		}

		if pm.Grid[pm.GuardPosition.Y][pm.GuardPosition.X] {
			// Loop check
			if dir, ok := visited[pm.GuardPosition]; ok && dir == pm.Direction {
				return true
			}

			visited[pm.GuardPosition] = pm.Direction
			pm.Direction = (pm.Direction + 1) % 4
			pm.GuardPosition = lastPos
			continue
		}
	}
}

// Move the guard one space in the current direction.
func (pm *PatrolMap) Move() {
	switch pm.Direction {
	case NORTH:
		pm.GuardPosition.Y--
	case EAST:
		pm.GuardPosition.X++
	case SOUTH:
		pm.GuardPosition.Y++
	case WEST:
		pm.GuardPosition.X--
	}
}

// PositionsWhichCauseALoop returns the number of obstacle positions that can cause the guard to loop.
func (pm *PatrolMap) CountPositionsWhichCauseALoop() int {
	loopObstacles := make([]util.Point, 0)
	originalPosition := pm.GuardPosition

	// Only test positions we know the guard will normally visit
	testPositions := pm.PointsVisited()
	for _, pos := range testPositions {
		// Reset the map
		pm.GuardPosition = originalPosition
		pm.Direction = NORTH
		// Add the obstacle
		pm.Grid[pos.Y][pos.X] = true
		// Loop discovered
		if pm.LoopCheck() {
			loopObstacles = append(loopObstacles, util.Point{X: pos.X, Y: pos.Y})
		}
		pm.Grid[pos.Y][pos.X] = false
	}

	return len(loopObstacles)
}

/* ----------------------------- Helper Methods ----------------------------- */

// Directions
const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

// parsePatrolMap returns a PatrolMap from the input string.
func parsePatrolMap(input string) PatrolMap {
	pos := util.Point{X: 0, Y: 0}
	bounds := len(util.GetLines(input))
	grid := make([][]bool, bounds)

	for y, line := range util.GetLines(input) {
		grid[y] = make([]bool, len(line))
		for x, cell := range line {
			switch cell {
			case '^':
				pos = util.Point{X: x, Y: y}
			case '#':
				grid[y][x] = true
			}
		}
	}

	return PatrolMap{
		GuardPosition: pos,
		Grid:          grid,
		Direction:     NORTH,
		Bounds:        util.Point{X: bounds, Y: bounds},
	}
}
