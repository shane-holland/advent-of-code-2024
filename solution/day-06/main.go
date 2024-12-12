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

/**
 * Find the number of spaces visited by the patrol.
 */
func part1(input string) string {
	patrolMap := parsePatrolMap(input)
	return fmt.Sprintf("%d", patrolMap.PointsVisited())
}

/**
 * Find the number of obstacle placements that cause a loop.
 */
func part2(input string) string {
	patrolMap := parsePatrolMap(input)
	loopObstacles := make([]util.Point, 0)

	for y := 0; y < patrolMap.Bounds.Y; y++ {
		for x := 0; x < patrolMap.Bounds.X; x++ {
			if (!slices.Contains(patrolMap.Obstacles, util.Point{X: x, Y: y})) {
				// Reset the map
				patrolMap := parsePatrolMap(input)
				patrolMap.Obstacles = append(patrolMap.Obstacles, util.Point{X: x, Y: y})
				// Loop discovered
				if patrolMap.PointsVisited() == -1 {
					loopObstacles = append(loopObstacles, util.Point{X: x, Y: y})
				}
			}
		}
	}

	return fmt.Sprintf("%d", len(loopObstacles))
}

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

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

type PatrolMap struct {
	GuardPosition util.Point
	Obstacles     []util.Point
	Direction     int
	Bounds        util.Point
}

func (pm *PatrolMap) PointsVisited() int {
	visited := make(map[util.Point]int)

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
			return -1
		}

		if pm.GuardPosition.X < 0 || pm.GuardPosition.Y < 0 || pm.GuardPosition.X >= pm.Bounds.X || pm.GuardPosition.Y >= pm.Bounds.Y {
			return len(visited)
		}
		visited[pm.GuardPosition] = pm.Direction
	}
}

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
