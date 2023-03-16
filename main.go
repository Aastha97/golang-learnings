package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	name     string
	strategy func(score, turnTotal int) bool
}

func players(name string, strategy func(score, turnTotal int) bool) Player {
	return Player{name, strategy}
}

func (p Player) Play() int {
	score := 0
	turnTotal := 0
	for {
		roll := rand.Intn(6) + 1
		if roll == 1 {
			turnTotal = 0
			break
		} else {
			turnTotal += roll
			if p.strategy(score, turnTotal) {
				score += turnTotal
				break
			}
		}
	}
	return score
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run ./main <player1_strategy> <player2_strategy>")
		return
	}

	rangeSplit1 := strings.Split(os.Args[1], "-")

	start1, _ := strconv.Atoi(rangeSplit1[0])

	//end1, _ := strconv.Atoi(rangeSplit1[1])

	rangeSplit2 := strings.Split(os.Args[2], "-")

	start2, _ := strconv.Atoi(rangeSplit2[0])

	end2, _ := strconv.Atoi(rangeSplit2[1])

	p1 := players("Holding at "+strconv.Itoa(start1), func(score, turnTotal int) bool {
		return turnTotal >= start1
	})

	for i := start2; i <= end2; i++ {
		if i != start1 {
			p2 := players("Holding at "+strconv.Itoa(i), func(score, turnTotal int) bool {
				return turnTotal >= i
			})
			p1Wins := 0
			p2Wins := 0

			for i := 0; i < 10; i++ {
				p1Score := 0
				p2Score := 0

				for p1Score < 100 && p2Score < 100 {
					p1Score += p1.Play()
					if p1Score >= 100 {
						p1Wins++
						break
					}

					p2Score += p2.Play()
					if p2Score >= 100 {
						p2Wins++
						break
					}
				}

			}
			fmt.Printf("%s vs %s: wins: %d/10 (%.1f%%), losses: %d/10 (%.1f%%)\n",
				p1.name, p2.name, p1Wins, float64(p1Wins)*100/10, 10-p1Wins, float64(10-p1Wins)*100/10)
		}

	}

}
