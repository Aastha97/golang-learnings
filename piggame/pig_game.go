package piggame

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
	score, turnTotal := 0, 0
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

func PigGame(args []string) {
	rangeSplit1 := strings.Split(os.Args[1], "-")

	start1, _ := strconv.Atoi(rangeSplit1[0])

	end1, _ := strconv.Atoi(rangeSplit1[1])

	rangeSplit2 := strings.Split(os.Args[2], "-")

	start2, _ := strconv.Atoi(rangeSplit2[0])

	end2, _ := strconv.Atoi(rangeSplit2[1])

	for j := start1; j <= end1; j++ {
		p1 := players("Holding at "+strconv.Itoa(j), func(score, turnTotal int) bool {
			return turnTotal >= j
		})
		p1TotalWins, p2TotalWins, totalgames := 0, 0, 0

		for i := start2; i <= end2; i++ {
			if i != j {
				p2 := players("Holding at "+strconv.Itoa(i), func(score, turnTotal int) bool {
					return turnTotal >= i
				})
				p1Wins, p2Wins := 0, 0

				for i := 0; i < 10; i++ {
					p1Score, p2Score := 0, 0

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
				p1TotalWins += p1Wins
				p2TotalWins += p2Wins
				totalgames = p1TotalWins + p2TotalWins
			}
		}
		fmt.Printf("Result: Wins, losses staying at k =  %d: %d/%d (%.1f%%), %d/%d (%.1f%%)\n", j, p1TotalWins, totalgames, float64(p1TotalWins*100/totalgames), p2TotalWins, totalgames, float64(p2TotalWins*100/totalgames))
	}

}
