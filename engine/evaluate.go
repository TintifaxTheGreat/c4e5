package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"math/bits"
)

const cbCenter uint64 = 0x00003C3C3C3C0000
const cbBoard0 uint64 = 0xff818181818181ff
const cbBoard1 uint64 = 0x007e424242427e00
const cbCenter1 uint64 = 0x00003c24243c0000
const cbCenter0 uint64 = 0x0000001818000000
const cbSafeKing uint64 = 0xc3000000000000c3
const cbGoodBishop uint64 = 0x42006666004200
const cbGoodQueen uint64 = 0x3c1800000000183c

//const cbGoodPawn uint64 = 0x1c1c000000
const cbBaseLine uint64 = 0xff000000000000ff

func evaluate(b *dragontoothmg.Board, depth int) int {
	var value int = 0
	movesCount := b.Fullmoveno
	piecesCount := bits.OnesCount64(b.White.All | b.Black.All)

	value += bits.OnesCount64(b.White.Pawns&cbCenter0) * 30
	value -= bits.OnesCount64(b.Black.Pawns&cbCenter0) * 30

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

	value += bits.OnesCount64(b.White.Rooks&openFiles(b)) * 20
	value -= bits.OnesCount64(b.Black.Rooks&openFiles(b)) * 20

	value += bits.OnesCount64(b.White.Rooks&halfOpenFiles(b)) * 10
	value -= bits.OnesCount64(b.Black.Rooks&halfOpenFiles(b)) * 10

	value -= bits.OnesCount64(b.White.Knights&cbBoard0) * 20
	value += bits.OnesCount64(b.Black.Knights&cbBoard0) * 20

	if movesCount < 12 {
		value += bits.OnesCount64(b.White.Queens&cbGoodQueen) * 120 // TODO was 8
		value -= bits.OnesCount64(b.Black.Queens&cbGoodQueen) * 120

	}

	if piecesCount > 20 {
		value -= bits.OnesCount64(b.White.Queens&cbCenter) * 30
		value += bits.OnesCount64(b.Black.Queens&cbCenter) * 30

		value -= bits.OnesCount64(b.White.Bishops&cbBaseLine) * 20
		value += bits.OnesCount64(b.Black.Bishops&cbBaseLine) * 20

		value -= bits.OnesCount64(b.White.Knights&cbBaseLine) * 20
		value += bits.OnesCount64(b.Black.Knights&cbBaseLine) * 20

		value += bits.OnesCount64(b.White.Kings&cbSafeKing) * 100
		value -= bits.OnesCount64(b.Black.Kings&cbSafeKing) * 100

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
