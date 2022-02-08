package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"log"
)

func negamax(board *dragontoothmg.Board, depth int, alpha int, beta int) int {
	children := board.GenerateLegalMoves()
	if len(children) == 0 {
		return -40000 - depth
		// TODO consider draws
	}

	if depth < 1 {
		return evaluate(board)
	}

	for _, child := range children {
		value := 0
		isCapture := testCapture(child, board)
		unapplyFunc := board.Apply(child)
		if (depth == 1) && isCapture {
			log.Println("entering q move ", child.String())
			value = -quiescense(board, depth-1+QDEPTH, -beta, -alpha)
		} else {
			value = -negamax(board, depth-1, -beta, -alpha)
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

func quiescense(board *dragontoothmg.Board, depth int, alpha int, beta int) int {
	children := board.GenerateLegalMoves()
	if len(children) == 0 {
		return -40000 - depth
		// TODO consider draws
	}

	if depth < 1 {
		return evaluate(board)
	}

	for _, child := range children {
		if !testCapture(child, board) {
			continue
		}
		value := 0
		unapplyFunc := board.Apply(child)
		value = -quiescense(board, depth-1, -beta, -alpha)
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
