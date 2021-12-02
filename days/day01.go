package days

import (
	"fmt"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

// Day01 - https://adventofcode.com/2021/day/1
func Day01() {
	inputs := inputs.Day01
	count := 0
	for i := 1; i < len(inputs); i++ {
		if (inputs[i-1] < inputs[i]) {
			count++
		}
	}
	fmt.Println(count)
}