package cmd

import (
	"golang-learnings/pigGame"
	"os"

	"github.com/spf13/cobra"
)

var problemStatement string
var limit int

var rootCmd = &cobra.Command{
	Use:   "practice",
	Short: "A simple CLI application",
	Long:  `A simple CLI application to demonstrate the usage of Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().HasFlags() {
			cmd.Usage()
			return
		}

		switch problemStatement {
		case "piggame":
			player1Strategy := os.Args[3]
			player2Strategy := os.Args[4]
			pigGame.PigGame(player1Strategy, player2Strategy)

		case "quiz1":

		default:
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&problemStatement, "problemStatement", "p", "", "Problem Statement")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Limit for quiz 1")

}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
