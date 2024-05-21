package bloom

import (
	"fmt"
	"hash"
	"hash/fnv"
	"sync"

	"github.com/binaek/gocoll/trie"
)

type BloomConfig func(*BloomFilter)

func WithHashers(hashers []hash.Hash) BloomConfig {
	return func(bf *BloomFilter) {
		bf.hashers = hashers
	}
}

// BloomFilter represents a Bloom Filter data structure.
type BloomFilter struct {
	trie    trie.Tree[struct{}]
	hashers []hash.Hash
	lock    sync.RWMutex
}

// NewBloomFilter creates and returns a new BloomFilter instance.
func NewBloomFilter(config ...BloomConfig) *BloomFilter {
	bf := &BloomFilter{
		trie:    trie.NewConcurrentTree[struct{}](),
		hashers: []hash.Hash{fnv.New64(), fnv.New64a()},
	}
	for _, c := range config {
		c(bf)
	}
	return bf
}

// Add inserts a token into the Bloom Filter.
func (bf *BloomFilter) Add(token string) {
	hashes := bf.getHashes(token)
	for _, hash := range hashes {
		bf.trie.InsertB(hash, struct{}{} /* add an empty struct */)
	}
}

// Test checks if a token is possibly in the Bloom Filter.
func (bf *BloomFilter) Test(token string) bool {
	hashes := bf.getHashes(token)
	for _, hash := range hashes {
		_, found := bf.trie.Find(string(hash))
		if found {
			return true
		}
	}
	return false
}

func (bf *BloomFilter) getHashes(token string) [][]byte {
	bf.lock.Lock()
	defer bf.lock.Unlock()

	hashes := make([][]byte, len(bf.hashers))
	for i, hasher := range bf.hashers {
		hasher.Reset()
		hasher.Write([]byte(token))
		// add a salt to reduce the chance of hash collision
		hasher.Write([]byte(fmt.Sprint(i)))
		hashes[i] = hasher.Sum(nil)
	}
	return hashes
}
