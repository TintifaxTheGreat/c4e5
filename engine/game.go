package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"time"
)

type Game struct {
	MaxDepth      int
	IncQuietDepth int
	Board         dragontoothmg.Board
	Playing       bool
	MoveTime      time.Duration
	BoardHistory  BoardHistory
}

func NewGame(fen string, maxDepth, incQuietDepth int, moveTime time.Duration) *Game {
	if fen == "" {
		fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}

	if moveTime == 0 {
		moveTime = defaultTime
	}

	if maxDepth == 0 {
		maxDepth = maxDepth
	}

	if incQuietDepth == 0 {
		incQuietDepth = initQuietDepth
	}

	return &Game{
		MaxDepth:      maxDepth,
		IncQuietDepth: initQuietDepth,
		Board:         dragontoothmg.ParseFen(fen),
		Playing:       true,
		MoveTime:      moveTime,
		BoardHistory:  *NewBoardHistory(),
	}
}
