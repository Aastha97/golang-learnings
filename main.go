package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run ./main <player1_strategy> <player2_strategy>")
		return
	}

	pigGame(os.Args)
}
