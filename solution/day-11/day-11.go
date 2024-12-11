package day11

import (
	"fmt"
	"strconv"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}


/**
 * Count the number of new stones after 25 blinks.
 */
func part1(input string) string {
	stoneGraph := NewStoneGraph(input)
	return fmt.Sprintf("%d", stoneGraph.Blink(25))
}

/**
 * Count the number of new stones after 75 blinks.
 */
func part2(input string) string {
	stoneGraph := NewStoneGraph(input)
	return fmt.Sprintf("%d", stoneGraph.Blink(75))
}

/**
 * Creates a new StoneGraph from the given input.
 */
func NewStoneGraph(input string) StoneGraph {
	return StoneGraph{Stones: input, Graph: map[string][]string{"0": {"1"}}}
}

/**
 * Represents a directional graph of stones (vertices) and resulting stones from a single blink (edges).
 */
type StoneGraph struct {
	Stones string
	Graph map[string][]string
}

/**
 * Simulates the blinking of the stones for a given number of times.
 * Returns the number of stones which will exist after the given number of blinks.
 */
func (g *StoneGraph) Blink(times int) int{
	count := 0

	for _, stone := range strings.Split(g.Stones, " ") {
		count += g.EdgesAfterSteps(stone, times)
	}

	return count
}

/**
 * Retrieves the edges for a given node, or generates them if they don't exist.
 * Rules:
 * 	- If the stone is engraved with the number 0, it is replaced by a 
 *		stone engraved with the number 1.
 * 	- If the stone is engraved with a number that has an even number of digits, 
 *		it is replaced by two stones. The left half of the digits are engraved 
 *		on the new left stone, and the right half of the digits are engraved on 
 *		the new right stone. (The new numbers don't keep extra leading zeroes: 
 *		1000 would become stones 10 and 0.)
 * 	- If none of the other rules apply, the stone is replaced by a new stone; 
 *		the old stone's number multiplied by 2024 is engraved on the new stone.
 */
func (g *StoneGraph) getEdges(node string) []string {
	if edges, ok := g.Graph[node]; ok {
		// Edges exist, return them
		return edges
	} 
	// Generate the edges
	edges := []string{}
	if len(node) % 2 == 0 {
	   mid := len(node) / 2
	   edges = append(edges, node[:mid], strconv.Itoa(util.AtoI(node[mid:])))
   	} else {
	   edges = append(edges, strconv.Itoa(util.AtoI(node) * 2024))
   	}
   	if len(node) <= 4 {
	   g.Graph[node] = edges
   	}
	return edges
}

/**
 * Returns the number of stones which will exist after a given number of blinks, 
 * starting from a given stone.
 */
func (g *StoneGraph) EdgesAfterSteps(start string, blinks int) int {
	next := map[string]int{start: 1}
	for i := 0; i < blinks; i++ {
		queue := next
		next = make(map[string]int)
		for node, count := range queue {
			for _, edge := range g.getEdges(node) {
				if _, ok := next[edge]; !ok {
					next[edge] = count
				} else {
					next[edge] += count
				}
			}
		}
	}

	stoneCount := 0
	for _, count := range next {
		stoneCount += count
	}
	return stoneCount
}