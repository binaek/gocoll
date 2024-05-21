package bloom

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"testing"
)

// TestBloomFilter tests the basic functionality of the BloomFilter.
func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter()

	tokens := []string{"apple", "banana", "grape", "orange"}
	for _, token := range tokens {
		bf.Add(token)
	}

	for _, token := range tokens {
		if !bf.Test(token) {
			t.Errorf("expected %s to be in the BloomFilter", token)
		}
	}

	if bf.Test("pineapple") {
		t.Errorf("expected pineapple not to be in the BloomFilter")
	}
}

// TestBloomFilterWithHashers tests the BloomFilter with custom hashers.
func TestBloomFilterWithHashers(t *testing.T) {
	hashers := []hash.Hash{
		sha256.New(),
		sha512.New(),
	}

	bf := NewBloomFilter(WithHashers(hashers))

	tokens := []string{"apple", "banana", "grape", "orange"}
	for _, token := range tokens {
		bf.Add(token)
	}

	for _, token := range tokens {
		if !bf.Test(token) {
			t.Errorf("expected %s to be in the BloomFilter", token)
		}
	}

	if bf.Test("pineapple") {
		t.Errorf("expected pineapple not to be in the BloomFilter")
	}
}

// BenchmarkBloomFilterAdd benchmarks the Add method of the BloomFilter.
func BenchmarkBloomFilterAdd(b *testing.B) {
	bf := NewBloomFilter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Add("benchmark-token")
	}
}

// BenchmarkBloomFilterTest benchmarks the Test method of the BloomFilter.
func BenchmarkBloomFilterTest(b *testing.B) {
	bf := NewBloomFilter()
	bf.Add("benchmark-token")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Test("benchmark-token")
	}
}

// BenchmarkBloomFilterAddWithCustomHashers benchmarks the Add method with custom hashers.
func BenchmarkBloomFilterAddWithCustomHashers(b *testing.B) {
	hashers := []hash.Hash{
		sha256.New(),
		sha512.New(),
	}

	bf := NewBloomFilter(WithHashers(hashers))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Add("benchmark-token")
	}
}

// BenchmarkBloomFilterTestWithCustomHashers benchmarks the Test method with custom hashers.
func BenchmarkBloomFilterTestWithCustomHashers(b *testing.B) {
	hashers := []hash.Hash{
		sha256.New(),
		sha512.New(),
	}

	bf := NewBloomFilter(WithHashers(hashers))
	bf.Add("benchmark-token")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Test("benchmark-token")
	}
}
