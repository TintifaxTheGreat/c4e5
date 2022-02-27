package engine

import "github.com/dylhunn/dragontoothmg"

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

func OpenFiles(b *dragontoothmg.Board) uint64 {
	return ^fileFill(b.White.Pawns) & ^fileFill(b.Black.Pawns)
}
