package main

import (
	"fmt"
)

type GameSquare struct {
	player int
	move   string
}

func clearGameBoard(gameBoard []GameSquare) {
	for i := 0; i < 9; i++ {
		gameBoard[i].player = 0
		gameBoard[i].move = "[ ]"

	}
}

func renderGameBoard(gameBoard []GameSquare) {
	i := 0

	fmt.Println("")

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			fmt.Printf("%s", gameBoard[i].move)
			i++
		}
		fmt.Println("")
	}

	fmt.Println("")
}

func main() {
	var input string
	gameOn := true
	playerTurn := 1
	var gameBoard [9]GameSquare

	fmt.Println(`Type "q" and Enter to end game:`)

	clearGameBoard(gameBoard[:])

	// Main game loop
	for gameOn {
		fmt.Println("Game: On> ")

		renderGameBoard(gameBoard[:])

		fmt.Printf(`Ready Player %d (Type "x","o" and Enter): `, playerTurn)

		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("Game: Off")
			break
		}

		fmt.Println("Input: ", input)
	}
}
