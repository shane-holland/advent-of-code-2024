/* -------------------------------------------------------------------------- */
/*                      --- Day 13: Claw Contraption ---                      */
/* -------------------------------------------------------------------------- */
package day13

import (
	"fmt"
	"regexp"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	return part1(input), part2(input)
}

// Part 1: Return the minimum cost to win the prize.
// Limit each button press to 100.
func part1(input string) string {
	clawMachines := parseClawMachines(input)
	cost := 0
	for _, clawMachine := range clawMachines {
		cost += clawMachine.MinimumCost(100)
	}
	
	return fmt.Sprintf("%d", cost)
}

// Part 2: Return the minimum cost to win the prize
// with the prize coordinates increased by 10000000000000.
// There is no limit to the number of button presses.
func part2(input string) string {
	clawMachines := parseClawMachines(input)
	cost := 0
	for _, clawMachine := range clawMachines {
		clawMachine.Prize.X += 10000000000000
		clawMachine.Prize.Y += 10000000000000
		cost += clawMachine.MinimumCost(-1)
	}
	
	return fmt.Sprintf("%d", cost)
}

/* ---------------------- Button Definition and Methods --------------------- */
// Button struct definition.
type Button struct {
	Action util.Point
	Cost int
}

/* ------------------- Claw Machine Definition and Methods ------------------ */
// ClawMachine struct definition.
type ClawMachine struct {
	ButtonA Button
	ButtonB Button
	Prize util.Point
}

// Return the minimum cost to win the prize, limiting each button press to the pressLimit.
func (cm ClawMachine) MinimumCost(pressLimit int) int {
	aPresses, bPresses := cm.PressesToWin()

	// Check to see if the press numbers are whole numbers
	// If they are not, then the prize cannot be won.
	if aPresses != float64(int(aPresses)) || bPresses != float64(int(bPresses)) {
		return 0
	}

	aButtonPresses := int(aPresses)
	bButtonPresses := int(bPresses)
	if (pressLimit == -1 || (aButtonPresses <= pressLimit && bButtonPresses <= pressLimit)) {
		return aButtonPresses * cm.ButtonA.Cost + bButtonPresses * cm.ButtonB.Cost
	}
	return 0
}

// Return the number of button presses required to win the prize.
func (cm ClawMachine) PressesToWin() (buttonA, buttonB float64) {

	// Given the puzzle input, we can determine the minimum number of button presses by 
	// converting the X values and Y values into linear equations and solving for the intersection.
	
	// Example:
	//
	// 	Button A: X+94, Y+34
	// 	Button B: X+22, Y+67
	//	Prize: X=8400, Y=5400
	//
	// Converted into linear equations:
	// 	8400 = 94X + 22Y
	//	5400 = 34X + 67Y
	//
	// Solving for the intersection, we get: X = 80, Y = 40
	//	Button A presses: 80
	//  Button B presses: 40

	// Solve for Button B first
	buttonB = float64((cm.Prize.Y * cm.ButtonA.Action.X) - (cm.ButtonA.Action.Y * cm.Prize.X))
	buttonB = buttonB / float64(cm.ButtonA.Action.X * cm.ButtonB.Action.Y - cm.ButtonA.Action.Y * cm.ButtonB.Action.X)

	// Use Button B to solve for Button A
	buttonA = (float64(cm.Prize.X) - (float64(cm.ButtonB.Action.X) * buttonB)) / float64(cm.ButtonA.Action.X)

	return buttonA, buttonB
}
/* ---------------------------- Helper Functions ---------------------------- */

// Parses all claw machines from the input.
func parseClawMachines(input string) []ClawMachine {
	clawMachines := []ClawMachine{}
	for _, lines := range strings.Split(input, "\n\n") {
		clawMachines = append(clawMachines, parseClawMachine(lines))
	}
	return clawMachines
}

// Parses a single claw machine from the input.
func parseClawMachine(input string) ClawMachine {
	// This pattern should match both the button and prize coordinates.
	pattern := `.*: X.(\d+), Y.(\d+)`

	re := regexp.MustCompile(pattern)
	lines := util.GetLines(input)

	coords := []util.Point{}
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		x := util.AtoI(match[1])
		y := util.AtoI(match[2])
		coords = append(coords, util.Point{X: x, Y: y})
	}

	return ClawMachine{
		ButtonA: Button{Action: coords[0], Cost: 3},
		ButtonB: Button{Action: coords[1], Cost: 1},
		Prize: coords[2],
	}
}

