package days

import (
	"strconv"
	"strings"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

type Line struct {
	start Coordinate
	end Coordinate
}

type Coordinate struct {
	x int
	y int
}

// Day05
// Part 1:
// 1. Convert inputs to usable format
// 2. Build "grid" based on highest point in coordinates
// 3. Increment points on the grid based on coordinates start -> end points
// 4. Count all points on grid that are >= 2
// 5. Return count
func Day05() {
	coordinatesInput := inputs.Day05Sample
	lines := generateLines(coordinatesInput)
	grid := drawGrid(lines)

}

// generateLines returns a slice of Lines that are 
// parsed from a slice of strings where an element in that
// slice would look like "0,9 -> 5,9". White space is trimmed,
// '->' is removed, and a `Line` object is created
func generateLines(coordinatesInput []string) []Line {
	lines := []Line{}
	for _, input := range coordinatesInput {
		input = (strings.ReplaceAll(input, " ", ""))
		coordinatePair := strings.Split(input, "->")
		startCoordinateSplit := strings.Split(coordinatePair[0], ",")
		endCoordinateSplit := strings.Split(coordinatePair[1], ",")
		startX, _ := strconv.Atoi(startCoordinateSplit[0])
		startY, _ := strconv.Atoi(startCoordinateSplit[1])
		endX, _ := strconv.Atoi(endCoordinateSplit[0])
		endY, _ := strconv.Atoi(endCoordinateSplit[1])

		startCoordinate := Coordinate {
			x: startX,
			y: startY,
		}
		endCoordinate := Coordinate {
			x: endX,
			y: endY,
		}
		lines = append(lines, Line{
			start: startCoordinate,
			end: endCoordinate,
		})
	}

	return lines
}