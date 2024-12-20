/* -------------------------------------------------------------------------- */
/*                         --- Day 5: Print Queue ---                         */
/* -------------------------------------------------------------------------- */
package day05

import (
	"fmt"
	"slices"
	"strings"

	"shaneholland.dev/aoc-2024/util"
)

type Puzzle struct{}

func (d Puzzle) Solve(input string) (string, string) {
	printQueue := parsePrintQueue(input)
	return part1(printQueue), part2(printQueue)
}

// Part 1: Find the sum of the middle pages of all valid print jobs.
func part1(printQueue PrintQueue) string {
	middlePageSum := 0
	for _, printJob := range printQueue.GetValidPrintJobs() {
		middlePageSum += printJob.Pages[(len(printJob.Pages)-1)/2]
	}

	return fmt.Sprintf("%d", middlePageSum)
}

// Part 2: Find the sum of the middle pages of all corrected print jobs.
func part2(printQueue PrintQueue) string {
	middlePageSum := 0
	for _, printJob := range printQueue.GetCorrectedPrintJobs() {
		middlePageSum += printJob.Pages[(len(printJob.Pages)-1)/2]
	}

	return fmt.Sprintf("%d", middlePageSum)
}

/* --------------------------- PrintJob Definition -------------------------- */

// PrintJob represents a print job with a list of pages.
type PrintJob struct {
	Pages []int
}

/* -------------------- PrintQueue Definition and Methods ------------------- */

// PrintQueue represents a print queue with a list of print jobs and page rules.
type PrintQueue struct {
	Jobs      []PrintJob
	PageRules map[int][]int
}

// GetValidPrintJobs returns a list of all valid print jobs in the print queue.
func (pq *PrintQueue) GetValidPrintJobs() []PrintJob {
	valid := make([]PrintJob, 0)
	for _, job := range pq.Jobs {
		if pq.isPrintJobValid(job) {
			valid = append(valid, job)
		}
	}
	return valid
}

// GetCorrectedPrintJobs returns a list of all corrected print jobs in the print queue.
func (pq *PrintQueue) GetCorrectedPrintJobs() []PrintJob {
	corrected := make([]PrintJob, 0)
	for _, job := range pq.Jobs {
		if !pq.isPrintJobValid(job) {
			corrected = append(corrected, pq.correctPrintJob(job))
		}
	}
	return corrected
}

// isPrintJobValid returns true if the print job is valid.
func (pq *PrintQueue) isPrintJobValid(printJob PrintJob) bool {
	for i := 0; i < len(printJob.Pages); i++ {
		for j := 0; j < i; j++ {
			if slices.Contains(pq.PageRules[printJob.Pages[i]], printJob.Pages[j]) {
				return false
			}
		}
	}
	return true
}

// correctPrintJob returns the corrected print job.
func (pq *PrintQueue) correctPrintJob(printJob PrintJob) PrintJob {
	corrected := PrintJob{Pages: slices.Clone(printJob.Pages)}
	for i := 0; i < len(corrected.Pages); i++ {
		for j := 0; j < i; j++ {
			if slices.Contains(pq.PageRules[corrected.Pages[i]], corrected.Pages[j]) {
				// move the element at corrected[j] to be after corrected[i]
				modified := slices.Clone(corrected.Pages[:j])
				modified = append(modified, corrected.Pages[j+1:i]...)
				modified = append(modified, corrected.Pages[i], corrected.Pages[j])
				modified = append(modified, corrected.Pages[i+1:]...)
				corrected = PrintJob{modified}
				i--
				break
			}
		}
		if pq.isPrintJobValid(corrected) {
			return corrected
		}
	}

	return corrected
}

/* ----------------------------- Helper Methods ----------------------------- */

// parsePrintQueue returns a PrintQueue from the input string.
func parsePrintQueue(input string) PrintQueue {
	lines := util.GetLines(input)
	pageBreak := slices.Index(lines, "")
	pageRules := make(map[int][]int)

	// Parse the page rules
	for _, line := range lines[:pageBreak] {
		first := util.AtoI(strings.Split(line, "|")[0])
		second := util.AtoI(strings.Split(line, "|")[1])

		if !slices.Contains(pageRules[first], second) {
			pageRules[first] = append(pageRules[first], second)
		}
	}

	printJobs := make([]PrintJob, 0)
	// Parse the print jobs
	for _, line := range lines[pageBreak+1:] {
		printJob := make([]int, 0)
		for _, page := range strings.Split(line, ",") {
			pageInt := util.AtoI(page)
			printJob = append(printJob, pageInt)
		}

		printJobs = append(printJobs, PrintJob{Pages: printJob})
	}

	return PrintQueue{Jobs: printJobs, PageRules: pageRules}
}
