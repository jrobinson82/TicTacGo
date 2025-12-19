package main

import (
	"fmt"
	"strconv"
)

func clearGameBoard(gameBoard []string) {
	for i := 0; i < 9; i++ {
		gameBoard[i] = "[ ]"
	}
}

func renderGameBoard(gameBoard []string, showCellNumbers ...bool) {
	i := 0
	cellNumbers := false

	if len(showCellNumbers) > 0 {
		cellNumbers = showCellNumbers[0]
	}

	fmt.Println("")

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if !cellNumbers {
				fmt.Printf("%s", gameBoard[i])
			} else {
				fmt.Printf("%d%s ", i+1, gameBoard[i])
			}

			i++
		}
		fmt.Println("")
	}

	fmt.Println("")
}

func turnToggle(turn *string) {
	if *turn == "x" {
		*turn = "o"
	} else {
		*turn = "x"
	}
}

func main() {
	var input string
	gameOn := true
	validInput := false
	playerTurn := "x"
	var gameBoard [9]string

	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Println("****** Welcome To TIC TAC [GO]! ******")
	fmt.Println("--------------------------------------")
	fmt.Println("Instructions:")
	fmt.Println(`When it's you turn type a cell number (1-9)`)
	fmt.Println(`and press Enter to mark a square.`)
	fmt.Println(`Type "q" to quit game.`)
	fmt.Println(`Here are the cell numbers...`)

	clearGameBoard(gameBoard[:])
	renderGameBoard(gameBoard[:], true)

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")
	fmt.Println()

	// Main game loop
	for gameOn {
		fmt.Println("Game: On> ")

		renderGameBoard(gameBoard[:])

		fmt.Printf(`Ready Player "%s"> `, playerTurn)
		fmt.Scanln(&input)

		num, err := strconv.Atoi(input)

		if input == "q" {
			validInput = true
			fmt.Println("Game: Off<")
			break
		} else {
			if err != nil {
				fmt.Printf(`"%s" is an invalid command.`, input)
			} else {
				if num >= 1 && num <= 9 {
					validInput = true
					gameBoard[num-1] = "[" + playerTurn + "]"
				} else {
					fmt.Printf(`"%d" is out of range. Please select a number between 1 and 9.`, num)
				}
			}
		}

		fmt.Println()
		fmt.Println("#####################################")
		fmt.Println()

		if validInput {
			turnToggle(&playerTurn)
		}
	}
}
