package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"math/bits"
)

const CENTER uint64 = 0x00003C3C3C3C0000
const BORD_0 uint64 = 0xff818181818181ff
const BORD_1 uint64 = 0x007e424242427e00
const CENT_1 uint64 = 0x00003c24243c0000
const CENT_0 uint64 = 0x0000001818000000
const SAFE_KING uint64 = 0xc3000000000000c3
const GOOD_BISHOP uint64 = 0x42006666004200
const BASE_LINE uint64 = 0xff000000000000ff

func evaluate(b *dragontoothmg.Board) int {
	var value int = 0

	value += bits.OnesCount64(b.White.Pawns&CENT_0) * 4
	value -= bits.OnesCount64(b.Black.Pawns&CENT_0) * 4

	value += bits.OnesCount64(b.White.Pawns & CENT_1)
	value -= bits.OnesCount64(b.Black.Pawns & CENT_1)

	value += bits.OnesCount64(b.White.Pawns) * 20
	value -= bits.OnesCount64(b.Black.Pawns) * 20

	value += bits.OnesCount64(b.White.Knights) * 60
	value -= bits.OnesCount64(b.Black.Knights) * 60

	value += bits.OnesCount64(b.White.Bishops) * 61
	value -= bits.OnesCount64(b.Black.Bishops) * 61

	value += bits.OnesCount64(b.White.Rooks) * 100
	value -= bits.OnesCount64(b.Black.Rooks) * 100

	value += bits.OnesCount64(b.White.Queens) * 180
	value -= bits.OnesCount64(b.Black.Queens) * 180

	// TODO only use in early stage of game
	value += bits.OnesCount64(b.White.Knights & CENTER)
	value -= bits.OnesCount64(b.Black.Knights & CENTER)

	value -= bits.OnesCount64(b.White.Queens & CENTER)
	value += bits.OnesCount64(b.Black.Queens & CENTER)

	value -= bits.OnesCount64(b.White.Knights & BASE_LINE)
	value += bits.OnesCount64(b.Black.Knights & BASE_LINE)
	value -= bits.OnesCount64(b.White.Bishops & BASE_LINE)
	value += bits.OnesCount64(b.Black.Bishops & BASE_LINE)

	value += bits.OnesCount64(b.White.Kings&SAFE_KING) * 5
	value -= bits.OnesCount64(b.Black.Kings&SAFE_KING) * 5

	value += bits.OnesCount64(b.White.Bishops & GOOD_BISHOP)
	value -= bits.OnesCount64(b.Black.Bishops & GOOD_BISHOP)

	//bbDefendingKing := bbWhiteKing
	//if pos.Turn() == chess.Black {
	if b.Wtomove == false {
		value *= (-1)
		//bbDefendingKing = bbBlackKing
	}

	//TODO only use this in the endgame
	/*
		if value < 0 {
			value += distance(bbWhiteKing, bbBlackKing)

			value += bits.OnesCount64(bbDefendingKing&CENT_0) * 3
			value += bits.OnesCount64(bbDefendingKing&CENT_1) * 2
			value += bits.OnesCount64(bbDefendingKing&BORD_1) * 1

		} else {
			value -= bits.OnesCount64(bbDefendingKing&CENT_0) * 3
			value -= bits.OnesCount64(bbDefendingKing&CENT_1) * 2
			value -= bits.OnesCount64(bbDefendingKing&BORD_1) * 1
		}
	*/

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
