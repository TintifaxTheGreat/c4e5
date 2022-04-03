package engine

type BoardHistory map[uint64]uint8

func NewBoardHistory() *BoardHistory {
	h := make(BoardHistory)
	return &h
}

func (g *Game) StoreBoardHistory() {
	key := g.Board.Hash()
	_, ok := g.BoardHistory[key]

	if ok {
		g.BoardHistory[key]++
	} else {
		g.BoardHistory[key] = 1
	}
}

func (g *Game) TestBoardHistory() uint8 {
	key := g.Board.Hash()
	v, ok := g.BoardHistory[key]

	if ok {
		return v
	}
	return 0
}
