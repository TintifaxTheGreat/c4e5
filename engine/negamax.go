package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

func negamax(board *dragontoothmg.Board, hashmap *HashMap, depth int, alpha int, beta int) int {
	children := board.GenerateLegalMoves()
	if len(children) == 0 {
		if board.OurKingInCheck() == true {
			value := -40000 - depth
			hashmap.Put(depth, value, board)
			return value
		}
		hashmap.Put(depth, 0, board)
		return 0
	}

	if depth < 1 {
		value := evaluate(board)
		hashmap.Put(0, value, board)
		return value
	}

	pvs := true
	for _, child := range children {
		value := 0
		isCapture := testCapture(child, board)

		unapplyFunc := board.Apply(child)
		v, ok := hashmap.Get(depth, board)

		if ok {
			value = -v

			//a := alpha

			//value_calc := -negamax(board, hashmap, depth-1, -maxInt, -minInt)
			/*
				if value != value_calc {
					log.Print(board.ToFen(), " ", value, " ", value_calc, " ", depth-1)
				}

			*/
		} else {
			if (depth == 1) && isCapture {
				value = -quiescense(board, depth-1+initQuietDepth, -beta, -alpha)
			} else {
				if pvs {
					value = -negamax(board, hashmap, depth-1, -beta, -alpha)
					hashmap.Put(depth-1, value, board)
				} else {
					value = -negamax(board, hashmap, depth-1, -alpha-1, -alpha)
					if value > alpha {
						value = -negamax(board, hashmap, depth-1, -beta, -alpha)
						hashmap.Put(depth-1, value, board)
					}
				}

			}

		}
		unapplyFunc()

		if value >= beta {
			return beta
		}
		if value > alpha {
			alpha = value
			pvs = false
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
