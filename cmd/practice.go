package cmd

import (
	"fmt"
	"golang-learnings/pigGame"
	"os"

	"github.com/spf13/cobra"
)

var problemStatement int

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
		case 1:
			fmt.Println("")
		default:
			if len(os.Args) != 3 {
				fmt.Println("Usage: go run ./golang-learnings <player1_strategy> <player2_strategy>")
				return
			}
			pigGame.PigGame(os.Args)
		}
	},
}

func init() {
	rootCmd.Flags().IntVarP(&problemStatement, "problemStatement", "p", 0, "Problem Statement")

}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
