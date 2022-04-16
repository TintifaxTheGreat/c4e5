package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"math/bits" // TODO think if this is wise
)

func evaluate(b *dragontoothmg.Board) int {
	var value int = 0
	movesCount := b.Fullmoveno
	piecesCount := bits.OnesCount64(b.White.All | b.Black.All)

	bOpenFiles := openFiles(b)
	bHalfOpenFiles := halfOpenFiles(b)

	value += bits.OnesCount64(b.White.Pawns&cbCenter0) * 40
	value -= bits.OnesCount64(b.Black.Pawns&cbCenter0) * 40

	value += bits.OnesCount64(b.White.Pawns) * 190
	value -= bits.OnesCount64(b.Black.Pawns) * 200

	value += bits.OnesCount64(b.White.Knights) * 590
	value -= bits.OnesCount64(b.Black.Knights) * 600

	value += bits.OnesCount64(b.White.Bishops) * 610
	value -= bits.OnesCount64(b.Black.Bishops) * 620

	value += bits.OnesCount64(b.White.Rooks) * 940
	value -= bits.OnesCount64(b.Black.Rooks) * 950

	value += bits.OnesCount64(b.White.Queens) * 1790
	value -= bits.OnesCount64(b.Black.Queens) * 1800

	value += bits.OnesCount64(b.White.Rooks&bOpenFiles) * 20
	value -= bits.OnesCount64(b.Black.Rooks&bOpenFiles) * 20

	value += bits.OnesCount64(b.White.Rooks&bHalfOpenFiles) * 10
	value -= bits.OnesCount64(b.Black.Rooks&bHalfOpenFiles) * 10

	value -= bits.OnesCount64(b.White.Knights&cbBoard0) * 20
	value += bits.OnesCount64(b.Black.Knights&cbBoard0) * 20

	if movesCount < 12 {
		value += bits.OnesCount64(b.White.Queens&cbGoodQueen) * 120
		value -= bits.OnesCount64(b.Black.Queens&cbGoodQueen) * 120
	}

	if piecesCount > 20 {
		value -= bits.OnesCount64(b.White.Queens&cbCenter) * 40
		value += bits.OnesCount64(b.Black.Queens&cbCenter) * 40

		value -= bits.OnesCount64(b.White.Bishops&cbBaseLine) * 30
		value += bits.OnesCount64(b.Black.Bishops&cbBaseLine) * 30

		value -= bits.OnesCount64(b.White.Knights&cbBaseLine) * 30
		value += bits.OnesCount64(b.Black.Knights&cbBaseLine) * 30

		value += bits.OnesCount64(b.White.Kings&cbSafeKing) * 130
		value -= bits.OnesCount64(b.Black.Kings&cbSafeKing) * 130

		value += bits.OnesCount64(b.White.Bishops&cbGoodBishop) * 20
		value -= bits.OnesCount64(b.Black.Bishops&cbGoodBishop) * 20
	}

	bbDefendingKing := b.White.Kings
	if b.Wtomove == false {
		value *= -1
		bbDefendingKing = b.Black.Kings
	}

	if piecesCount < 8 {
		if value < 0 {
			value += distance(b.White.Kings, b.Black.Kings) * 20
			value += countFiguresMoves(b, bbDefendingKing) * 10
			value += bits.OnesCount64(bbDefendingKing&cbCenter0) * 80
			value += bits.OnesCount64(bbDefendingKing&cbCenter1) * 40
			value += bits.OnesCount64(bbDefendingKing&cbBoard1) * 10
			value -= bits.OnesCount64(bbDefendingKing&cbBoard0) * 50
		} else {
			value -= distance(b.White.Kings, b.Black.Kings) * 20
			value -= countFiguresMoves(b, bbDefendingKing) * 10
			value -= bits.OnesCount64(bbDefendingKing&cbCenter0) * 80
			value -= bits.OnesCount64(bbDefendingKing&cbCenter1) * 40
			value -= bits.OnesCount64(bbDefendingKing&cbBoard1) * 10
			value += bits.OnesCount64(bbDefendingKing&cbBoard0) * 50
		}
	}

	return value
}
