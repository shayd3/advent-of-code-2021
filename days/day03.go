package days

import (
	"fmt"
	"log"
	"strconv"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

type ReportResult struct {
	gamma int64
	epsilon int64
	co2Scrubber int64
	oxygenGenerator int64
}


func Day03() {
	inputs := inputs.Day03
	reportResults := generateReportResults(inputs)
	fmt.Printf("Report Product: %d", getReportProduct(reportResults.gamma, reportResults.epsilon))
	
}

func getReportProduct(gamma int64, epsilon int64) int64 {
	return gamma * epsilon
}

func generateReportResults(inputs []string) ReportResult {
	reportResults := ReportResult{}
	gammaRaw := ""
	epsilonRaw := ""

	for col := 0; col < len(inputs[0]); col++ {
		zeros := 0
		ones := 0
		for row := 0; row < len(inputs); row++ {
			bit := inputs[row][col]
			if (bit == '0') {
				zeros++
			} else {
				ones++
			}
		}
		most, least := getFrequency(ones, zeros)
		gammaRaw += strconv.Itoa(most)
		epsilonRaw += strconv.Itoa(least)
	}
	reportResults.gamma = convertToBinary(gammaRaw)
	reportResults.epsilon = convertToBinary(epsilonRaw)
	return reportResults
}
func convertToBinary(str string) int64 {
	output, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func getFrequency(ones int, zeros int) (most int, least int) {
	if(ones > zeros) {
		return 1, 0
	} else {
		return 0, 1
	}
}