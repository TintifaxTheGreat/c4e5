package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

func negamax(board *dragontoothmg.Board, hashmap *HashMap, depth int, alpha int, beta int, unsorted bool) (int, dragontoothmg.Move) {
	v, priorBestMove, ok := hashmap.Get(depth, board)
	//ok = false //TODO
	if ok {
		return v, priorBestMove
	}

	children := board.GenerateLegalMoves()
	if len(children) == 0 {
		if board.OurKingInCheck() == true {
			value := -40000 - depth
			hashmap.Put(999, value, board, 0)
			return value, 0
		}
		hashmap.Put(999, 0, board, 0)
		return 0, 0
	}

	if depth < 1 {
		value := evaluate(board)
		hashmap.Put(0, value, board, 0)
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

		isCapture := testCapture(child, board)

		unapplyFunc := board.Apply(child)

		if (depth == 1) && isCapture {
			value = -quiescense(board, depth-1+initQuietDepth, -beta, -alpha)
		} else {
			if pvs {
				value, valueMove = negamax(board, hashmap, depth-1, -beta, -alpha, true)
				value *= -1

			} else {
				value, valueMove = negamax(board, hashmap, depth-1, -alpha-1, -alpha, true)
				value *= -1
				if value > alpha {
					value, valueMove = negamax(board, hashmap, depth-1, -beta, -alpha, true)
					value *= -1
				}
			}
		}

		unapplyFunc()

		if value >= beta {
			hashmap.Put(depth-1, beta, board, bestMove)
			return beta, bestMove
		}
		if value > alpha {
			alpha = value
			bestMove = valueMove
			pvs = false
		}

	}
	hashmap.Put(depth-1, alpha, board, bestMove)
	return alpha, bestMove
}

func quiescense(board *dragontoothmg.Board, depth int, alpha int, beta int) int {
	children := board.GenerateLegalMoves()
	if len(children) == 0 {
		return -40000 - depth
		// TODO consider draws
	}

	staticValue := evaluate(board)

	if depth < 1 {
		return staticValue
	}

	var unapplyFunc func()
	for _, child := range children {
		value := 0
		if testCapture(child, board) {
			value = staticValue
		} else {
			unapplyFunc = board.Apply(child)
			value = -evaluate(board)
			unapplyFunc()
		}

		if value >= beta {
			return beta
		}
		if value > alpha {
			alpha = value
		}
	}
	return alpha
}

// TODO from where is this snippet?
func testCapture(m dragontoothmg.Move, b *dragontoothmg.Board) bool {
	toBitboard := (uint64(1) << m.To())
	return (toBitboard&b.White.All != 0) || (toBitboard&b.Black.All != 0)
}
