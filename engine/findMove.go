package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"sort"
)

const MAX_INT int = 1000000
const MIN_INT int = -1000000

const DEPTH = 5
const QDEPTH = 3

func (g *Game) FindMove() dragontoothmg.Move {
	moves := g.Board.GenerateLegalMoves()
	if len(moves) == 1 {
		return moves[0]
	}
	alpha := MIN_INT
	beta := MAX_INT
	index := MIN_INT

	priorValues := make(map[dragontoothmg.Move]int)
	for _, move := range moves {
		unapplyFunc := g.Board.Apply(move)
		priorValues[move] = -evaluate(&g.Board)
		unapplyFunc()
	}

	keys := make([]dragontoothmg.Move, 0, len(priorValues))
	for key := range priorValues {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		// TODO order checks first
		return priorValues[keys[i]] > priorValues[keys[j]]
	})

	for i, move := range keys {
		unapplyFunc := g.Board.Apply(move)
		value := -negamax(&g.Board, DEPTH, -beta, -alpha)
		unapplyFunc()
		if value >= beta {
			index = i
			break
		}
		if value > alpha {
			alpha = value
			index = i
		}
	}

	if index == MIN_INT {
		panic("no move found")
	}

	return keys[index]

}
