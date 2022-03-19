package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Hash struct {
	depth int
	value int
	move  dragontoothmg.Move
}

var cacheHit int
var cacheMiss int

type HashMap map[uint64]*Hash

func NewHashMap() *HashMap {
	m := make(HashMap)
	return &m
}

// TODO consider hash entries w. higher depth and the effect of pruneWorseIndex

func (h HashMap) Put(depth int, value int, b *dragontoothmg.Board, m dragontoothmg.Move) {
	key := b.Hash()
	hash, ok := h[key]

	if !ok || (hash.depth <= depth) {
		h[key] = &Hash{
			depth: depth,
			value: value,
			move:  m,
		}
	}
}

func (h HashMap) Get(depth int, b *dragontoothmg.Board) (int, dragontoothmg.Move, bool) {
	key := b.Hash()
	hash, ok := h[key]

	if ok {
		if hash.depth < depth {
			cacheMiss++
			return 0, hash.move, false
		}
		cacheHit++
		return hash.value, hash.move, true
	}
	cacheMiss++
	return 0, 0, false
}
