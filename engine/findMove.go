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

	var bestMove dragontoothmg.Move = 0
	var sortedMoves []dragontoothmg.Move

	curDepth := 1
	for curDepth <= g.Depth {

		priorValues := make(map[dragontoothmg.Move]int)

		for i, move := range moves {
			unapplyFunc := g.Board.Apply(move)
			v, ok := g.HashMap.Get(curDepth, &g.Board)
			if ok {
				priorValues[move] = v
			} else {
				priorValues[move] = -negamax(&g.Board, g.HashMap, curDepth, -beta, -alpha)
				g.HashMap.Put(curDepth, priorValues[move], &g.Board)
			}
			unapplyFunc()
		}

		// TODO consider winning move

		sortedMoves = make([]dragontoothmg.Move, 0, len(priorValues))
		for key := range priorValues {
			sortedMoves = append(sortedMoves, key)
		}
		sort.Slice(sortedMoves, func(i, j int) bool {
			return priorValues[sortedMoves[i]] > priorValues[sortedMoves[j]]
		})

		bestMove = sortedMoves[0]
		moves = sortedMoves
		curDepth++
	}

	if bestMove == 0 {
		panic("no move found")
	}

	return bestMove

}
