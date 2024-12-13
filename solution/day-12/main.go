/* -------------------------------------------------------------------------- */
/*                        --- Day 12: Garden Groups ---                       */
/* -------------------------------------------------------------------------- */
package day12

import (
	"fmt"
	"slices"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}


// Part 1: Return the cost of fencing in the garden. (Perimeter * Area)
func part1(input string) string {
	garden := parseGarden(input)
	return fmt.Sprintf("%d", garden.GetFencingCost(false))
}


// Part 2: Return the cost of fencing in the garden with a bulk discount. (Sides * Area) 
func part2(input string) string {
	garden := parseGarden(input)
	return fmt.Sprintf("%d", garden.GetFencingCost(true))
}

/* ---------------------- Garden Definition and Methods --------------------- */

// Garden represents a collection of regions.
// Each region is a collection of plots.
type Garden struct {
	Regions []Region
}

// GetFencingCost returns the cost of fencing in the garden.
// If bulkDiscount is true, then the cost is calculated as Sides * Area.
// Otherwise, the cost is calculated as Perimeter * Area.
func (g *Garden) GetFencingCost(bulkDiscount bool) int {
	cost := 0
	for _, region := range g.Regions {
		cost += region.FencingCost(bulkDiscount)
	}
	return cost
}

/* ---------------------- Region Definition and Methods --------------------- */

/**
 * Region represents a collection of plots.
 * Each plot is a point on a grid.
 */
type Region struct {
	Plots []util.Point
	Graph map[util.Point][]util.Point
}

// FencingCost returns the cost of fencing in the region.
// If bulkDiscount is true, then the cost is calculated as Sides * Area.
// Otherwise, the cost is calculated as Perimeter * Area.
func (region Region) FencingCost(bulkDiscount bool) int {
	if bulkDiscount {
		return region.CountSides() * region.GetArea()
	} 
	return region.GetPerimeter() * region.GetArea()
}

// GetArea returns the number of plots in the region.
func (r Region) GetArea() int {
	return len(r.Plots)
}

// GetPerimeter returns the number of plots on the edge of the region.
func (r Region) GetPerimeter() int {
	perimeter := 0
	for _, edges := range r.Graph {
		// If the length of the edges is less than 4, then the plot is on the edge of the region
		perimeter += 4 - len(edges)
	}
	return perimeter
}

// CountSides returns the number of sides in the region.
func (r Region) CountSides() int {
	// The number of sides will always be equal to the number of corners of the region
	corners := 0

	for node := range r.Graph {
		corners += r.countOuterCorners(node)
		corners += r.countInteriorCorners(node)
	}

	return corners
}

// countOuterCorners returns the number of outward facing corners the node represents on the region.
// Example: where O is the node, # is a plot, and . is a plot outside the region.
//     . . . .
//     . O # #			This node is an outer corner
//     . # # #
func (r Region) countOuterCorners(node util.Point) int {
	edges := r.Graph[node]
	if len(edges) == 0 {
		return 4
	} else if len(edges) == 1 {
		return 2
	} else if len(edges) == 2 {
		if edges[0].X != edges[1].X && edges[0].Y != edges[1].Y {
			return 1
		}
	}
	return 0
}

// countInteriorCorners returns the number of inward facing corners the node represents in the region.
// Example: where O is the node, # is a plot, and . is a plot outside the region.
//     . . . #
//     . . . #
//     . # # O
func (r Region) countInteriorCorners(node util.Point) int {
	edges := r.Graph[node]
	corners := 0

	if len(edges) < 2 {
		return 0
	}

	for i := -1; i <= 1; i += 2 {
		if slices.Contains(edges, util.Point{X: node.X + i, Y: node.Y}) {
			for j := -1; j <= 1; j += 2 {
				if slices.Contains(edges, util.Point{X: node.X, Y: node.Y + j}) && !slices.Contains(r.Plots, util.Point{X: node.X + i, Y: node.Y + j}) {
					corners += 1
				}
			}
		}
	}
	return corners
}

// GetSubRegions returns a list of sub-regions within the current region.
// Each sub-region is a connected component of plots within the region.
func (r Region) GetComponentRegions() []Region {
	subRegions := []Region{}
	untested := slices.Clone(r.Plots)

	for len(untested) > 0 {
		region := Region{Plots: []util.Point{}, Graph: map[util.Point][]util.Point{}}
		queue := []util.Point{untested[0]}
		untested = untested[1:]

		for len(queue) > 0 {
			plot := queue[0]
			queue = queue[1:]
			region.Plots = append(region.Plots, plot)
			region.Graph[plot] = r.Graph[plot]

			for _, edge := range r.Graph[plot] {
				if index := slices.Index(untested, edge); index >= 0 {
					queue = append(queue, edge)
					untested = append(untested[:index], untested[index+1:]...)
				}
			}
		}

		subRegions = append(subRegions, region)
	}

	return subRegions
}

/* ----------------------------- Helper Methods ----------------------------- */

// parseGarden parses the input string representing a garden and returns a Garden struct.
// Each character in the input represents a plot, and adjacent plots with the same character
// are considered connected.
func parseGarden(input string) Garden {
    lines := util.GetLines(input)
    bounds := len(lines)

    // inBounds checks if a point is within the garden bounds.
    inBounds := func(p util.Point) bool {
        return p.X >= 0 && p.X < bounds && p.Y >= 0 && p.Y < bounds
    }

    // getEdges returns the adjacent points of a given point that have the same character.
    getEdges := func(p util.Point) []util.Point {
        edges := []util.Point{}
        for _, delta := range []util.Point{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}} {
            adj := util.Point{X: p.X + delta.X, Y: p.Y + delta.Y}
            if inBounds(adj) && lines[p.Y][p.X] == lines[adj.Y][adj.X] {
                edges = append(edges, adj)
            }
        }
        return edges
    }

    regionMap := make(map[string]Region)

    for y, line := range lines {
        for x, char := range line {
            region := string(char)
            plot := util.Point{X: x, Y: y}
            if _, ok := regionMap[region]; !ok {
                regionMap[region] = Region{Plots: []util.Point{}, Graph: map[util.Point][]util.Point{}}
            }
            r := regionMap[region]
            r.Plots = append(r.Plots, plot)
            r.Graph[plot] = getEdges(plot)
            regionMap[region] = r
        }
    }

	// There may be multiple regions of the same type in the Garden.
	// Break them into their connected components.
	regions := []Region{}
	for _, region := range regionMap {
		regions = append(regions, region.GetComponentRegions()...)
	}

    return Garden{Regions: regions}
}
