package days

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	log.SetOutput(ioutil.Discard)
	// coordinatesInput := inputs.Day05Sample
	coordinatesInput := inputs.Day05
	lines := generateLines(coordinatesInput)
	gridPt1 := drawGrid(lines)
	gridPt2 := drawGrid(lines)
	gridPt1 = drawLinesOnGrid(gridPt1, lines, false)
	gridPt2 = drawLinesOnGrid(gridPt2, lines, true)
	
	fmt.Printf("Total number of overlapping points (Horizonal/Vertical): %d\n", countOverlappingPoints(gridPt1))
	fmt.Printf("Total number of overlapping points (Horizonal/Vertical/Diag): %d\n", countOverlappingPoints(gridPt2))
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func countOverlappingPoints(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val >= 2 {
				count++
			}
		}
	}
	return count
}

// drawLinesOnGrid will draw lines on the grid from the 
// list of lines
// TODO: The calculation of where to increment could be done better. 
// Especially with the diagonal method
func drawLinesOnGrid(grid [][]int, lines []Line, drawDiagonalLines bool) [][]int {
	for _, line := range lines {
		// Horizontal Line
		if line.start.y == line.end.y {
			init, ceiling := getLoopParameters(line.start.x, line.end.x)
			for i := init; i <= ceiling; i++ {
				grid[line.start.y][i]++
			}
			
		}
		// Vertical Line
		if line.start.x == line.end.x {
			init, ceiling := getLoopParameters(line.start.y, line.end.y)
			for i := init; i <= ceiling; i++ {
				grid[i][line.start.x]++
			}
		}

		// Diagonal Line
		if drawDiagonalLines {
			isDiagonal, _, _ := getLineDetails(line)
			if isDiagonal {
				points := getDiagonalPoints(line)
				for _,point := range points {
					grid[point.y][point.x]++
				}
			}
		}
		
	}

	return grid
}

// if x1 < x2 -> Right
// if x1 > x2 -> Left
// if y1 < y2 -> Down
// if y1 > y2 -> Up
func getDiagonalPoints(line Line) []Coordinate {
	coordinates := []Coordinate{}
	_, _, yRange := getLineDetails(line)
	for i := 0; i <= yRange; i++ {
		log.Printf("Point #%d\n", i)
		coordinate := Coordinate{}
		// Right
		if line.start.x < line.end.x {
			coordinate.x = line.start.x + i
		}
		// Left
		if line.start.x > line.end.x {
			coordinate.x = line.start.x - i
		}
		// Down
		if line.start.y < line.end.y {
			coordinate.y = line.start.y + i
		}
		// Up
		if line.start.y > line.end.y {
			coordinate.y = line.start.y - i
		}
		coordinates = append(coordinates, coordinate)

	}
	return coordinates
}


func getLineDetails(line Line) (isDiagonal bool, xRange int, yRange int) {
	xRange = int(math.Abs(float64(line.start.x - line.end.x)))
	yRange = int(math.Abs(float64(line.start.y - line.end.y)))
	isDiagonal = xRange == yRange

	return isDiagonal, xRange, yRange
}

func getLoopParameters(boundryA int, boundryB int) (init int, ceiling int) {
	if boundryA < boundryB {
		init = boundryA
		ceiling = boundryB
	} else {
		init = boundryB
		ceiling = boundryA
	}
	
	return init, ceiling
}

func drawGrid(lines []Line) (grid [][]int) {
	maxX := 0
	maxY := 0

	for _, line := range lines {
		// Check x
		if line.start.x > maxX {
			maxX = line.start.x
		}
		if line.end.x > maxX {
			maxX = line.end.x
		}

		// Check y
		if line.start.y > maxY {
			maxY = line.start.y
		}
		if line.end.y > maxY {
			maxY = line.end.y
		}
	}
	// Add 1 for both the x and y axis for the slice size
	grid = make([][]int, maxY + 1)
	for i := range grid {
		grid[i] = make([]int, maxX + 1)
	}
	return grid
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