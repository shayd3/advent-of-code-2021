package days

import (
	"fmt"
	"io/ioutil"
	"log"
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

type Winner struct {
	game BingoBoard
	lastCalledNumber int
}

// Day04 Solution for AoC
// 1. Create boards
// 2. Go through BingoCalls and Mark BingoNumbers
// 3. Check if any board is a winner after 5 numbers called (Cols or row, no diagonal)
// 4. When a winner is found, add all non-marked numbers
func Day04() {
	log.SetOutput(ioutil.Discard)
	bingoCalls := inputs.BingoCalls
	bingoBoards := inputs.BingoBoards
	
	games := makeBoards(bingoBoards)
	winner := playGames(games, bingoCalls)
	fmt.Printf("We have a winner!\nWinning game: %v\nLast called number: %d\n", winner.game.board, winner.lastCalledNumber)
	fmt.Printf("Final standings: %d\n", calculateWinningBoard(winner))
}

func makeBoards(boardInput string) (bingoGames []BingoBoard) {
	boardStrings := strings.Split(string(boardInput), "\n\n")
	// Get each board block from string
	for gameNumber, boardString := range boardStrings {
		log.Printf("Game #%d:\n%s\n", gameNumber, boardString)
		boardLineStrings := strings.Split(boardString, "\n")
		game := BingoBoard{}
		// Loop each line in board
		for lineNum, boardLine := range boardLineStrings {
			log.Printf("Line %d => %s\n", lineNum, boardLine)
			line := [5]BingoNumber{}
			// Get each number on line
			for numberPos, numVal := range strings.Fields(boardLine) {
				log.Printf("Position: %d: %s\n", numberPos, numVal)
				num, _ := strconv.Atoi(numVal)
				line[numberPos] = BingoNumber{ value: num, marked: false}
			}
			game.board[lineNum] = line
		}
		log.Println("-------")

		bingoGames = append(bingoGames, game)
	}

	return bingoGames
}

func playGames(games []BingoBoard, bingoCalls []int) Winner {
	winner := Winner{}
	for turn, call := range bingoCalls {
		log.Printf("Turn #%d : Number called: %d\n", turn, call)
		// Loop over each game
		for gameNumber, game := range games {
			for rowNumber, row := range game.board {
				for numPosition, num := range row {
					if !num.marked && num.value == call {
						log.Printf("Number found! Game: %d, Row: %d, Found: %d\n", gameNumber, rowNumber, call)
						games[gameNumber].board[rowNumber][numPosition].marked = true
						if checkBoard(games[gameNumber]){
							winner.lastCalledNumber = call
							winner.game = games[gameNumber]
							return winner
						}
					}
				}
			}
			
		}
	}

	return winner
}

func checkBoard(game BingoBoard) bool {
	isWinner := false
	// Check rows
	for _, row := range game.board {
		if row[0].marked && 
		row[1].marked && 
		row[2].marked && 
		row[3].marked && 
		row[4].marked {
			isWinner = true
		}
	}
	// Check cols
	for colPos := range game.board[0] {
		if game.board[0][colPos].marked && 
		game.board[1][colPos].marked && 
		game.board[2][colPos].marked && 
		game.board[3][colPos].marked && 	
		game.board[4][colPos].marked {
				isWinner = true
		}
	}

	return isWinner
}

func calculateWinningBoard(winner Winner) int {
	total := 0
	for _, row := range winner.game.board {
		for _, num := range row {
			if !num.marked {
				total += num.value
			}
		}
	}

	return total * winner.lastCalledNumber
}