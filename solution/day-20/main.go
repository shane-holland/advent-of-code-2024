/* -------------------------------------------------------------------------- */
/*                           Day 20: Race Condition                           */
/* -------------------------------------------------------------------------- */
package day20

import (
	"fmt"
	"math"

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
	raceTrack := NewRaceTrack(input)

	threshold := 100
	if len(raceTrack.Track) < 100 {
		threshold = 50
	}

	return fmt.Sprintf("%v", raceTrack.CountCheatDistances(2, threshold))
}

// Part 2:
func part2(input string) string {
	raceTrack := NewRaceTrack(input)

	threshold := 100
	if len(raceTrack.Track) < 100 {
		threshold = 50
	}

	return fmt.Sprintf("%v", raceTrack.CountCheatDistances(20, threshold))
}

type RaceTrack struct {
	Track map[util.Point]int
	Walls []util.Point
	Start util.Point
	End   util.Point
	Bounds int
}

func (rt RaceTrack) GetCheatDistances() map[util.Point]int {
	cheatDistances := make(map[util.Point]int)
	for _, wall := range rt.Walls {
		dist := 0
		for _, d := range Neighbors {
			p1 := util.Point{X: wall.X + d.X, Y: wall.Y + d.Y}
			if v1, ok := rt.Track[p1]; ok {
				for _, d2 := range Neighbors {
					p2 := util.Point{X: wall.X + d2.X, Y: wall.Y + d2.Y}
					if v2, ok := rt.Track[p2]; ok {
						dist = int(math.Max(float64(dist), float64(util.AbsInt(v1-v2))))
					}
				}
			}
		}
		if dist > 2 {
			cheatDistances[wall] = dist - 2
		}

	}
	return cheatDistances
}

func (rt RaceTrack) CountCheatDistances(duration, threshold int) int {
	distances := make(map[cheatpath]int)
	duration = duration - 1

	for _, wall := range rt.Walls {
		for _, d := range Neighbors {
			p1 := util.Point{X: wall.X + d.X, Y: wall.Y + d.Y}
			if v1, ok := rt.Track[p1]; ok {
				// Diamond search pattern
				for y:=p1.Y -duration; y <= p1.Y + duration; y++ {
					if y >= 0 && y < rt.Bounds {
						for x := p1.X-(y-(p1.Y - duration)); x <= p1.X + (y-(p1.Y - duration)); x++ {
							if x >= 0 && x < rt.Bounds {
								p2 := util.Point{X: x, Y: y}
								adjust := util.ManhattanDistance(p1,p2) 
								if v2, ok := rt.Track[p2]; ok && adjust <= threshold && (v2 - v1) - adjust >= threshold {
									if delta, ok := distances[cheatpath{p1.X, p1.Y, x, y}]; ok {
										distances[cheatpath{p1.X, p1.Y, x, y}] = int(math.Max(float64(delta), float64((v2 - v1)-adjust)))
									} else {
										distances[cheatpath{p1.X, p1.Y, x, y}] = (v2 - v1)-adjust
									}
								}
							}

						}

					}
				}
			}
		}

	}
	// dCount := make(map[int]int) 
	// for _, d := range distances {
	// 	dCount[d]++
	// }
	

	return len(distances)
}

type cheatpath struct {
	x1 int
	y1 int
	x2 int
	y2 int
}



var Neighbors = []util.Point{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}

func NewRaceTrack(input string) RaceTrack {
	grid := util.GetLines(input)
	bounds := len(grid)
	// Map candidate walls and start/end positions
	walls, start, end := []util.Point{}, util.Point{}, util.Point{}
	for y, row := range grid {
		for x, cell := range row{
			switch cell {
			case '#':
				// Check for neighboring paths
				for _, d := range Neighbors {
					if x+d.X >= 0 && x+d.X < bounds && y+d.Y >=0 && y+d.Y < bounds {
						if grid[y+d.Y][x+d.X] != '#' {
							walls = append(walls, util.Point{X: x, Y: y})
						}
					}
				}

			case 'S':
				start = util.Point{X: x, Y: y}
			case 'E':
				end = util.Point{X: x, Y: y}
			}
		}
	}

	track := make(map[util.Point]int)
	node := start
	d := 0

	for node != end {
		track[node] = d
		d++
		for _, d := range Neighbors {
			testPoint := util.Point{X: node.X + d.X, Y: node.Y + d.Y}
			if _, ok := track[testPoint]; !ok && grid[testPoint.Y][testPoint.X] != '#' {
				node = testPoint
				break
			}
		}
	}
	track[end] = d

	return RaceTrack{Track: track, Walls: walls, Start: start, End: end, Bounds: bounds}
}

