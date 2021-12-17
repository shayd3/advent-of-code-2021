package days

import (
	"fmt"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

const NEW_BORN_LANTERNFISH_LIFE_TIMER = 8
const LANTERNFISH_LIFE_TIMER = 6

type Lanternfish struct {
	lifespan int
	lifespanRefresh bool
}

func Day06() {
	// inputs := inputs.Day06Sample
	inputs := inputs.Day06
	totalFishPart1 := simulateLanternfish(inputs, 80, 8, 6)
	totalFishPart2 := simulateLanternfish(inputs, 256, 8, 6)
	fmt.Printf("Total number of Lanternfishes after %d days: %d\n", 18, totalFishPart1)
	fmt.Printf("Total number of Lanternfishes after %d days: %d\n", 256, totalFishPart2)
}

// simulateLanternfish takes in a starting list of fishes and days to simulate
// returns total number of lanternfish after n number of given days
func simulateLanternfish(fishes []int, days int, newBornLifespan int, fishLifespanDayRefresh int) int {
	fishPopulationSimulation := make([]int, newBornLifespan + 1)
	// Initlize population simulation with fishes input
	for _, fish := range fishes {
		fishPopulationSimulation[fish]++
	}
	for day := 0; day < days; day++ {
		// Hold first value for all 0s
		tmp := fishPopulationSimulation[0]
		for i := 1; i < len(fishPopulationSimulation); i++ {
			fishPopulationSimulation[i-1] = fishPopulationSimulation[i]
		}
		fishPopulationSimulation[fishLifespanDayRefresh] += tmp
		fishPopulationSimulation[newBornLifespan] = tmp
		//fmt.Printf("Fish population Day#%d => \n\t%v\n", day+1, fishPopulationSimulation)
	}
	

	sum := 0
	for _, val := range fishPopulationSimulation {
		sum += val
	}
	return sum
}