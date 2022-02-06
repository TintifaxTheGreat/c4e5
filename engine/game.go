package engine

import "github.com/dylhunn/dragontoothmg"

type Game struct {
	Fen     string // remove this
	Board   dragontoothmg.Board
	Playing bool
}
