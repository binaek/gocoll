package trie

import (
	"sync"
	"time"
)

// Tree represents a generic Trie (prefix Tree) structure.
type Tree[T any] struct {
	root     *node[T]
	syncSafe bool
	lock     *sync.RWMutex
}

// NewTree creates and returns a new non-thread-safe Tree instance.
func NewTree[T any]() Tree[T] {
	return Tree[T]{
		root:     newNode[T](),
		syncSafe: false,
		lock:     nil,
	}
}

// NewConcurrentTree creates and returns a new thread-safe Tree instance.
func NewConcurrentTree[T any]() Tree[T] {
	return Tree[T]{
		root:     newNode[T](),
		syncSafe: true,
		lock:     &sync.RWMutex{},
	}
}

// Insert adds a key-value pair to the Trie. It returns the old value (if any) and a boolean indicating if a value was replaced.
func (t *Tree[T]) Insert(key string, value T) (oldValue T, replaced bool) {
	return t.insert([]byte(key), value, nil)
}

// InsertWithExpiry adds a key-value pair to the Trie with an expiry duration. It returns the old value (if any) and a boolean indicating if a value was replaced.
func (t *Tree[T]) InsertWithExpiry(key string, value T, expiry time.Duration) (oldValue T, replaced bool) {
	expiryTime := time.Now().Add(expiry)
	return t.insert([]byte(key), value, &expiryTime)
}

// InsertB adds a key-value pair to the Trie using a byte slice key.
func (t *Tree[T]) InsertB(key []byte, value T) (oldValue T, replaced bool) {
	return t.insert(key, value, nil)
}

// InsertBWithExpiry adds a key-value pair to the Trie with an expiry duration using a byte slice key.
func (t *Tree[T]) InsertBWithExpiry(key []byte, value T, expiry time.Duration) (oldValue T, replaced bool) {
	expiryTime := time.Now().Add(expiry)
	return t.insert(key, value, &expiryTime)
}

// Find retrieves the value associated with the given key. It returns nil if the key does not exist or the value has expired.
func (t *Tree[T]) Find(key string) (value T, found bool) {
	if t.syncSafe {
		t.lock.RLock()
		defer t.lock.RUnlock()
	}
	zero := new(T)
	node := t.root
	for i := 0; i < len(key); i++ {
		if node.children[key[i]] == nil {
			return *zero, false
		}
		node = node.children[key[i]]
	}
	val, found := node.getValue()
	if found {
		return val, true
	}
	return *zero, false
}

// Remove deletes the key-value pair from the Trie. It returns the old value (if any) and a boolean indicating if a value was removed.
func (t *Tree[T]) Remove(key string) (oldValue T, removed bool) {
	if t.syncSafe {
		t.lock.Lock()
		defer t.lock.Unlock()
	}
	node := t.root
	for i := 0; i < len(key); i++ {
		if node.children[key[i]] == nil {
			return *new(T), false
		}
		node = node.children[key[i]]
	}
	if node.isEnd {
		oldValue, _ = node.getValue()
		node.isEnd = false
		node.value = nil
		removed = true
	} else {
		removed = false
	}
	return oldValue, removed
}

// insert adds a key-value pair to the Trie with an optional expiry time. It returns the old value (if any) and a boolean indicating if a value was replaced.
func (t *Tree[T]) insert(key []byte, value T, expiry *time.Time) (oldValue T, replaced bool) {
	if t.syncSafe {
		t.lock.Lock()
		defer t.lock.Unlock()
	}
	node := t.root
	for i := 0; i < len(key); i++ {
		if node.children[key[i]] == nil {
			node.children[key[i]] = newNode[T]()
		}
		node = node.children[key[i]]
	}
	if node.isEnd {
		oldValue, _ = node.getValue()
		replaced = true
	} else {
		replaced = false
	}
	node.setValue(value, expiry)
	return oldValue, replaced
}
