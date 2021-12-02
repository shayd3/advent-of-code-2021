package days

import (
	"fmt"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

// Day01 - https://adventofcode.com/2021/day/1
func Day01() {
	inputs := inputs.Day01
	fmt.Println(getNumberOfIncreases(inputs, 1))
	fmt.Println(getNumberOfIncreases(inputs, 3))
}

// getNumberOfIncreases calculates how many times 
// there are positive increases on the previous number
// in the given input. 'window' is how many of the inputs
// are included in the comparison sums
func getNumberOfIncreases(inputs []int, window int) int {
	count := 0
	// Check if window is out of bounds or if no window at all
	if (len(inputs) < window || window <= 0 ) {
		return count
	}

	for i := window; i < len(inputs); i++ {
		// Only compare sums of windows if window is more that 1
		if(window > 1) {
			windowASum, windowBSum := 0, 0
			for _, valueA := range inputs[i-window : i] {
				windowASum += valueA
			}
			for _, valueB := range inputs[i-window+1 : i+1] {
				windowBSum += valueB
			}
			if(windowASum < windowBSum) {
				count++
			}

		} else {
			if (inputs[i-1] < inputs[i]) {
				count++
			}
		}
		
	}
	return count
}