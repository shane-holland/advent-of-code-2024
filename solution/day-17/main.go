/* -------------------------------------------------------------------------- */
/*                   --- Day 17: Chronospatial Computer ---                   */
/* -------------------------------------------------------------------------- */
package day17

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"slices"
	"strconv"
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

// Part 1: Retrieve the output after running the program on the 3-bit computer
func part1(input string) string {
	computer := NewComputer(input)
	for computer.RunNextInstruction() {
		// Continue until execution halts
	}
	return fmt.Sprintf("%v", computer.GetOuput())
}

// Part 2: Determine the correct 'A' Register value which generates the Program as the output
func part2(input string) string {
	computer := NewComputer(input)
	return fmt.Sprintf("%d", computer.GetSelfProducingRegister())
	
}

/* --------------------- Computer Definition and Methods -------------------- */

// Define Computer, which consists of 3 register values (A,B,C), A program which is a list of 3 bit integers, 
//   an instruction pointer, and it's output
type Computer struct {
	Registers map[rune]int
	Program   []int

	pointer int
	Output  []int
}

// Returns the values in the output buffer as a string separated by commas
func (c Computer) GetOuput() string {
	output := ""
	for _, n := range c.Output {
		output += strconv.Itoa(n) + ","
	}
	return output[:len(output)-1]
}

// Determine the 'A' Register value which causes the program to output itself
func (c *Computer) GetSelfProducingRegister() int {
	components := len(c.Program) - 1
	// Set the register value to 8 to the power of the number of digits in the program minus 1
	//  This will allow us to generate an output of appropriate length 
	register := math.Pow(8, float64(components))

	matched := 0
	for {
		// Reset the computer
		c.Registers['A'] = int(register)  // Set the 'A' register to the test value
		c.Registers['B'] = 0 
		c.Registers['C'] = 0
		c.Output = make([]int, 0)
		c.pointer = 0

		// Run the program until completion
		for c.RunNextInstruction() {
			// Continue until execution halts
		}

		if slices.Equal(c.Output, c.Program) {
			return int(register)
		} else if register >= math.Pow(8, float64(components+1)) {
			// Failed to find the correct register value
			return -1
		}

		if !slices.Equal(c.Output[components-matched:], c.Program[components-matched:]) {
			// The last set of numbers do not match, increment the test value by the appropriate factor
			register += math.Pow(8, float64(components-matched-1))
		} else {
			// Match found, increment the number of matches to reduce the increment factor
			matched++
		}
	}
}

// Returns the value of a 'combo operand' for a given literal operand
func (c Computer) GetComboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.Registers['A']
	case 5:
		return c.Registers['B']
	case 6:
		return c.Registers['C']
	}
	log.Fatalf("Invalid combo operand encountered: %d", operand)
	return -1
}

// Completes the next instruction at the instruction pointer, using the operand at the next position
func (c *Computer) RunNextInstruction() bool {
	if c.pointer >= len(c.Program) || c.Program[c.pointer] > 7 {
		// Halt Execution
		return false
	}

	// Literal Operand
	literal := c.Program[c.pointer+1]
	// Combo Operand
	combo := c.GetComboOperand(literal)

	switch c.Program[c.pointer] {
	// opcode 0 - adv
	case 0:
		c.Registers['A'] = int(math.Floor(float64(c.Registers['A']) / math.Pow(2, float64(combo))))
	// opcode 1 - bxl
	case 1:
		c.Registers['B'] = c.Registers['B'] ^ literal
	// opcode 2 - bst
	case 2:
		c.Registers['B'] = combo % 8
	// opcode 3 - jnz
	case 3:
		if c.Registers['A'] != 0 {
			c.pointer = literal
			c.pointer -= 2
		}
	// opcode 4 - bxc
	case 4:
		c.Registers['B'] = c.Registers['B'] ^ c.Registers['C']
	// opcode 5 - out
	case 5:
		c.Output = append(c.Output, combo % 8)
	// opcode 6 - bdv
	case 6:
		c.Registers['B'] = int(math.Floor(float64(c.Registers['A']) / math.Pow(2, float64(combo))))
	// opcode 7 - cdv
	case 7:
		c.Registers['C'] = int(math.Floor(float64(c.Registers['A']) / math.Pow(2, float64(combo))))
	default:
		return false
	}

	c.pointer += 2

	return true
}

/* ----------------------------- Helper Methods ----------------------------- */

// Generate a Computer from an input string
func NewComputer(input string) Computer {
	computer := Computer{Registers: make(map[rune]int), Program: make([]int, 0)}

	for _, l := range util.GetLines(strings.Split(input, "\n\n")[0]) {
		r, v := parseRegister(l)
		computer.Registers[r] = v
	}
	computer.Program = parseProgram(strings.Split(input, "\n\n")[1])
	return computer
}

// Parse a register value from the input string
func parseRegister(line string) (rune, int) {
	match := regexp.MustCompile(`Register (\w): (\d+)`).FindStringSubmatch(line)

	if len(match) == 0 {
		log.Fatalf("Unable to parse Register: %s", line)
	}
	return []rune(match[1])[0], util.AtoI(match[2])
}

// Retrieve the program as a list of 3 bit integers
func parseProgram(line string) []int {
	values := strings.Split(strings.Replace(line, "Program: ", "", 1), ",")
	numbers := make([]int, len(values))
	for i, v := range values {
		numbers[i] = util.AtoI(v)
	}
	return numbers
}
