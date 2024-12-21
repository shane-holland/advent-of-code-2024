// The util package contains utility functions that are used throughout the application.
package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

/**
 * ReadFile reads the contents of a file and returns it as a string.
 */
func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

/**
 * GetLines splits a string into an array of strings by newline characters.
 */
func GetLines(input string) []string {
	return strings.Split(input, "\n")
}

/**
 * GetColumns splits a string into an array of strings by newline characters, then returns an array of strings
 * where each string is a column of the input.
 */
func GetColumns(input string) []string {
	lines := GetLines(input)
	columns := make([]string, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		column := ""
		for _, line := range lines {
			column += string(line[i])
		}
		columns[i] = column
	}
	return columns
}

/**
 * Function to return the Integer value of a string, or exit the program if the string is not an integer.
 */
func AtoI(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Failed to convert %s to an integer.", val)
	}
	return num
}

/**
 * Function to find the absolute value of an integer
 */
func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Point is a struct to represent a point in 2D space.
type Point struct {
	X int
	Y int
}

func ManhattanDistance(p1, p2 Point) int {
	return AbsInt(p1.X - p2.X) + AbsInt(p1.Y - p2.Y)
}