package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"log"
	"sort"
)

func (g *Game) FindMove() dragontoothmg.Move {
	hashmap := NewHashMap()

	moves := g.Board.GenerateLegalMoves()
	if len(moves) == 1 {
		return moves[0]
	}

	alpha := minInt
	beta := maxInt

	var bestMove dragontoothmg.Move = 0

	curDepth := 0
	for curDepth <= g.Depth {

		priorValues := make(map[dragontoothmg.Move]int)

		for i, move := range moves {
			unapplyFunc := g.Board.Apply(move)
			v, _, ok := hashmap.Get(curDepth, &g.Board)

			if ok {
				priorValues[move] = -v
			} else {
				priorValues[move], _ = g.negamax(hashmap, curDepth, -beta, -alpha, false, false)
				priorValues[move] *= -1
			}
			unapplyFunc()
			priorValues[move] -= i
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
		bestValue := priorValues[bestMove]

		if bestValue > mateLevel {
			break
		}

		cutIndex := len(sortedMoves)
		if curDepth > latePruningDepth {
			for i, move := range sortedMoves {
				if priorValues[move] < bestValue-pruneThreshold {
					cutIndex = i
					log.Print("cut at ", i)
					break
				}
			}
		}
		moves = sortedMoves[:cutIndex]

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

	//log.Print(cacheHit, " ", cacheMiss, " ", float64(cacheHit)/float64(cacheMiss+cacheHit))
	return bestMove

}
