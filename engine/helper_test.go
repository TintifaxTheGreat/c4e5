package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenFiles(t *testing.T) {
	g := NewGame("rnbqkbnr/p1ppp1p1/8/8/8/8/P1P1PPP1/RNBQKBNR w KQkq - 0 1")
	assert.Equal(t, uint64(0x8282828282828282), OpenFiles(&g.Board))
	assert.Equal(t, uint64(0x2828282828282828), HalfOpenFiles(&g.Board))
}
