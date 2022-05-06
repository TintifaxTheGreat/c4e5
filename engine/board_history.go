package engine

type boardHistory map[uint64]uint8

func newBoardHistory() *boardHistory {
	h := make(boardHistory)
	return &h
}

func (g *Game) StoreBoardHistory() {
	key := g.Board.Hash()
	_, ok := g.boardHistory[key]

	if ok {
		g.boardHistory[key]++
	} else {
		g.boardHistory[key] = 1
	}
}

func (g *Game) testBoardHistory() uint8 {
	key := g.Board.Hash()
	v, ok := g.boardHistory[key]

	if ok {
		return v
	}
	return 0
}
