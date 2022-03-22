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
			/*
				if i > pruneWorseIndex(curDepth) { //TODO think about this
					continue
				}

			*/

			unapplyFunc := g.Board.Apply(move)
			v, _, ok := hashmap.Get(curDepth, &g.Board)

			if ok {
				priorValues[move] = -v
			} else {
				priorValues[move], _ = g.negamax(hashmap, curDepth, g.QuietDepth, -beta, -alpha, false)
				priorValues[move] *= -1
			}
			unapplyFunc()
			priorValues[move] -= i // TODO is this wise?
		}

		sortedMoves := make([]dragontoothmg.Move, 0, len(priorValues))
		for key := range priorValues {
			sortedMoves = append(sortedMoves, key)
		}
		sort.Slice(sortedMoves, func(i, j int) bool {
			return priorValues[sortedMoves[i]] >= priorValues[sortedMoves[j]]
		})

		if !g.Playing {
			break
		}

		bestMove = sortedMoves[0]

		if priorValues[bestMove] > mateLevel {
			break
		}

		moves = sortedMoves

		log.Print("\nDepth: ", curDepth)
		/*
			for _, m := range moves {
				log.Print(m.String(), " ", priorValues[m])
			}
		*/

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
		return maxInt
	case 1:
		return maxInt
	case 2:
		return maxInt
	case 3:
		return maxInt
	case 4:
		return 16
	default:
		return 4
	}
}
