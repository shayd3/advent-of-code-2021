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
	fmt.Printf("Report Product: %d", getPowerConsumption(reportResults))
	
}

func generateReportResults(inputs []string) ReportResult {
	reportResults := ReportResult{}
	gamma, epsilon := getPowerConsumptionProperties(inputs)

	reportResults.gamma = gamma
	reportResults.epsilon = epsilon
	return reportResults
}

func getPowerConsumptionProperties(inputs []string) (gamma int64, epsilon int64) {
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

	return convertToBinary(gammaRaw), convertToBinary(epsilonRaw)
}

func getFrequency(ones int, zeros int) (most int, least int) {
	if(ones > zeros) {
		return 1, 0
	} else {
		return 0, 1
	}
}

func convertToBinary(str string) int64 {
	output, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func getPowerConsumption(reportResult ReportResult) int64 {
	return reportResult.gamma * reportResult.epsilon
}