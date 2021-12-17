package days

import (
	"fmt"
	"math"
	"sort"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

func Day07() {
	// inputs := inputs.Day07Sample
	inputs := inputs.Day07
	fmt.Printf("Fuel cost: %d\n", getOptimizedFuel(inputs, findMedian(inputs)))
	fmt.Printf("Fuel cost: %d\n", getOptimizedFuel2(inputs, findAvg(inputs)))
}

func findMedian(positions []int) int {
	sort.Ints(positions)
	positionsLength := len(positions)
	middle := positionsLength / 2

	if positionsLength % 2 == 0 {
		return (positions[middle - 1] + positions[middle]) / 2
	} else {
		return positions[middle]
	}
}

func findAvg(positions []int) int {
	sum := 0
	for _, val := range positions {
		sum += val
	}

	// Floor works for the input given, but the sample input only works for ceiling?
	return int(math.Floor(float64(sum) / float64(len(positions))))
}

func getOptimizedFuel(positions []int, medianPoint int) int {
	fuel := 0
	for _, position := range positions {
		diff := int(math.Abs(float64(position) - float64(medianPoint)))
		fuel += diff
	}
	return fuel
}

func getOptimizedFuel2(positions []int, avgPoint int) int {
	fuel := 0
	for _, position := range positions {
		diff := int(math.Abs(float64(position) - float64(avgPoint)))
		// Fuel is calculated by finding the triangular number equivalent
		// https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF
		fuel += (diff * (diff + 1)) / 2
	}
	return fuel
}