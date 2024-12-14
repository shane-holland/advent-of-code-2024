/* -------------------------------------------------------------------------- */
/*                      --- Day 14: Restroom Redoubt ---                      */
/* -------------------------------------------------------------------------- */
package day14

import (
	"fmt"
	"math"
	"regexp"
	"slices"

	"shaneholland.dev/aoc-2024/util"
)

/* ------------------------------- Main Method ------------------------------ */
type Puzzle struct{}

// The Solve method is called to solve the puzzle.
func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

/* -------------------------------- Solution -------------------------------- */

// Part 1: Calculate the Safety Factor of the lobby after 100 seconds. 
func part1(input string) string {
	lobby := parseLobby(input)
	lobby.Update(100)
	return fmt.Sprintf("%d", lobby.SafetyFactor())
}

// Part 2: Determine the number of seconds it takes for the robots to form a Christmas tree.
func part2(input string) string {
	// After watching for a pattern, we noticed that, starting at 97 seconds, 
	// a vertical formation appears every 101 seconds.  We updated our script to draw the lobby
	// every 101 seconds starting at 97 seconds.  After watching that for a while we 
	// saw the first Christmas tree appear at 7672 seconds!

	// We continued to watch until the next tree appeared, at 18075 seconds.
	// This means that the pattern repeats every 10403 seconds (Bounds.X * Bounds.Y).

	// Another way it seems that we can solve this is by looking for minimum safety factor.
	// At the point where the Tree appears, the safety factor is at its minimum.
	lobby := parseLobby(input)

	minSafetyFactor := math.Inf(1)
	secondsAtMinimumSafetyFactor := 0
	for i:=0; i <= lobby.Bounds.X * lobby.Bounds.Y; i++ {
		lobby.Update(i)
		safetyFactor := lobby.SafetyFactor()

		if safetyFactor < int(minSafetyFactor) {
			minSafetyFactor = float64(safetyFactor)
			secondsAtMinimumSafetyFactor = i
		}
	}
	return fmt.Sprintf("%d", secondsAtMinimumSafetyFactor)
}

/* ---------------------- Lobby Definition and Methods ---------------------- */

// A Lobby containing robots moving in straight lines.
type Lobby struct {
	Robots []Robot
	Bounds util.Point
	Positions []util.Point
}

// Update the positions of the robots after a given number of seconds.
func (l *Lobby) Update(seconds int) {
	l.Positions = make([]util.Point, len(l.Robots))
	for i, r := range l.Robots {
		l.Positions[i] = l.getWrappedPosition(r.move(seconds))
	}
}

// Return the product of the number of robots in each quadrant.
func (l Lobby) SafetyFactor() int {
	quadrants := []int{0, 0, 0, 0}
	for _, p := range l.Positions {
		if p.X < (l.Bounds.X-1)/2 && p.Y < (l.Bounds.Y-1)/2 {
			quadrants[0]++
		} else if  p.X < (l.Bounds.X-1)/2 && p.Y > (l.Bounds.Y-1)/2 {
			quadrants[1]++
		} else if  p.X > (l.Bounds.X-1)/2 && p.Y > (l.Bounds.Y-1)/2 {
			quadrants[2]++
		} else if  p.X > (l.Bounds.X-1)/2 && p.Y < (l.Bounds.Y-1)/2 {
			quadrants[3]++
		}
	}	
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

// Get the wrapped position of a point within the bounds.
// If the point is outside the bounds, it will wrap around to the other side.
func (l Lobby) getWrappedPosition(pos util.Point) util.Point{
	var x,y int
	if pos.X < 0 {
		x = (l.Bounds.X - (util.AbsInt(pos.X) % l.Bounds.X)) % l.Bounds.X
	}
	if pos.Y < 0 {
		y = (l.Bounds.Y - (util.AbsInt(pos.Y) % l.Bounds.Y)) % l.Bounds.Y
	}
	if pos.X >= 0 {
		x = pos.X % l.Bounds.X
	}
	if pos.Y >= 0 {
		y = pos.Y % l.Bounds.Y
	}
	return util.Point{X: x, Y: y}
}

func (l Lobby) ToString() string {
	output := ""
	for y:=0; y<l.Bounds.Y; y++ {
		for x:=0; x<l.Bounds.X; x++ {
			if slices.Contains(l.Positions, util.Point{X: x, Y: y}) {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

/* ---------------------- Robot Definition and Methods --------------------- */
// A Robot which starts at a point and moves in a direction at a constant speed.
type Robot struct {
	Start util.Point
	Vector util.Point
}

// Returns the position of the robot after given number of seconds has passed.
func (robot Robot) move(seconds int) util.Point {
	return util.Point {
		X: robot.Start.X + robot.Vector.X * seconds, 
		Y: robot.Start.Y + robot.Vector.Y * seconds,
	}
}

/* ---------------------------- Helper Functions ---------------------------- */

// Define the bounds of the lobby, which is different for the test input.
var TEST_BOUNDS = util.Point{X: 11, Y: 7}
var BOUNDS = util.Point{X: 101, Y: 103}

// Parse the input string into a Lobby containing robots moving in straight lines
func parseLobby(input string) Lobby {

	pattern := `p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	robots := make([]Robot, len(matches))
	positions := make([]util.Point, len(matches))
	for i, match := range matches {
		robots[i] = Robot{
			Start: util.Point{X: util.AtoI(match[1]), Y: util.AtoI(match[2])},
			Vector: util.Point{X: util.AtoI(match[3]), Y: util.AtoI(match[4])},
		}
		positions[i] = robots[i].Start
	}

	if (len(robots) < 50) {
		return Lobby{Robots: robots, Bounds: TEST_BOUNDS, Positions: positions}
	} 
	return Lobby{Robots: robots, Bounds: BOUNDS, Positions: positions}
}