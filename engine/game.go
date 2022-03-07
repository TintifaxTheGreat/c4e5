package engine

import "github.com/dylhunn/dragontoothmg"

const initDepth = 4
const initQuietDepth = 3

type Game struct {
	Depth      int
	QuietDepth int
	Board      dragontoothmg.Board
	Playing    bool
	HashMap    *HashMap
}

func NewGame(fen string) *Game {
	if fen == "" {
		fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}
	return &Game{
		Depth:   initDepth,
		Board:   dragontoothmg.ParseFen(fen),
		Playing: true,
		HashMap: NewHashMap(),
	}
}
