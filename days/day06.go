package days

import (
	"fmt"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

const NEW_BORN_LANTERNFISH_LIFE_TIMER = 8
const LANTERNFISH_LIFE_TIMER = 6
const TOTAL_DAYS = 80

type Lanternfish struct {
	lifespan int
	lifespanRefresh bool
}

func Day06() {
	//inputs := inputs.Day06Sample
	inputs := inputs.Day06
	lanternfishes := parseLanternfish(inputs)
	totalFish := simulateLanternfish(lanternfishes, TOTAL_DAYS)
	fmt.Printf("Total number of Lanternfishes after %d days: %d\n", TOTAL_DAYS, totalFish)
}

func parseLanternfish(input []int) []Lanternfish {
	lanternfishes := []Lanternfish{}
	for _, fishLifespan := range input {
		lanternfishes = append(lanternfishes, Lanternfish{lifespan: fishLifespan})
	}
	return lanternfishes
}

// simulateLanternfish takes in a starting list of fishes and days to simulate
// returns total number of lanternfish after n number of given days
func simulateLanternfish(fishes []Lanternfish, days int) int {
	for i := 0; i < days; i++ {
		fmt.Printf("Day #%d => Current Total: %d\n", i+1, len(fishes))
		newFishToSpawn := 0
		for i := range fishes {
			if fishes[i].lifespan == 0 {
				newFishToSpawn++
				fishes[i].lifespan = LANTERNFISH_LIFE_TIMER
				fishes[i].lifespanRefresh = true
			}
			if fishes[i].lifespanRefresh == false {
				fishes[i].lifespan--
			}
			fishes[i].lifespanRefresh = false
		}
		for i := 0; i < newFishToSpawn; i++ {
			fishes = append(fishes, Lanternfish{ lifespan: NEW_BORN_LANTERNFISH_LIFE_TIMER})
		}
	}
	return len(fishes)
}