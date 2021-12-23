package days

import (
	"fmt"
	"strings"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

const DELIMITER string = "|"

type Entry struct {
	signalPaterns []string
	outputValue []string
}

func Day08() {
	// input := inputs.Day08Sample
	input := inputs.Day08
	entries := parseInput(input)
	count := findEasyDigitsCount(entries)
	fmt.Printf("Total times 1, 4, 7, or 8 appeared: %d\n", count)
}

func parseInput(input []string) []Entry {
	entries := []Entry{}
	for _, val := range input {
		valSplitOnDelim := strings.Split(val, DELIMITER)
		signalPaterns := strings.Fields(valSplitOnDelim[0])
		outputValue := strings.Fields(valSplitOnDelim[1])

		entries = append(entries, Entry{ signalPaterns: signalPaterns, outputValue: outputValue })
	}

	return entries
}
/*
break down of lengths:
len = number

2 => 1
3 => 7
4 => 4
7 => 8

5 => 2,3,5
	- Contains 7 -> 3
	- Else [2,5]
		- intersects "4" in 3 places -> 5
		- else -> 2
6 => 0,6,9
	- Contains 4 -> 9
	- Else -> [0,6]
*/

// findEasyDigitsCount finds total count of how many times
// 1, 4, 7, and 8 in a slice of Entry objects
func findEasyDigitsCount(entries []Entry) int {
	count := 0
	for _, entry := range entries {
		for _, outputVal := range entry.outputValue {
			switch len(outputVal) {
				case 2: count++; // 1
				case 3: count++; // 7
				case 4: count++; // 4
				case 7: count++; // 8
			}
		}
	}
	return count
} 