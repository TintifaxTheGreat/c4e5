package engine

import "github.com/dylhunn/dragontoothmg"

type Hash struct {
	depth int
	value int
}

type HashMap map[uint64]*Hash

func NewHashMap() *HashMap {
	m := make(HashMap)
	return &m
}
func (h HashMap) Put(depth int, value int, b *dragontoothmg.Board) {
	key := b.Hash()
	h[key] = &Hash{depth: depth, value: value}
}

func (h HashMap) Get(depth int, b *dragontoothmg.Board) (int, bool) {
	key := b.Hash()
	hash, ok := h[key]
	if ok && (hash.depth < depth) {
		ok = false
	}
	if ok {
		return hash.value, true
	}
	return 0, false
}
