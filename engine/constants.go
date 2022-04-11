package engine

import "time"

// Bitmaps
const (
	cbCenter     uint64 = 0x00003C3C3C3C0000
	cbBoard0     uint64 = 0xff818181818181ff
	cbBoard1     uint64 = 0x007e424242427e00
	cbCenter1    uint64 = 0x00003c24243c0000
	cbCenter0    uint64 = 0x0000001818000000
	cbSafeKing   uint64 = 0xc3000000000000c3
	cbGoodBishop uint64 = 0x42006666004200
	cbGoodQueen  uint64 = 0x3c1800000000183c
	cbBaseLine   uint64 = 0xff000000000000ff
)

// Search
const (
	maxInt           int = 1000000
	minInt           int = -1000000
	pruneThreshold   int = 40
	initMaxDepth         = 99
	initQuietDepth       = 1
	pvsDepth             = 2
	latePruningDepth     = 3
)

// Game
const (
	defaultTime = 10000 * time.Millisecond
)

// Evaluation
const (
	mateLevel = 55000
	mate      = 60000
)
