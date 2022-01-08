package days

import (
	"fmt"
	"strings"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

const DELIMITER string = "|"

type Entry struct {
	signalPatterns []string
	outputValue []string
}

func Day08() {
	// input := inputs.Day08Sample
	input := inputs.Day08
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
		signalPatterns := strings.Fields(valSplitOnDelim[0])
		outputValue := strings.Fields(valSplitOnDelim[1])

		entries = append(entries, Entry{ signalPatterns: signalPatterns, outputValue: outputValue })
	}

	return entries
}

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
		digitMap := make(map[int]string)

		// Known easy digits
		for _, signalPattern := range entry.signalPatterns {
			switch len(signalPattern) {
				case 2: digitMap[1] = signalPattern;
				case 4: digitMap[4] = signalPattern; 
				case 3: digitMap[7] = signalPattern;
				case 7: digitMap[8] = signalPattern;
			}
		}
		// There is absolutely a better way of doing it, but brute force is the easiest way atm
		// Difficult digits - Find 0,6,9
		for _, signalPattern := range entry.signalPatterns {
			// 0,6,9
			if len(signalPattern) == 6 {
				// If signalPattern contains "4", it's 9
				if  strings.ContainsRune(signalPattern, rune(digitMap[4][0])) &&
				strings.ContainsRune(signalPattern, rune(digitMap[4][1])) &&
				strings.ContainsRune(signalPattern, rune(digitMap[4][2])) &&
				strings.ContainsRune(signalPattern, rune(digitMap[4][3])) {
					digitMap[9] = signalPattern
				// If signalPattern contains "1", it's 0
				} else if (strings.ContainsRune(signalPattern, rune(digitMap[1][0])) && 
				strings.ContainsRune(signalPattern, rune(digitMap[1][1]))) {
					digitMap[0] = signalPattern
				} else {
					digitMap[6] = signalPattern
				}
			}
		}

		// More difficult digits - find 2,3,5
		for _, signalPattern := range entry.signalPatterns {
			if len(signalPattern) == 5 {
				// if 5 is in 6
				if  strings.ContainsRune(digitMap[6], rune(signalPattern[0])) &&
				strings.ContainsRune(digitMap[6], rune(signalPattern[1])) &&
				strings.ContainsRune(digitMap[6], rune(signalPattern[2])) &&
				strings.ContainsRune(digitMap[6], rune(signalPattern[3])) &&
				strings.ContainsRune(digitMap[6], rune(signalPattern[4])) {
					digitMap[5] = signalPattern
				// if 1 is in 3
				} else if strings.ContainsRune(signalPattern, rune(digitMap[1][0])) && 
				strings.ContainsRune(signalPattern, rune(digitMap[1][1])) {
					digitMap[3] = signalPattern
				} else {
					digitMap[2] = signalPattern
				}
			}
		}
		// strVal := ""
		// for _, outputVal := range entry.outputValue {
			
		// }
		// intOutput, _ := strconv.Atoi(strOutput)
		// sum += intOutput
		fmt.Println(digitMap)
	}
	return sum
} 