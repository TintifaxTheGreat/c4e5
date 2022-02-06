package cmd

import (
	"github.com/dylhunn/dragontoothmg"
	"github.com/spf13/cobra"
	"github.com/tintifaxthegreat/c4e5/engine"
	"log"
	"strconv"
	"time"
)

var depth string

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringVarP(&depth, "depth", "d", "1", "levels of perft to be calculated")
}

var versionCmd = &cobra.Command{
	Use:   "perft",
	Short: "Perft",
	Long:  "Just perft",
	Run: func(cmd *cobra.Command, args []string) {
		level, _ := strconv.Atoi(depth)

		board := dragontoothmg.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
		start := time.Now()
		count := engine.Perft(&board, level)
		duration := time.Since(start)
		log.Print(count)
		log.Print(duration.Seconds())
	},
}
