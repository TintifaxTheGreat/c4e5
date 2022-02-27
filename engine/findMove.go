package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"sort"
)

const maxInt int = 1000000
const minInt int = -1000000

func (g *Game) FindMove() dragontoothmg.Move {
	moves := g.Board.GenerateLegalMoves()
	if len(moves) == 1 {
		return moves[0]
	}
	alpha := minInt
	beta := maxInt
	index := minInt

	priorValues := make(map[dragontoothmg.Move]int)
	for _, move := range moves {
		unapplyFunc := g.Board.Apply(move)
		//priorValues[move] = -evaluate(&g.Board)
		priorValues[move] = -negamax(&g.Board, 2, -beta, -alpha)
		unapplyFunc()
	}

	keys := make([]dragontoothmg.Move, 0, len(priorValues))
	for key := range priorValues {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return priorValues[keys[i]] > priorValues[keys[j]]
	})

	for i, move := range keys {
		unapplyFunc := g.Board.Apply(move)
		value := -negamax(&g.Board, g.Depth, -beta, -alpha)
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

	if index == minInt {
		panic("no move found")
	}

	return keys[index]

}
