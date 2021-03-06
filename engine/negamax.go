package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

func (g *Game) negamax(hashmap *hashMap, depth, alpha, beta int, unsorted, isQuiescence bool) (int, dragontoothmg.Move) {
	v, priorBestMove, ok := hashmap.get(depth, &g.Board)
	if ok {
		return v, priorBestMove
	}

	if g.testBoardHistory() > 2 {
		return 0, 0 // TODO improve this
	}

	children := generateMovesPrime(&g.Board)
	if len(children) == 0 {
		if g.Board.OurKingInCheck() == true {
			value := -mate - depth
			hashmap.put(maxInt, value, &g.Board, 0)
			return value, 0
		}
		hashmap.put(maxInt, 0, &g.Board, 0)
		return 0, 0
	}

	if !g.Playing {
		return 0, 0
	}

	if depth < 1 {
		value := evaluate(&g.Board)
		hashmap.put(0, value, &g.Board, 0)
		return value, 0
	}

	if unsorted && (priorBestMove != 0) {
		for i, child := range children {
			if child == priorBestMove {
				children[i] = children[0]
				children[0] = priorBestMove
			}
		}
	}

	var bestMove dragontoothmg.Move = 0

	pvs := true
	for _, child := range children {

		value := 0
		var valueMove dragontoothmg.Move = 0

		isCapture := testCapture(child, &g.Board)

		unapplyFunc := g.Board.Apply(child)
		var newDepth int

		if !unsorted && pvs {
			newDepth += pvsDepth
		} else {

			if depth == 1 && isCapture && !isQuiescence {
				isQuiescence = true
				newDepth = depth + g.incQuietDepth - 1
			} else {
				newDepth = depth - 1
			}
		}

		if pvs {
			value, valueMove = g.negamax(hashmap, newDepth, -beta, -alpha, true, isQuiescence)
			value *= -1

		} else {
			value, valueMove = g.negamax(hashmap, newDepth, -alpha-1, -alpha, true, isQuiescence)
			value *= -1
			if value > alpha {
				value, valueMove = g.negamax(hashmap, newDepth, -beta, -alpha, true, isQuiescence)
				value *= -1
			}
		}

		unapplyFunc()

		if value >= beta {
			//hashmap.put(depth-1, beta, &g.Board, bestMove)
			return beta, bestMove
		}
		if value > alpha {
			alpha = value
			bestMove = valueMove
			pvs = false
		}
	}
	hashmap.put(depth-1, alpha, &g.Board, bestMove)
	return alpha, bestMove
}
