package engine

import (
	"github.com/dylhunn/dragontoothmg"
)

type Hash struct {
	depth int
	value int
	white bool
}

var cacheHit int
var cacheMiss int

type HashMap map[uint64]*Hash

func NewHashMap() *HashMap {
	m := make(HashMap)
	return &m
}
func (h HashMap) Put(depth int, value int, b *dragontoothmg.Board) {
	key := b.Hash()
	hash, ok := h[key]

	if !ok || (hash.depth <= depth) {
		h[key] = &Hash{
			depth: depth,
			value: value,
		}
	}
}

func (h HashMap) Get(depth int, b *dragontoothmg.Board) (int, bool) {
	key := b.Hash()
	hash, ok := h[key]

	if ok && (hash.depth < depth) {
		ok = false
	}
	if ok {
		cacheHit++
		v := hash.value
		return v, true
	}
	cacheMiss++
	return 0, false
}
