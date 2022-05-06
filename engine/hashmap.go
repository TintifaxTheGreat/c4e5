package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type hashMap map[uint64]*Hash

func newHashMap() *hashMap {
	m := make(hashMap)
	return &m
}

func (h hashMap) put(depth int, value int, b *dragontoothmg.Board, m dragontoothmg.Move) {
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

func (h hashMap) get(depth int, b *dragontoothmg.Board) (int, dragontoothmg.Move, bool) {
	key := b.Hash()
	hash, ok := h[key]

	if ok {
		if hash.depth < depth {
			return 0, hash.move, false
		}
		return hash.value, hash.move, true
	}
	return 0, 0, false
}
