package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

const DELIMITER string = "|"

type Entry struct {
	signalPaterns []string
	outputValue []string
}

func Day08() {
	input := inputs.Day08Sample
	// input := inputs.Day08
	entries := parseInput(input)
	count := findEasyDigitsCount(entries)
	fmt.Printf("Total times 1, 4, 7, or 8 appeared: %d\n", count)
	totalSum := totalOutputValues(entries)
	fmt.Printf("Total sum of all output values: %d\n", totalSum)
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
		- if contains 7 -> 0
		- else -> 6
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

// totalOutputValues - Calculates digit for each slice of 
// Entry.outputValue and returns sum
func totalOutputValues(entries []Entry) int {
	sum := 0
	for _, entry := range entries {
		strOutput := ""
		for _, outputVal := range entry.outputValue {
			switch len(outputVal) {
			case 2: strOutput += "1"; // 1
			case 3: strOutput += "7"; // 7
			case 4: strOutput += "4"; // 4
			case 7: strOutput += "8"; // 8
			case 5:  { // 2,3,5
				if strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'e') && strings.ContainsRune(outputVal, 'f') && strings.ContainsRune(outputVal, 'c') && strings.ContainsRune(outputVal, 'b') || 
				   strings.ContainsRune(outputVal, 'a') && strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'c') && strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'e') || 
				   strings.ContainsRune(outputVal, 'a') && strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'e') && strings.ContainsRune(outputVal, 'f') && strings.ContainsRune(outputVal, 'g') {
					strOutput += "3"
				} else if strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'a') && strings.ContainsRune(outputVal, 'f') && strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'c') ||
						  strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'c') && strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'e') && strings.ContainsRune(outputVal, 'f') {
					strOutput += "5"
				} else { // 2
					strOutput += "2"
				}
			}
			case 6: {// 0,6,9
				// Check if 6 segments contains 4
				if strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'c') && strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'e') && strings.ContainsRune(outputVal, 'f') && strings.ContainsRune(outputVal, 'g') || 
				   strings.ContainsRune(outputVal, 'a') && strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'c') && strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'f') && strings.ContainsRune(outputVal, 'g') ||
				   strings.ContainsRune(outputVal, 'a') && strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'c') && strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'e') && strings.ContainsRune(outputVal, 'f') {
					strOutput += "9"
				// If 6 segments contains 7, it's 0
				} else if strings.ContainsRune(outputVal, 'd') && strings.ContainsRune(outputVal, 'e') && strings.ContainsRune(outputVal, 'a') && strings.ContainsRune(outputVal, 'f') && strings.ContainsRune(outputVal, 'b') && strings.ContainsRune(outputVal, 'g') && strings.ContainsRune(outputVal, 'c') {
					strOutput += "0"
				} else {
					strOutput += "6"
				}
			}
			}
		}
		intOutput, _ := strconv.Atoi(strOutput)
		sum += intOutput
	}
	return sum
} 