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
	input := inputs.Day08Sample
	entries := parseInput(input)
	for _, entry := range entries {
		fmt.Println(entry)
	}
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