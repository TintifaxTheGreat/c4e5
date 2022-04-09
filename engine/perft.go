package engine

import "github.com/dylhunn/dragontoothmg"

func Perft(board *dragontoothmg.Board, depth int) int {
	count := 0
	moves := board.GenerateLegalMoves()

	if depth == 1 {
		return len(moves)
	}

	for _, move := range moves {
		var unapply func() = board.Apply(move)
		count += Perft(board, depth-1)
		unapply()
	}

	return count
}

func PerftPrime(board *dragontoothmg.Board, depth int) int {
	count := 0
	moves := generateMovesPrime(board)

	if depth == 1 {
		return len(moves)
	}

	for _, move := range moves {
		var unapply func() = board.Apply(move)
		count += PerftPrime(board, depth-1)
		unapply()
	}

	return count
}
