package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

func negamax(board dragontoothmg.Board, depth int, alpha int, beta int, quiescence bool) int {
	children := board.GenerateLegalMoves()
	if len(children) == 0 {
		return -40000 - depth
		// TODO consider draws
	}

	if depth < 1 {
		return evaluate(board)
	}

	for i, child := range children {
		if quiescence && i > 3 {
			return alpha
		}
		value := 0
		unapplyFunc := board.Apply(child)
		if !quiescence && (isCapture(child, &board) || board.OurKingInCheck()) {
			value = -negamax(board, depth+1, -beta, -alpha, true)
		} else {
			value = -negamax(board, depth-1, -beta, -alpha, quiescence)
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

// TODO remove this
func isCapture(m dragontoothmg.Move, b *dragontoothmg.Board) bool {
	toBitboard := (uint64(1) << m.To())
	return (toBitboard&b.White.All != 0) || (toBitboard&b.Black.All != 0)
}
