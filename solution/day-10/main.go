package day10

import (
	"fmt"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

/**
 * Function to find the minimum of two integers.
 */
func part1(input string) string {
	grid := parseGrid(input)
	trailMap := NewTopographicMap(grid)

	scoreSum := 0
	for _, trailhead := range trailMap.TrailHeads {
		trailScore := trailMap.score(trailhead)
		scoreSum += trailScore
	}
	return fmt.Sprintf("%d", scoreSum)
}

/**
 * Function to find the minimum of two integers.
 */
func part2(input string) string {
	grid := parseGrid(input)
	trailMap := NewTopographicMap(grid)

	ratingSum := 0
	for _, trailhead := range trailMap.TrailHeads {
		trailScore := trailMap.rating(trailhead)
		ratingSum += trailScore
	}

	return fmt.Sprintf("%d", ratingSum)
}

type Plot struct {
	X      int
	Y      int
	Height int
}

type TopographicMap struct {
	TrailHeads  []Plot
	TrailEdges  map[Plot]map[Plot]int
	TrailBounds int
}

func (t TopographicMap) score(start Plot) int {
	parents := make(map[Plot]Plot)

	score := 0
	queue := []Plot{start}
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for neighbor, weight := range t.TrailEdges[node] {
			if _, ok := parents[neighbor]; weight > 0 && !ok {
				parents[neighbor] = node
				queue = append(queue, neighbor)
			}
		}
		if node.Height == 9 {
			score += 1
		}
	}
	return score
}

func (t TopographicMap) rating(start Plot) int {
	rating := 0
	queue := []Plot{start}
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for neighbor, weight := range t.TrailEdges[node] {
			if weight > 0 {
				queue = append(queue, neighbor)
			}
		}
		if node.Height == 9 {
			rating += 1
		}
	}
	return rating
}

func NewTopographicMap(grid [][]int) TopographicMap {
	bounds := len(grid)
	trailheads := make([]Plot, 0)
	edges := make(map[Plot]map[Plot]int)

	// Build graph
	for y, row := range grid {
		for x, height := range row {
			if height == 0 {
				trailheads = append(trailheads, Plot{x, y, height})
			}
			node := Plot{x, y, height}
			edges[node] = getEdges(node, grid)
		}
	}

	return TopographicMap{trailheads, edges, bounds}
}

func getEdges(node Plot, grid [][]int) map[Plot]int {
	bounds := len(grid)
	edges := make(map[Plot]int)

	for i := -1; i <= 1; i += 2 {
		if node.X+i >= 0 && node.X+i < bounds && grid[node.Y][node.X+i] == node.Height+1 {
			edges[Plot{node.X + i, node.Y, grid[node.Y][node.X+i]}] = 1
		}
		if node.Y+i >= 0 && node.Y+i < bounds && grid[node.Y+i][node.X] == node.Height+1 {
			edges[Plot{node.X, node.Y + i, grid[node.Y+i][node.X]}] = 1
		}
	}

	return edges
}

func parseGrid(data string) [][]int {
	grid := make([][]int, 0)
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		current := make([]int, 0)
		for _, col := range row {
			current = append(current, util.AtoI(string(col)))
		}
		grid = append(grid, current)
	}
	return grid
}
