/* -------------------------------------------------------------------------- */
/*                        --- Day 16: Reindeer Maze ---                       */
/* -------------------------------------------------------------------------- */
package day16

import (
	"fmt"
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

// Part 1: What is the lowest score a Reindeer could get traversing from
//
//	Start (S) to End (E)?
func part1(input string) string {
	maze := NewMaze(input)
	return fmt.Sprintf("%d", maze.LowestScore())
}

// Part 2: How many tiles are part of at least one of the best paths through the maze?
func part2(input string) string {
	maze := NewMaze(input)
	return fmt.Sprintf("%d", maze.TilesOnBestPaths())
}

/* ----------------------- Maze Definition and Methods ---------------------- */

// The maze consists of a graph with start and end vertices
type Maze struct {
	Graph map[util.Point][]util.Point
	Start util.Point
	End   util.Point
}

// Return the lowest score possible when traversing from start to end
func (m Maze) LowestScore() int {
	return m.getLowestScores()[m.End]
}

// Return the number of tiles on the map which occur in any of the "best" paths
func (m Maze) TilesOnBestPaths() int {
	scoreMap := m.getLowestScores()
	queue := []util.Point{m.End}

	// Map of vertices on the best paths
	bestPathNodes := make(map[util.Point]struct{}, 0)
	bestPathNodes[m.End] = struct{}{}

	// This is a sort of BFS implementation
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// We've reached the start node, we've collected all the nodes which
		//   are part of the best paths
		if node == m.Start {
			return len(bestPathNodes)
		}

		// Find the node which is next from our current node.
		// We'll use this to determine which neighbor nodes are part of the path
		subqueue := make([]util.Point, 0)
		next := m.End
		for _, neighbor := range m.Graph[node] {
			if _, ok := bestPathNodes[neighbor]; ok {
				next = neighbor
				break
			}
		}
		// Loop through neighbors to collect those which are part of the best paths
		for _, neighbor := range m.Graph[node] {
			points := scoreMap[neighbor]
			if (points == scoreMap[node]-1) || (points == scoreMap[node]-1001) ||
				(points == scoreMap[next]-2) ||
				(points == scoreMap[next]-1002 && neighbor.X != next.X && neighbor.Y != next.Y) {
				// A difference of 1 or 1001 from the current node means this is part of the path
				// A difference of 2 from "next" means that the neighbor is in line and part of the path
				// A difference of 1002 from "next" and no shared axis means the neighbor is around
				// 	 the corner but part of the path
				bestPathNodes[neighbor] = struct{}{}
				subqueue = append(subqueue, neighbor)
			}
		}
		queue = append(queue, subqueue...)
	}
	return len(bestPathNodes)
}

// Return a map of vertices and the lowest score possible to reach them
func (m Maze) getLowestScores() map[util.Point]int {
	type Node struct {
		vertex    util.Point
		direction int
	}
	// Track each vertex and the points it took to get there
	visited := map[util.Point]int{m.Start: 0}
	// Our queue will need to keep track of our last direction of travel
	queue := []Node{{m.Start, EAST}}

	// Removes a node from the queue.  Used when we find a better path than is queued
	removeNodeFromQueue := func(node util.Point) {
		for i, n := range queue {
			if n.vertex == node {
				queue = slices.Delete(queue, i, i+1)
			}
		}
	}

	// BFS to traverse to the end node
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range m.Graph[node.vertex] {
			// Determine the direction this neighbor is from our vertex
			dir := Directions[util.Point{X: neighbor.X - node.vertex.X, Y: neighbor.Y - node.vertex.Y}]
			points := visited[node.vertex] + getPoints(node.direction, dir)

			// If this node is already in the queue, we need to see if this path is better
			if n, ok := visited[neighbor]; ok {
				if points > n {
					// There's a better path to get here, ignore this one
					continue
				}
				// The path in the queue is worse, remove it
				removeNodeFromQueue(neighbor)
			}
			// Update visisted and add the next node to the queue
			visited[neighbor] = points
			queue = append(queue, Node{neighbor, dir})
		}

	}

	return visited
}

/* ----------------------------- Helper Methods ----------------------------- */

// Directions
const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

// Direction Map: Retrieve the direction of one node to another by
//
//		subtracting the Source nodes coordinates from the Target node and looking up
//	 the result in the map.
var Directions = map[util.Point]int{
	{X: 0, Y: -1}: NORTH,
	{X: 1, Y: 0}:  EAST,
	{X: 0, Y: 1}:  SOUTH,
	{X: -1, Y: 0}: WEST,
}

// NewMaze creates a new Maze from the input string.
func NewMaze(input string) *Maze {
	maze := &Maze{Graph: make(map[util.Point][]util.Point)}
	grid := util.GetLines(input)

	for y, row := range grid {
		for x, cell := range row {
			node := util.Point{X: x, Y: y}
			switch cell {
			case 'S':
				maze.Start = node
			case 'E':
				maze.End = node
			case '#':
				continue
			}
			maze.Graph[node] = getEdges(node, grid)

		}
	}
	return maze
}

// Return a list of edges for a given node in the grid
func getEdges(node util.Point, grid []string) []util.Point {
	edges := make([]util.Point, 0)
	for delta := range Directions {
		if grid[node.Y+delta.Y][node.X+delta.X] != '#' {
			edges = append(edges, util.Point{X: node.X + delta.X, Y: node.Y + delta.Y})
		}
	}
	return edges
}

// Returns the number of points accumulated based on the direction the next node
//
//	related to the current node, and the previous direction of travel
func getPoints(lastDir, newDir int) int {
	// Traversing one space, add one point
	points := 1
	if util.AbsInt(newDir-lastDir)%2 == 1 {
		// This is a 90 degree turn, add 1000 points for turning
		points += 1000
	}
	return points
}
