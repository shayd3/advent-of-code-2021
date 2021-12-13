package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shayd3/advent-of-code-2021/inputs"
)

type BingoBoard struct {
	board [5][5]BingoNumber
}

type BingoNumber struct {
	value int
	marked bool
}

func Day04() {
	//bingoCalls := inputs.BingoCalls
	// bingoBoards := inputs.BingoBoards
	bingoBoards := inputs.BingoBoardsSample
	games := makeBoards(bingoBoards)
	
}

func makeBoards(boardInput string) (bingoGames []BingoBoard) {
	boardStrings := strings.Split(string(boardInput), "\n\n")
	// Get each board block from string
	for gameNumber, boardString := range boardStrings {
		fmt.Printf("Game #%d:\n%s\n", gameNumber, boardString)
		boardLineStrings := strings.Split(boardString, "\n")
		game := BingoBoard{}
		// Loop each line in board
		for lineNum, boardLine := range boardLineStrings {
			fmt.Printf("Line %d => %s\n", lineNum, boardLine)
			line := [5]BingoNumber{}
			// Get each number on line
			for numberPos, numVal := range strings.Fields(boardLine) {
				fmt.Printf("Position: %d: %s\n", numberPos, numVal)
				num, _ := strconv.Atoi(numVal)
				line[numberPos] = BingoNumber{ value: num, marked: false}
			}
			game.board[lineNum] = line
		}
		fmt.Println("-------")

		bingoGames = append(bingoGames, game)
	}

	return bingoGames
}