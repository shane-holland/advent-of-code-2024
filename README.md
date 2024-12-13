# Advent of Code in Go

Welcome to my Advent of Code repository! 🎄 This repository contains my solutions for [Advent of Code](https://adventofcode.com/), implemented in the Go programming language.

## Table of Contents
- [About Advent of Code](#about-advent-of-code)
- [Why Go?](#why-go)
- [Repository Structure](#repository-structure)
- [How to Run](#how-to-run)

## About Advent of Code
Advent of Code is an annual coding challenge that runs every December. It features 25 days of programming puzzles, with a new challenge unlocked daily. It's a great way to practice problem-solving, learn new programming skills, and join a vibrant community of developers.

## Why Go?
Go (or Golang) is a simple, efficient, and powerful programming language. It is particularly well-suited for:
- Writing clear and maintainable code.
- Fast execution and compilation.
- Strong support for concurrency.

I've been meaning to learn Go for a few years now, and had never made the time.  I thought using Go for this year's advent of code would be a good way for me to get familiar with the fundamentals of the language.

## Repository Structure
```
├── README.md              # This file
├── data/
|   ├── day-01.txt         # Day 1 Puzzle Input (Not committed)
|   └── day-02.txt         # Day 2 Puzzle Input (Not committed)
├── solution/              
|   ├── day-01/            # Solutions for Day 1
|   |   ├── main.go
|   |   ├── main_test.go
|   |   └── test-data.txt  # Day 1 Test Data
|   ├── day-02/            # Solutions for Day 2
|   |   ├── main.go
|   |   ├── main_test.go 
|   |   └── test-data.txt  # Day 2 Test Data
|   ├── template/          # A template folder which can be copied 
|   |   ├── main.go        #   for a new days puzzle
|   |   ├── main_test.go
|   |   └── test-data.txt
|   ├── solution.go        # Solution Interface
|   └── solution-map.go    # Map of solutions
├── util/                  # Utility functions used across days
│   └── util.go
├── main.go                # Application entry point.
└── go.mod                 # Go module file
```

Each day has a solution directory containing:
- `main.go`: The solution for the day's puzzle.
- `main_test.go`: A unit test which tests the solution against the test data.
- `test-data.txt`: The test input for the day's puzzle.

Additionally, each day's real input should be stored in the `data` directory using the format `day-{nn}.txt` where `{nn}` is the current day represented as a two digit number with leading zero where applicable.

## How to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/shane-holland/advent-of-code-2024
   cd advent-of-code-2024
   ```

2. Install Go if you haven't already: [Go Installation Guide](https://golang.org/doc/install)

3. Run the solution:
   ```bash
   go run main.go -day n 
   ```

---

Happy coding and may your Advent of Code journey be joyful and enlightening! 🎅
