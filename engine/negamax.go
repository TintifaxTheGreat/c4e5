package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

func (g *Game) negamax(hashmap *HashMap, depth, quietDepth, alpha, beta int, unsorted bool) (int, dragontoothmg.Move) {
	v, priorBestMove, ok := hashmap.Get(depth, &g.Board)
	if ok {
		return v, priorBestMove
	}

	children := g.Board.GenerateLegalMoves()
	if len(children) == 0 {
		if g.Board.OurKingInCheck() == true {
			value := -40000 - depth
			hashmap.Put(maxInt, value, &g.Board, 0)
			return value, 0
		}
		hashmap.Put(maxInt, 0, &g.Board, 0)
		return 0, 0
	}

	if !g.Playing {
		return 0, 0
	}

	if depth < 1 {
		value := evaluate(&g.Board)
		hashmap.Put(0, value, &g.Board, 0)
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
		if isCapture && (quietDepth > 0) {
			quietDepth--
			newDepth = depth
		} else {
			newDepth = depth - 1
		}

		if pvs {
			value, valueMove = g.negamax(hashmap, newDepth, quietDepth, -beta, -alpha, true)
			value *= -1

		} else {
			value, valueMove = g.negamax(hashmap, newDepth, quietDepth, -alpha-1, -alpha, true)
			value *= -1
			if value > alpha {
				value, valueMove = g.negamax(hashmap, newDepth, quietDepth, -beta, -alpha, true)
				value *= -1
			}
		}
		unapplyFunc()

		if value >= beta {
			hashmap.Put(depth-1, beta, &g.Board, bestMove)
			return beta, bestMove
		}
		if value > alpha {
			alpha = value
			bestMove = valueMove
			pvs = false
		}
	}
	hashmap.Put(depth-1, alpha, &g.Board, bestMove)
	return alpha, bestMove
}

// TODO from where is this snippet?
func testCapture(m dragontoothmg.Move, b *dragontoothmg.Board) bool {
	toBitboard := (uint64(1) << m.To())
	return (toBitboard&b.White.All != 0) || (toBitboard&b.Black.All != 0)
}
