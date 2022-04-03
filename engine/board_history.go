package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type BoardHistory map[uint64]bool

func NewBoardHistory() *BoardHistory {
	h := make(BoardHistory)
	return &h
}

func (h BoardHistory) Store(b *dragontoothmg.Board) {
	key := b.Hash()
	_, ok := h[key]
	
	if !ok {
		h[key] = true
	}
}

func (h BoardHistory) Test(b *dragontoothmg.Board) bool {
	key := b.Hash()
	_, ok := h[key]
	return ok
}
