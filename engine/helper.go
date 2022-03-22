package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"math/bits"
)

func northFill(p uint64) uint64 {
	p |= (p << 8)
	p |= (p << 16)
	p |= (p << 32)
	return p
}

func southFill(p uint64) uint64 {
	p |= (p >> 8)
	p |= (p >> 16)
	p |= (p >> 32)
	return p
}

func fileFill(p uint64) uint64 {
	return northFill(p) | southFill(p)
}

func openFiles(b *dragontoothmg.Board) uint64 {
	return ^fileFill(b.White.Pawns) & ^fileFill(b.Black.Pawns)
}

func halfOpenFiles(b *dragontoothmg.Board) uint64 {
	fw := fileFill(b.White.Pawns)
	fb := fileFill(b.Black.Pawns)
	return (fw & ^fb) | (^fw & fb)
}

func countFiguresMoves(b *dragontoothmg.Board, fig uint64) int {
	count := 0
	square := uint8(bits.TrailingZeros64(fig))

	moves := b.GenerateLegalMoves()
	for _, move := range moves {
		if move.From() == square {
			count++
		}
	}

	return count
}

func distance(x uint64, y uint64) int {
	xLz := bits.LeadingZeros64(x)
	yLz := bits.LeadingZeros64(y)
	fx := xLz % 8
	fy := yLz % 8
	rx := xLz / 8
	ry := yLz / 8

	fD := fy - fx
	if fD < 0 {
		fD = -fD
	}

	rD := ry - rx
	if rD < 0 {
		rD = -rD
	}

	if rD < fD {
		return fD
	}
	return rD
}
