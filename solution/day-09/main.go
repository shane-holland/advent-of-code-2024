package day09

import (
	"fmt"
	"sort"

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
	diskMap := parseDiskMap(input)
	diskMap.BlockDefrag(false)
	return fmt.Sprintf("%d", diskMap.Checksum())
}

/**
 * Function to find the minimum of two integers.
 */
func part2(input string) string {
	diskMap := parseDiskMap(input)
	diskMap.BlockDefrag(true)
	return fmt.Sprintf("%d", diskMap.Checksum())
}

/**
 * Struct to represent a file
 */
type File struct {
	// The file id
	Id int
	// The file size
	Size int
}

/**
 * Struct to represent a disk map
 */
type DiskMap struct {
	// The original diskmap string
	Map string
	// Map of file positions to file objects
	Files map[int]File
	// Map of free space positions to free space sizes
	FreeSpace map[int]int
}

/**
 * Function to parse the input into a DiskMap
 */
func parseDiskMap(input string) DiskMap {
	diskMap := DiskMap{
		Map:       input,
		Files:     make(map[int]File),
		FreeSpace: make(map[int]int),
	}

	position := 0

	for i, c := range input {
		size := util.AtoI(string(c))

		if (i % 2) == 0 {
			diskMap.Files[position] = File{Id: i / 2, Size: size}
		} else if size > 0 {
			diskMap.FreeSpace[position] = size
		}
		position += size
	}
	return diskMap
}

/**
 * Returns a sorted list of file positions
 */
func (diskMap *DiskMap) sortedFilePositions(reverse bool) []int {
	filePositions := make([]int, 0)
	for pos := range diskMap.Files {
		filePositions = append(filePositions, pos)
	}

	if reverse {
		sort.Sort(sort.Reverse(sort.IntSlice(filePositions)))
	} else {
		sort.Ints(filePositions)
	}

	return filePositions
}

/**
 * Returns a sorted list of free space positions
 */
func (diskMap *DiskMap) sortedFreeSpacePositions() []int {
	freeSpacePositions := make([]int, 0)
	for pos := range diskMap.FreeSpace {
		freeSpacePositions = append(freeSpacePositions, pos)
	}
	sort.Ints(freeSpacePositions)

	return freeSpacePositions
}

/**
 * Generates a checksum for the disk map
 */
func (diskMap *DiskMap) Checksum() int {
	checksum := 0

	for _, pos := range diskMap.sortedFilePositions(false) {
		fileId := diskMap.Files[pos].Id
		size := diskMap.Files[pos].Size
		for i := 0; i < size; i++ {
			checksum += fileId * (pos + i)
		}
	}

	return checksum
}

/**
 * Simple defrag which moves individual file blocks from the rightmost files to the first available free space
 */
func (diskMap *DiskMap) BlockDefrag(wholeFiles bool) {
	for _, pos := range diskMap.sortedFilePositions(true) {
		fileId := diskMap.Files[pos].Id
		fileSize := diskMap.Files[pos].Size
		for _, freePos := range diskMap.sortedFreeSpacePositions() {
			freeSize := diskMap.FreeSpace[freePos]

			if freePos > pos {
				if wholeFiles {
					break
				}
				return
			}

			// Free space can accommodate the file blocks
			if freeSize >= fileSize {
				diskMap.moveFileToFreeSpace(pos, freePos)
				break
			} else if !wholeFiles {
				// Free space cannot accommodate the file blocks, move the ones that can fit
				fileSize -= freeSize
				diskMap.Files[pos] = File{Id: fileId, Size: fileSize}

				diskMap.Files[freePos] = File{Id: fileId, Size: freeSize}
				delete(diskMap.FreeSpace, freePos)
			}
		}
	}
}

func (diskMap *DiskMap) moveFileToFreeSpace(filePos int, freePos int) {
	freeSize := diskMap.FreeSpace[freePos]
	fileSize := diskMap.Files[filePos].Size

	diskMap.FreeSpace[freePos+fileSize] = freeSize - fileSize
	delete(diskMap.FreeSpace, freePos)
	diskMap.Files[freePos] = File{Id: diskMap.Files[filePos].Id, Size: fileSize}
	delete(diskMap.Files, filePos)
	if freeSize-fileSize == 0 {
		delete(diskMap.FreeSpace, freePos+fileSize)
	}
}
