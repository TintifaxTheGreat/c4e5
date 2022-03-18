package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"log"
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

	cacheMiss = 0
	cacheHit = 0

	var bestMove dragontoothmg.Move = 0

	curDepth := 0
	for curDepth <= g.Depth {

		priorValues := make(map[dragontoothmg.Move]int)

		for i, move := range moves {
			if i > pruneWorseIndex(curDepth) {
				continue
			}

			unapplyFunc := g.Board.Apply(move)
			v, _, ok := hashmap.Get(curDepth, &g.Board)

			if ok {
				priorValues[move] = -v
			} else {
				priorValues[move], _ = negamax(&g.Board, hashmap, curDepth, -beta, -alpha, false)
				priorValues[move] *= -1
			}
			unapplyFunc()
		}

		sortedMoves := make([]dragontoothmg.Move, 0, len(priorValues))
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

	log.Print(cacheHit, " ", cacheMiss, " ", float64(cacheHit)/float64(cacheMiss+cacheHit))
	return bestMove

}
func pruneWorseIndex(depth int) int {
	switch depth {
	case 0:
		return 9999
	case 1:
		return 9999
	case 2:
		return 9999
	case 3:
		return 8
	case 4:
		return 8
	default:
		return 4
	}
}
