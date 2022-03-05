package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

func negamax(board *dragontoothmg.Board, hashmap *HashMap, depth int, alpha int, beta int) int {
	children := board.GenerateLegalMoves()
	if len(children) == 0 {
		if board.OurKingInCheck() == true {
			return -40000 - depth
		}
		return 0
	}

	if depth < 1 {
		return evaluate(board)
	}

	for _, child := range children {
		value := 0
		//isCapture := testCapture(child, board)
		unapplyFunc := board.Apply(child)

		v, ok := hashmap.Get(depth, board)
		if ok {
			value = v
		} else {
			value = -negamax(board, hashmap, depth-1, -beta, -alpha)
			hashmap.Put(depth, value, board)
		}
		unapplyFunc()

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
