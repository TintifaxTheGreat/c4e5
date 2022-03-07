package engine

import (
	"github.com/dylhunn/dragontoothmg"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestGet(t *testing.T) {
	alpha := minInt
	beta := maxInt
	depth := 4

	g := NewGame("4kN2/8/7K/b5B1/2N2R2/6rn/2K5/8 w - - 0 1")
	values := make(map[dragontoothmg.Move]int)
	moves := g.Board.GenerateLegalMoves()
	for _, move := range moves {
		unapplyFunc := g.Board.Apply(move)
		values[move] = -negamax(&g.Board, g.HashMap, depth, -beta, -alpha)
		g.HashMap.Put(depth, values[move], &g.Board)
		unapplyFunc()
	}

	sortedMoves := make([]dragontoothmg.Move, 0, len(values))
	for key := range values {
		sortedMoves = append(sortedMoves, key)
	}
	sort.Slice(sortedMoves, func(i, j int) bool {
		return values[sortedMoves[i]] > values[sortedMoves[j]]
	})

	for _, move := range sortedMoves {
		unapplyFunc := g.Board.Apply(move)
		storedValue, _ := g.HashMap.Get(depth, &g.Board)
		storedValue1, _ := g.HashMap.Get(depth-1, &g.Board)
		_, ok2 := g.HashMap.Get(depth+1, &g.Board)

		calculatedValue := -negamax(&g.Board, g.HashMap, depth, -beta, -alpha)

		assert.Equal(t, storedValue, calculatedValue)
		assert.Equal(t, storedValue1, calculatedValue)
		assert.Equal(t, false, ok2)

		unapplyFunc()
	}

	g = NewGame("4kN2/8/7K/b5B1/2N2R2/6rn/2K5/8 b - - 0 1")
	values = make(map[dragontoothmg.Move]int)
	moves = g.Board.GenerateLegalMoves()
	for _, move := range moves {
		unapplyFunc := g.Board.Apply(move)
		values[move] = -negamax(&g.Board, g.HashMap, depth, -beta, -alpha)
		g.HashMap.Put(depth, values[move], &g.Board)
		unapplyFunc()
	}

	sortedMoves = make([]dragontoothmg.Move, 0, len(values))
	for key := range values {
		sortedMoves = append(sortedMoves, key)
	}
	sort.Slice(sortedMoves, func(i, j int) bool {
		return values[sortedMoves[i]] > values[sortedMoves[j]]
	})

	for _, move := range sortedMoves {
		unapplyFunc := g.Board.Apply(move)
		storedValue, _ := g.HashMap.Get(depth, &g.Board)
		storedValue1, _ := g.HashMap.Get(depth-1, &g.Board)
		_, ok2 := g.HashMap.Get(depth+1, &g.Board)

		calculatedValue := -negamax(&g.Board, g.HashMap, depth, -beta, -alpha)

		assert.Equal(t, storedValue, calculatedValue)
		assert.Equal(t, storedValue1, calculatedValue)
		assert.Equal(t, false, ok2)

		unapplyFunc()
	}
}
