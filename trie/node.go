package trie

import (
	"time"
)

// valueWithExpiry represents a value stored in the Trie with an optional expiry time.
type valueWithExpiry[T any] struct {
	value  T
	expiry *time.Time
}

// node represents a node in the Trie.
type node[T any] struct {
	children map[byte]*node[T]
	isEnd    bool
	value    *valueWithExpiry[T]
}

// newNode creates and returns a new node instance.
func newNode[T any]() *node[T] {
	return &node[T]{children: make(map[byte]*node[T])} // we could preallocate the map with 256 to reduce allocations, but it's not worth it
}

// setValue sets the value and optional expiry time for a node, and marks the node as an end node.
func (n *node[T]) setValue(value T, expiry *time.Time) {
	n.isEnd = true
	n.value = &valueWithExpiry[T]{value: value, expiry: expiry}
}

// getValue returns the value of a node if it is an end node and not expired. It returns nil if the value is expired.
func (n *node[T]) getValue() (val T, notStale bool) {
	if n.isEnd {
		val := n.value
		if val.expiry != nil && val.expiry.Before(time.Now()) {
			return val.value, false
		}
		return val.value, true
	}
	t := new(T)
	return *t, false
}
