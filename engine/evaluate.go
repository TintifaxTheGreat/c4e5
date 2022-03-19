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

//const cbGoodPawn uint64 = 0x1c1c000000
const cbBaseLine uint64 = 0xff000000000000ff

func evaluate(b *dragontoothmg.Board) int {
	var value int = 0
	piecesCount := bits.OnesCount64(b.White.All | b.Black.All)

	value += bits.OnesCount64(b.White.Pawns&cbCenter0) * 2
	value -= bits.OnesCount64(b.Black.Pawns&cbCenter0) * 2

	value += bits.OnesCount64(b.White.Pawns) * 20
	value -= bits.OnesCount64(b.Black.Pawns) * 20

	value += bits.OnesCount64(b.White.Knights) * 60
	value -= bits.OnesCount64(b.Black.Knights) * 60

	value += bits.OnesCount64(b.White.Bishops) * 62
	value -= bits.OnesCount64(b.Black.Bishops) * 62

	value += bits.OnesCount64(b.White.Rooks) * 95
	value -= bits.OnesCount64(b.Black.Rooks) * 95

	value += bits.OnesCount64(b.White.Queens) * 180
	value -= bits.OnesCount64(b.Black.Queens) * 180

	value += bits.OnesCount64(b.White.Rooks&OpenFiles(b)) * 2
	value -= bits.OnesCount64(b.Black.Rooks&OpenFiles(b)) * 2

	value += bits.OnesCount64(b.White.Rooks & HalfOpenFiles(b))
	value -= bits.OnesCount64(b.Black.Rooks & HalfOpenFiles(b))

	//TODO Rook on half open line

	value -= bits.OnesCount64(b.White.Knights&cbBoard0) * 2 //TODO
	value += bits.OnesCount64(b.Black.Knights&cbBoard0) * 2

	if piecesCount > 20 {
		value -= bits.OnesCount64(b.White.Queens&cbCenter) * 3
		value += bits.OnesCount64(b.Black.Queens&cbCenter) * 3

		value -= bits.OnesCount64(b.White.Bishops&cbBaseLine) * 2 //TODO
		value += bits.OnesCount64(b.Black.Bishops&cbBaseLine) * 2

		value -= bits.OnesCount64(b.White.Knights&cbBaseLine) * 2 //TODO
		value += bits.OnesCount64(b.Black.Knights&cbBaseLine) * 2

		value += bits.OnesCount64(b.White.Kings&cbSafeKing) * 10
		value -= bits.OnesCount64(b.Black.Kings&cbSafeKing) * 10

		value += bits.OnesCount64(b.White.Bishops&cbGoodBishop) * 2
		value -= bits.OnesCount64(b.Black.Bishops&cbGoodBishop) * 2
	}

	bbDefendingKing := b.White.Kings
	if b.Wtomove == false {
		value *= -1
		bbDefendingKing = b.Black.Kings
	}

	if piecesCount < 8 {
		if value < 0 {
			value += distance(b.White.Kings, b.Black.Kings) * 3

			value += bits.OnesCount64(bbDefendingKing&cbCenter0) * 3
			value += bits.OnesCount64(bbDefendingKing&cbCenter1) * 2
			value += bits.OnesCount64(bbDefendingKing&cbBoard1) * 1
			value -= bits.OnesCount64(bbDefendingKing&cbBoard0) * 5

		} else {
			value -= distance(b.White.Kings, b.Black.Kings) * 3

			value -= bits.OnesCount64(bbDefendingKing&cbCenter0) * 3
			value -= bits.OnesCount64(bbDefendingKing&cbCenter1) * 2
			value -= bits.OnesCount64(bbDefendingKing&cbBoard1) * 1
			value += bits.OnesCount64(bbDefendingKing&cbBoard0) * 5
		}
	}

	return value
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
