/* -------------------------------------------------------------------------- */
/*                               Day 18: RAM Run                              */
/* -------------------------------------------------------------------------- */
package day18

import (
	"fmt"
	"math"
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

// Part 1: Calculate the minimum number of steps needed to reach the exit
func part1(input string) string {
	memoryGrid := NewMemoryGrid(input)
	bytes := 12 // Number of bytes to fall for test case
	if memoryGrid.Bounds > 10 {
		bytes = 1024 // Number of bytes to fall for puzzle input
	}

	for i := 0; i < bytes; i++ {
		memoryGrid.PushNextByte()
	}

	return fmt.Sprintf("%d", memoryGrid.ShortestPath())
}

// Part 2: Calculate coordinates of the first byte that will prevent the exit from being reachable from your starting position
func part2(input string) string {
	memoryGrid := NewMemoryGrid(input)
	// Get the first byte which blocks the path
	position := memoryGrid.FirstBlockingByte()
	// Convert to X,Y coordinates
	x := position % memoryGrid.Bounds
	y := int(math.Floor(float64(position) / float64(memoryGrid.Bounds)))

	return fmt.Sprintf("%d,%d", x, y)
}

/* -------------------- MemoryGrid Definition and Methods ------------------- */

// MemoryGrid which is a graph object describing vertices and edges, and a queue of incoming edges to destroy
type MemoryGrid struct {
	Graph    map[int][]int
	Bounds   int
	Incoming []int
}

// Return the shortest path from (0,0) to (bounds-1, bounds-1)
// This is a fairly simple BFS traversal
func (mg MemoryGrid) ShortestPath() int {
	end := (mg.Bounds * mg.Bounds) - 1
	// seen is a map of visited and distance travelled to get there
	seen := map[int]int{0: 0}
	queue := []int{0}

	// BFS search for end
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		d := seen[cur]

		if cur == end {
			return d
		}

		for _, edge := range mg.Graph[cur] {
			if _, ok := seen[edge]; !ok {
				seen[edge] = d + 1
				queue = append(queue, edge)
			}
		}
	}

	// End not reachable
	return -1
}

// Remove edges to the next item in the Incoming queue
func (mg *MemoryGrid) PushNextByte() {
	next := mg.Incoming[0]
	mg.Incoming = mg.Incoming[1:]

	// Remove edges which connect to this byte
	for _, v := range mg.Graph[next] {
		i := slices.Index(mg.Graph[v], next)
		mg.Graph[v] = slices.Delete(mg.Graph[v], i, i+1)
	}
}

// Retrieve the first item in the Incoming queue which makes traversal to the end impossible
//
//	To solve this, I created a stack from the original Incoming queue, to process it in reverse order
//	For each item in the stack, I rebuild the removed edges and test using ShortestPath to see if it's possible
//	to traverse the graph.  The first item which successfully traverses, is also the first byte which makes it impossible
//	to reach the end.
func (mg *MemoryGrid) FirstBlockingByte() int {
	stack := slices.Clone(mg.Incoming)

	// Push all the bytes
	for len(mg.Incoming) > 0 {
		mg.PushNextByte()
	}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Rebuild the edges
		for _, edge := range mg.Graph[cur] {
			mg.Graph[edge] = append(mg.Graph[edge], cur)
		}

		if mg.ShortestPath() != -1 {
			return cur
		}
	}
	return -1
}

/* ----------------------------- Helper Methods ----------------------------- */

func NewMemoryGrid(input string) MemoryGrid {
	lines := util.GetLines(input)
	incoming := make([]int, len(lines))
	// Determine bounds based on the data set
	bounds := 71 // Bounds for puzzle input
	if len(lines) < 1024 {
		bounds = 7 // Bounds for test data
	}

	// Parse incoming bytes
	for i, line := range lines {
		incoming[i] = parseMemoryAddress(line, bounds)
	}

	return MemoryGrid{Graph: generateGraph(bounds), Bounds: bounds, Incoming: incoming}

}

// Returns the address parsed from a line of input
func parseMemoryAddress(line string, bounds int) int {
	coords := strings.Split(line, ",")
	return mapMemoryAddress(util.AtoI(coords[0]), util.AtoI(coords[1]), bounds)
}

// Return an integer address representing x and y coordinates on a grid
func mapMemoryAddress(x, y, bounds int) int {
	return x + (y * bounds)
}

// Generates a graph of points in a bounds x bounds grid, each connected to their neighbors to the North, South, East, and West
//
//	Vertices are connected bi-directionally
func generateGraph(bounds int) map[int][]int {
	

	graph := make(map[int][]int)
	// Initialize the map
	for i := 0; i < bounds*bounds; i++ {
		graph[i] = make([]int, 0)
	}
	for y := 0; y < bounds; y++ {
		for x := 0; x < bounds; x++ {
			node := mapMemoryAddress(x, y, bounds)
			// For each valid neighbor, generate an edge
			for _, neighbor := range getEdges(x, y, bounds) {
				if !slices.Contains(graph[node], neighbor) {
					graph[node] = append(graph[node], neighbor)
					graph[neighbor] = append(graph[neighbor], node)
				}
			}
		}
	}

	return graph
}

// Returns true if the x,y position is within a grid of bounds * bounds dimensions
func inBounds(x, y, bounds int) bool {
	return x >= 0 && x < bounds && y >= 0 && y < bounds
}

// Returns all neighbors (up, down, left, right) which are in bounds
func getEdges(x,y,bounds int) []int {
	edges := []int{}
	// For each valid neighbor, generate an edge
	for _, d := range []util.Point{{X: 0, Y: 1}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: -1, Y: 0}} {
		if inBounds(x+d.X, y+d.Y, bounds) {
			edges = append(edges, mapMemoryAddress(x+d.X, y+d.Y, bounds))
		}
	}
	return edges
}
