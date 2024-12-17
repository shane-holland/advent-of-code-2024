/* -------------------------------------------------------------------------- */
/*                       --- Day 6: Guard Gallivant ---                       */
/* -------------------------------------------------------------------------- */
package day06

import (
	"fmt"
	"slices"
	"strings"

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
	return fmt.Sprintf("%d", patrolMap.PositionsWhichCauseALoop())
}

/* -------------------- PatrolMap Definition and Methods -------------------- */

// PatrolMap represents a map of a guard's patrol path.
// It contains the guard's position, obstacles, direction, and bounds.
type PatrolMap struct {
	GuardPosition util.Point
	Obstacles     []util.Point
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
		if slices.Contains(pm.Obstacles, pm.GuardPosition) {
			pm.Direction = (pm.Direction + 1) % 4
			pm.GuardPosition = lastPos
			continue
		}

		// Loop check
		if dir, ok := visited[pm.GuardPosition]; ok && dir == pm.Direction {
			return make([]util.Point, 0)
		}

		if pm.GuardPosition.X < 0 || pm.GuardPosition.Y < 0 || pm.GuardPosition.X >= pm.Bounds.X || pm.GuardPosition.Y >= pm.Bounds.Y {
			points := make([]util.Point, 0)
			for p := range visited {
				points = append(points, p)
			}
			return points
		}
		visited[pm.GuardPosition] = pm.Direction
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
func (pm *PatrolMap) PositionsWhichCauseALoop() int {
	loopObstacles := make([]util.Point, 0)
	original := PatrolMap{pm.GuardPosition, pm.Obstacles, pm.Direction, pm.Bounds}

	// Resets the patrol map for the next test
	reset := func() PatrolMap {
		return PatrolMap{original.GuardPosition, original.Obstacles, original.Direction, original.Bounds}
	}

	// Only test positions we know the guard will normally visit
	testPositions := pm.PointsVisited()
	for _, pos := range testPositions {
		// Reset the map
		*pm = reset()
		pm.Obstacles = append(pm.Obstacles, util.Point{X: pos.X, Y: pos.Y})
		// Loop discovered
		if pm.CountPointsVisited() == -1 {
			loopObstacles = append(loopObstacles, util.Point{X: pos.X, Y: pos.Y})
		}
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
	obstacles := []util.Point{}
	bounds := len(util.GetLines(input))

	for y, line := range util.GetLines(input) {
		if strings.Contains(line, "^") {
			pos = util.Point{X: strings.Index(line, "^"), Y: y}
		}
		offset := 0
		for strings.Contains(line, "#") {
			x := strings.Index(line, "#")
			obstacles = append(obstacles, util.Point{X: x + offset, Y: y})
			line = line[x+1:]
			offset += x + 1
		}
	}

	return PatrolMap{
		GuardPosition: pos,
		Obstacles:     obstacles,
		Direction:     NORTH,
		Bounds:        util.Point{X: bounds, Y: bounds},
	}
}
