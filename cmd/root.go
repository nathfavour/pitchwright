package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pitchwright",
	Short: "Generate stunning, intelligent pitch decks for your project.",
	Long:  `Pitchwright analyzes your project and generates professional pitch decks with advanced LLMs and animations.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement main pitch deck generation logic
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
