package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"sort"
)

const maxInt int = 1000000
const minInt int = -1000000

func (g *Game) FindMove() dragontoothmg.Move {
	hashmap := NewHashMap() // TODO think if this is wise

	moves := g.Board.GenerateLegalMoves()
	if len(moves) == 1 {
		return moves[0]
	}

	alpha := minInt
	beta := maxInt

	var bestMove dragontoothmg.Move = 0
	//var sortedMoves []dragontoothmg.Move TODO

	curDepth := 0
	for curDepth <= g.Depth {

		priorValues := make(map[dragontoothmg.Move]int)

		for _, move := range moves {
			unapplyFunc := g.Board.Apply(move)
			priorValues[move] = -negamax(&g.Board, hashmap, curDepth, -beta, -alpha)
			unapplyFunc()
		}

		// TODO consider winning move

		sortedMoves := make([]dragontoothmg.Move, 0, len(priorValues))
		for key := range priorValues {
			sortedMoves = append(sortedMoves, key)
		}
		sort.Slice(sortedMoves, func(i, j int) bool {
			return priorValues[sortedMoves[i]] > priorValues[sortedMoves[j]]
		})

		bestMove = sortedMoves[0]
		moves = sortedMoves
		curDepth += 2 // TODO think about this
	}

	if bestMove == 0 {
		panic("no move found")
	}
	return bestMove

}
