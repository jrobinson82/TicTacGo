package main

import (
	"fmt"
	"strconv"
)

const (
	version   = "0.2.0"
	moveEmpty = "[ ]"
	moveX     = "[x]"
	moveO     = "[o]"
	playerX   = "x"
	playerO   = "o"
)

type gameBlock struct {
	cell  int
	value string
}

func isPlayerWinner(board [][]gameBlock, player string) bool {
	isWinner := false
	var move string

	if player == playerX {
		move = moveX
	} else {
		move = moveO
	}

	// Handles 3 by 3 winds only...

	// --- Check rows ---
	for r := 0; r < 3; r++ {
		isWinner = true
		for c := 0; c < 3; c++ {
			if board[c][r].value != move {
				isWinner = false
				break
			}
		}
		if isWinner {
			break
		}
	}

	// --- Check columns ---
	if !isWinner {
		for c := 0; c < 3; c++ {
			isWinner = true
			for r := 0; r < 3; r++ {
				if board[c][r].value != move {
					isWinner = false
					break
				}
			}
			if isWinner {
				break
			}
		}
	}

	// --- Check diagonal (top-left → bottom-right) ---
	if !isWinner {
		isWinner = true
		for i := 0; i < 3; i++ {
			if board[i][i].value != move {
				isWinner = false
				break
			}
		}
	}

	// --- Check diagonal (top-right → bottom-left) ---
	if !isWinner {
		isWinner = true
		for i := 0; i < 3; i++ {
			if board[2-i][i].value != move {
				isWinner = false
				break
			}
		}
	}

	return isWinner
}

func editGameBlock(board [][]gameBlock, cell int, value string) {
	found := false

	for col := range board {
		if !found {
			for row := range board[col] {
				if cell == board[col][row].cell {
					board[col][row].value = value
					break
				}
			}
		} else {
			break
		}
	}
}

func createGameBoard(width, height int) [][]gameBlock {
	board := make([][]gameBlock, height)

	cell := 1

	for col := 0; col < height; col++ {
		board[col] = make([]gameBlock, width)

		for row := 0; row < width; row++ {
			board[col][row].cell = cell
			board[col][row].value = moveEmpty

			cell++
		}
	}

	return board
}

func clearGameBoard(board [][]gameBlock) {
	for col := range board {
		for row := range board[col] {
			board[col][row].value = moveEmpty
		}
	}
}

func renderGameBoard(board [][]gameBlock, showCellNumbers ...bool) {
	cellNumbers := false

	if len(showCellNumbers) > 0 {
		cellNumbers = showCellNumbers[0]
	}

	fmt.Println("")

	for col := range board {
		for row := range board[col] {
			if !cellNumbers {
				fmt.Printf("%s", board[col][row].value)
			} else {
				fmt.Printf("%d%s ", board[col][row].cell, board[col][row].value)
			}
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

	height := 3
	width := 3
	cellRange := height * width

	gameBoard := createGameBoard(width, height)

	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Println("****** Welcome To TIC TAC [GO]! ******")
	fmt.Println("--------------------------------------")
	fmt.Printf("\nVersion %s\n\n", version)
	fmt.Println("Instructions:")
	fmt.Println(`When it's you turn type a cell number (1-9)`)
	fmt.Println(`and press Enter to mark a square.`)
	fmt.Println(`Type "q" to quit game.`)
	fmt.Println(`Here are the cell numbers...`)

	clearGameBoard(gameBoard)
	renderGameBoard(gameBoard, true)

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
				if num >= 1 && num <= cellRange {

					switch playerTurn {
					case playerX:
						editGameBlock(gameBoard, num, moveX)
						validInput = true

					case playerO:
						editGameBlock(gameBoard, num, moveO)
						validInput = true

					default:
						// do nothing... for now.
					}

				} else {
					fmt.Printf(`"%d" is out of range. Please select a number between 1 and %d.`, num, cellRange)
				}
			}
		}

		fmt.Println()
		fmt.Println("#####################################")
		fmt.Println()

		if validInput {
			if !isPlayerWinner(gameBoard, playerTurn) {
				turnToggle(&playerTurn)
			} else {
				renderGameBoard(gameBoard[:])
				fmt.Printf(`Player "%s" Wins!`, playerTurn)
				gameOn = false
			}
		}
	}
}
