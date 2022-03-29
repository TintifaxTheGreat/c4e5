package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"time"
)

const initQuietDepth = 1 // TODO 1j
const maxDepth = 99
const defaultTime = 10000 * time.Millisecond // TODO was 10000
const mateLevel = 35000

type Game struct {
	Depth      int
	QuietDepth int
	Board      dragontoothmg.Board
	Playing    bool
	MoveTime   time.Duration
}

func NewGame(fen string, depth, quietDepth int, moveTime time.Duration) *Game {
	if fen == "" {
		fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}

	if moveTime == 0 {
		moveTime = defaultTime
	}

	if depth == 0 {
		depth = maxDepth
	}

	if quietDepth == 0 {
		quietDepth = initQuietDepth
	}

	return &Game{
		Depth:      depth,
		QuietDepth: initQuietDepth,
		Board:      dragontoothmg.ParseFen(fen),
		Playing:    true,
		MoveTime:   moveTime,
		//HashMap: NewHashMap(),
	}
}
