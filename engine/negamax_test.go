package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestCapture(t *testing.T) {
	g := NewGame("4kN2/8/7K/b5B1/2N2R2/6rn/2K5/8 w - - 0 1", 0, 0)
	moves := g.Board.GenerateLegalMoves()
	var captures []dragontoothmg.Move
	for _, move := range moves {
		if testCapture(move, &g.Board) {
			captures = append(captures, move)
		}
	}
	assert.Equal(t, 1, len(captures))

	g = NewGame("4kN2/4P3/7K/b5B1/2N2R2/6rn/2K5/8 b - - 0 1", 0, 0)
	moves = g.Board.GenerateLegalMoves()
	captures = nil
	for _, move := range moves {
		if testCapture(move, &g.Board) {
			captures = append(captures, move)
		}
	}
	assert.Equal(t, 3, len(captures))

	g = NewGame("4kN2/5P2/7K/b5B1/2N2R2/6rn/2K5/8 b - - 0 1", 0, 0)
	moves = g.Board.GenerateLegalMoves()
	captures = nil
	for _, move := range moves {
		if testCapture(move, &g.Board) {
			captures = append(captures, move)
		}
	}
	assert.Equal(t, 1, len(captures))
}
