package engine

import "github.com/dylhunn/dragontoothmg"

const initDepth = 99
const initQuietDepth = 2
const maxTime = 10000 // in milliseconds

type Game struct {
	Depth      int
	QuietDepth int
	Board      dragontoothmg.Board
	Playing    bool
}

func NewGame(fen string) *Game {
	if fen == "" {
		fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}
	return &Game{
		Depth:      initDepth,
		QuietDepth: initQuietDepth,
		Board:      dragontoothmg.ParseFen(fen),
		Playing:    true,
		//HashMap: NewHashMap(),
	}
}
