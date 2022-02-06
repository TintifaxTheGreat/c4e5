package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"math/rand"
)

func (g *Game) FindMove() dragontoothmg.Move {
	moves := g.Board.GenerateLegalMoves()
	return moves[rand.Intn(len(moves))]
}
