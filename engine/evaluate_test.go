package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeMove(t *testing.T) {
	g := new(Game)
	g.Board = dragontoothmg.ParseFen("r1b2k1r/pppq3p/2np1p2/8/2B2B2/8/PPP3PP/4RR1K w - - 0 1")
	move := g.FindMove()
	assert.Equal(t, "f4h6", move.String())

	g = new(Game)
	g.Board = dragontoothmg.ParseFen("1rb4r/pkPp3p/1b1P3n/1Q6/N3Pp2/8/P1P3PP/7K w - - 1 1")
	move = g.FindMove()
	assert.Equal(t, "b5d5", move.String())

	g = new(Game)
	g.Board = dragontoothmg.ParseFen("8/2Q5/8/6q1/2K5/8/8/7k b - - 0 1")
	move = g.FindMove()
	assert.Equal(t, "g5c1", move.String())

	g = new(Game)
	g.Board = dragontoothmg.ParseFen("4r1k1/5bpp/2p5/3pr3/8/1B3pPq/PPR2P2/2R2QK1 b - - 0 1")
	move = g.FindMove()
	assert.Equal(t, "e5e1", move.String())

	g = new(Game)
	g.Board = dragontoothmg.ParseFen("2b3rk/1q3p1p/p1p1pPpQ/4N3/2pP4/2P1p1P1/1P4PK/5R2 w - - 1 1")
	move = g.FindMove()
	assert.Equal(t, "f1h1", move.String())
}