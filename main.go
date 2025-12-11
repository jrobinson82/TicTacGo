package main

import (
	"fmt"
)

func main() {
	var input string
	gameOn := true

	fmt.Println(`Type "q" and Enter to end game:`)

	// Main game loop
	for gameOn == true {
		fmt.Println("Game: On> ")
		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("Game: Off")
			break
		}

		fmt.Println("Input: ", input)
	}
}
