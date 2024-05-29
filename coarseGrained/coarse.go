package coarseGrained

import (
	"sync"
	"trees/simpleTree"
)

type bst struct {
	mu   sync.RWMutex
	tree simpleTree.Tree
}

func (b *bst) Insert(val int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.tree.Insert(val)
}

func (b *bst) Find(val int) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.tree.Find(val)
}

func (b *bst) Remove(val int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.tree.Remove(val)
}

func (b *bst) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.tree.String()
}
