package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "c4e5",
	Short: "c4e5 is a UCI compatible chess engine.",
	Long:  "c4e5 is a UCI compatible chess engine.",
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
