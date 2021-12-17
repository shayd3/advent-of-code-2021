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
	fmt.Printf("Fuel cost: %d", getOptimizedFuel(inputs, findMedian(inputs)))
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

func getOptimizedFuel(positions []int, medianPoint int) int {
	fuel := 0
	for _, position := range positions {
		diff := int(math.Abs(float64(position) - float64(medianPoint)))
		fuel += diff
	}
	return fuel
}