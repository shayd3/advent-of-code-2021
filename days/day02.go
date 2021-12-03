package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

// Potentially make use of enum to assign directions to numbers?
type CourseStep struct {
	direction string
	count int
}

// Day02 - https://adventofcode.com/2021/day/2
func Day02() {
	inputs := inputs.Day02
	courseSteps := parseCourseSteps(inputs)
	position := getFinalPosition(courseSteps)
	fmt.Println(position)

	positionWithAim := getFinalPositionWithAim(courseSteps)
	fmt.Println(positionWithAim)
}

// parseCourseSteps takes slice of strings and 
// returns a slice of CourseSteps
func parseCourseSteps(inputs []string) []CourseStep {
	courseSteps := []CourseStep{}
	if(len(inputs) != 0) {
		for _, input := range inputs {
			courseStrSlice := strings.Fields(input)

			direction := courseStrSlice[0]
			count, err := strconv.Atoi(courseStrSlice[1])
			if err != nil {
				log.Fatal(err, courseStrSlice[1])
			}
			courseSteps = append(courseSteps, CourseStep{direction: direction, count: count})
		}
	}
	return courseSteps
}

func getFinalPosition(courseSteps []CourseStep) int {
	horizontalPos := 0
	depth := 0

	for _, step := range courseSteps {
		switch step.direction {
		case "forward":
			horizontalPos += step.count
		case "down":
			depth += step.count
		case "up":
			depth -= step.count
		}
	}
	return horizontalPos * depth
}

func getFinalPositionWithAim(courseSteps []CourseStep) int {
	horizontalPos := 0
	depth := 0
	aim := 0
	for _, step := range courseSteps {
		switch step.direction {
		case "forward":
			horizontalPos += step.count
			depth += step.count * aim
		case "down":
			aim += step.count
		case "up":
			aim -= step.count
		}
	}
	return horizontalPos * depth
}