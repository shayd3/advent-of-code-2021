package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

func Day09() {
	input := inputs.Day09Sample
	heatmap := generateHeatmap(input)
	fmt.Print(heatmap)
}

func generateHeatmap(input string) [][]int {
	heatmap := [][]int{}
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		heatmapLine := []int{}
		for _, char := range line {
			i, _ := strconv.Atoi(string(char))
			heatmapLine = append(heatmapLine, i)
		}
		heatmap = append(heatmap, heatmapLine)
	}
	return heatmap
}