package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenFiles(t *testing.T) {
	g := NewGame("rnbqkbnr/p1ppp1p1/8/8/8/8/P1P1PPP1/RNBQKBNR w KQkq - 0 1", 0, 0)
	assert.Equal(t, uint64(0x8282828282828282), openFiles(&g.Board))
	assert.Equal(t, uint64(0x2828282828282828), halfOpenFiles(&g.Board))
}

func TestCountFiguresMoves(t *testing.T) {
	g := NewGame("8/8/4k3/3R4/2K5/8/8/8 w - - 0 1", 0, 0)
	assert.Equal(t, 7, countFiguresMoves(&g.Board, g.Board.White.Kings))
	assert.Equal(t, 0, countFiguresMoves(&g.Board, g.Board.Black.Kings))

	g = NewGame("8/8/4k3/3R4/2K5/8/8/8 b - - 0 1", 0, 0)
	assert.Equal(t, 0, countFiguresMoves(&g.Board, g.Board.White.Kings))
	assert.Equal(t, 3, countFiguresMoves(&g.Board, g.Board.Black.Kings))
}
