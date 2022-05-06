package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"time"
)

type Game struct {
	MaxDepth      int
	incQuietDepth int
	Board         dragontoothmg.Board
	Playing       bool
	MoveTime      time.Duration
	boardHistory  boardHistory
}

func NewGame(fen string, maxDepth, incQuietDepth int, moveTime time.Duration) *Game {
	if fen == "" {
		fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}

	if moveTime == 0 {
		moveTime = defaultTime
	}

	if maxDepth == 0 {
		maxDepth = initMaxDepth
	}

	if incQuietDepth == 0 {
		incQuietDepth = initQuietDepth
	}

	return &Game{
		MaxDepth:      maxDepth,
		incQuietDepth: initQuietDepth,
		Board:         dragontoothmg.ParseFen(fen),
		Playing:       true,
		MoveTime:      moveTime,
		boardHistory:  *newBoardHistory(),
	}
}
