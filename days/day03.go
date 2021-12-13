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
	fmt.Printf("Report Product: %d\n", getPowerConsumption(reportResults))
	fmt.Printf("Life Support Rating: %d\n", getLifeSupportRating(reportResults))
}

func generateReportResults(inputs []string) ReportResult {
	reportResults := ReportResult{}
	gamma, epsilon := getPowerConsumptionProperties(inputs)
	co2Scrubber := getCO2ScrubberRating(inputs)
	oxygenGenerator := getOxygenGeneratorRating(inputs)
	
	reportResults.gamma = gamma
	reportResults.epsilon = epsilon
	reportResults.oxygenGenerator = oxygenGenerator
	reportResults.co2Scrubber = co2Scrubber
	return reportResults
}

func getPowerConsumptionProperties(inputs []string) (gamma int64, epsilon int64) {
	gammaRaw := ""
	epsilonRaw := ""

	// TODO: Extract the column loops into separate function
	// TODO: Figure out different way to extract gamma and epsilon values
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
		most, least := getFrequency(ones, zeros, 1)
		gammaRaw += strconv.Itoa(most)
		epsilonRaw += strconv.Itoa(least)
	}

	return convertToBinary(gammaRaw), convertToBinary(epsilonRaw)
}

func getOxygenGeneratorRating(inputs []string) (oxygenGeneratorRating int64) {
	foundRating := []string{}
	// TODO: Extract the column loops into separate function
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
		most, _ := getFrequency(ones, zeros, 1)
		byteMost := strconv.Itoa(most)
		tmp := []string{}
		for index, val := range inputs {
			if inputs[index][col] == byteMost[0] {
				tmp = append(tmp, val)
			}
		}
		inputs = tmp
		if len(inputs) == 1 {
			foundRating = inputs
			break
		}
	}

	return convertToBinary(foundRating[0])
}

func getCO2ScrubberRating(inputs []string) (co2ScrubberRating int64) {
	foundRating := []string{}
	// TODO: Extract the column loops into separate function
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
		_, least := getFrequency(ones, zeros, 0)
		byteLeast := strconv.Itoa(least)
		tmp := []string{}
		for index, val := range inputs {
			if inputs[index][col] == byteLeast[0] {
				tmp = append(tmp, val)
			}
		}
		inputs = tmp
		if len(inputs) == 1 {
			foundRating = inputs
			break
		}
	}

	return convertToBinary(foundRating[0])
}

func getFrequency(ones int, zeros int, tieBreaker int) (most int, least int) {
	if(ones == zeros) {
		return tieBreaker, tieBreaker
	} else if(ones > zeros) {
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

func getLifeSupportRating (reportResult ReportResult ) int64 {
	return reportResult.co2Scrubber * reportResult.oxygenGenerator
}