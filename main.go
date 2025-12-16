package main

import (
	"fmt"
)

type GameSquare struct {
	player int
	move   string
}

func main() {
	var input string
	gameOn := true
	var gameBoard [9]GameSquare

	// Clear Game Board
	for i := 0; i < 9; i++ {
		gameBoard[i].player = 0
		gameBoard[i].move = "[-]"

		fmt.Printf("Game | Square: %d Player: %d Move: %v \n", i+1, gameBoard[i].player, gameBoard[i].move)
	}

	fmt.Println(`Type "q" and Enter to end game:`)

	// Main game loop
	for gameOn {
		fmt.Println("Game: On> ")
		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("Game: Off")
			break
		}

		fmt.Println("Input: ", input)
	}
}
