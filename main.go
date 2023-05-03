package main

import (
	"fmt"
	"golang-learnings/piggame"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run ./main <player1_strategy> <player2_strategy>")
		return
	}
	piggame.PigGame(os.Args)
}
