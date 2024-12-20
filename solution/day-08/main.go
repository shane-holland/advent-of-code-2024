/* -------------------------------------------------------------------------- */
/*                    --- Day 8: Resonant Collinearity ---                    */
/* -------------------------------------------------------------------------- */
package day08

import (
	"fmt"
	"math"
	"regexp"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	antennaMap := parseAntennaMap(input)

	return part1(antennaMap), part2(antennaMap)
}

// Part 1: Count the number of antinodes in the antenna map.
func part1(antennaMap AntennaMap) string {
	antinodes := antennaMap.CountAntinodes(false)
	return fmt.Sprintf("%d", antinodes)
}

// Part 2: Count the number of antinodes in the antenna map with resonant harmonics.
func part2(antennaMap AntennaMap) string {
	antinodes := antennaMap.CountAntinodes(true)
	return fmt.Sprintf("%d", antinodes)
}

/* -------------------- AntennaMap Definition and Methods ------------------- */

// AntennaMap represents a map of antennas and their corresponding points.
type AntennaMap struct {
	Antennas map[string][]util.Point
	Bounds   util.Point
}

// CountAntinodes returns the number of antinodes in the antenna map.
// If resonantHarmonics is true, then the antinodes are calculated with resonant harmonics.
func (a *AntennaMap) CountAntinodes(resonantHarmonics bool) int {
	antinodes := make(map[util.Point]struct{})

	for _, antenna := range a.Antennas {
		for i, p1 := range antenna {
			for _, p2 := range antenna[i+1:] {
				if resonantHarmonics {
					for _, antinode := range a.getAllAntinodes(p1, p2) {
						antinodes[antinode] = struct{}{}
					}
				} else {
					for _, antinode := range a.getTwoAntinodes(p1, p2) {
						antinodes[antinode] = struct{}{}
					}
				}
			}

		}
	}

	return len(antinodes)
}

// pointInBounds returns true if the point is within the bounds of the antenna map.
func (a *AntennaMap) pointInBounds(point util.Point) bool {
	return point.X >= 0 && point.Y >= 0 && point.X < a.Bounds.X && point.Y < a.Bounds.Y
}

// getTwoAntinodes returns the two antinodes of two points.
func (a *AntennaMap) getTwoAntinodes(p1, p2 util.Point) []util.Point {
	antinodes := make([]util.Point, 0)

	node := util.Point{X: p1.X + (p1.X - p2.X), Y: p1.Y + (p1.Y - p2.Y)}
	if a.pointInBounds(node) {
		antinodes = append(antinodes, node)
	}
	node = util.Point{X: p2.X + (p2.X - p1.X), Y: p2.Y + (p2.Y - p1.Y)}
	if a.pointInBounds(node) {
		antinodes = append(antinodes, node)
	}
	return antinodes
}

// getAllAntinodes returns all antinodes between two points.
// The antinodes are calculated with resonant harmonics.
func (a *AntennaMap) getAllAntinodes(p1, p2 util.Point) []util.Point {
	antinodes := make([]util.Point, 0)
	antinodes = append(antinodes, []util.Point{p1, p2}...)

	factor := gcd((p1.X - p2.X), (p1.Y - p2.Y))
	offset1 := util.Point{X: (p1.X - p2.X) / factor, Y: (p1.Y - p2.Y) / factor}
	offset2 := util.Point{X: (p2.X - p1.X) / factor, Y: (p2.Y - p1.Y) / factor}
	point1 := util.Point{X: p1.X, Y: p1.Y}
	point2 := util.Point{X: p2.X, Y: p2.Y}
	for {
		point1 = util.Point{X: point1.X + offset1.X, Y: point1.Y + offset1.Y}
		point2 = util.Point{X: point2.X + offset2.X, Y: point2.Y + offset2.Y}
		if !a.pointInBounds(point1) && !a.pointInBounds(point2) {
			break
		}
		if a.pointInBounds(point1) {
			antinodes = append(antinodes, point1)
		}
		if a.pointInBounds(point2) {
			antinodes = append(antinodes, point2)
		}
	}

	return antinodes
}

/* ----------------------------- Helper Methods ----------------------------- */

// Parses an antenna map from the input string
func parseAntennaMap(input string) AntennaMap {
	antennas := make(map[string][]util.Point)
	inputLines := util.GetLines(input)

	pattern := `(\d|\w)`
	re := regexp.MustCompile(pattern)

	for y, line := range inputLines {
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			x := match[0]
			antennas[line[x:x+1]] = append(antennas[line[x:x+1]], util.Point{X: x, Y: y})
		}
	}
	return AntennaMap{Antennas: antennas, Bounds: util.Point{X: len(inputLines[0]), Y: len(inputLines)}}
}

// Returns the greatest common divisor of two numbers
func gcd(num1 int, num2 int) int {
	n1 := int(math.Max(math.Abs(float64(num1)), math.Abs(float64(num2))))
	n2 := int(math.Min(math.Abs(float64(num1)), math.Abs(float64(num2))))

	for {
		if n1%n2 == 0 {
			return n2
		}
		n1, n2 = n2, n1%n2
	}
}
