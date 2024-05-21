package trie

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertAndFind(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value
	oldValue, replaced := trie.Insert("hello", "world")
	if replaced || oldValue != "" {
		t.Errorf("expected replaced=false, oldValue='', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Find the inserted value
	value, found := trie.Find("hello")
	if !found || value != "world" {
		t.Errorf("expected found=true, value='world', got found=%v, value=%v", found, value)
	}

	// Insert a new value for the same key
	oldValue, replaced = trie.Insert("hello", "universe")
	if !replaced || oldValue != "world" {
		t.Errorf("expected replaced=true, oldValue='world', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Find the new value
	value, found = trie.Find("hello")
	if !found || value != "universe" {
		t.Errorf("expected found=true, value='universe', got found=%v, value=%v", found, value)
	}
}

func TestInsertWithExpiry(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value with an expiry of 1 second
	oldValue, replaced := trie.InsertWithExpiry("temp", "data", 1*time.Second)
	if replaced || oldValue != "" {
		t.Errorf("expected replaced=false, oldValue='', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Find the inserted value
	value, found := trie.Find("temp")
	if !found || value != "data" {
		t.Errorf("expected found=true, value='data', got found=%v, value=%v", found, value)
	}

	// Wait for expiry
	time.Sleep(2 * time.Second)

	// Find the value again
	value, found = trie.Find("temp")
	if found || value != "" {
		t.Errorf("expected found=false, value='', got found=%v, value=%v", found, value)
	}
}

func TestRemove(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value
	trie.Insert("hello", "world")

	// Remove the value
	oldValue, removed := trie.Remove("hello")
	if !removed || oldValue != "world" {
		t.Errorf("expected removed=true, oldValue='world', got removed=%v, oldValue=%v", removed, oldValue)
	}

	// Try to find the removed value
	value, found := trie.Find("hello")
	if found || value != "" {
		t.Errorf("expected found=false, value='', got found=%v, value=%v", found, value)
	}

	// Try to remove a non-existent key
	oldValue, removed = trie.Remove("nonexistent")
	if removed || oldValue != "" {
		t.Errorf("expected removed=false, oldValue='', got removed=%v, oldValue=%v", removed, oldValue)
	}
}

func TestExpiryDoesNotAffectOtherKeys(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value without expiry
	trie.Insert("permanent", "value")

	// Insert a value with expiry
	trie.InsertWithExpiry("temporary", "data", 1*time.Second)

	// Wait for expiry
	time.Sleep(2 * time.Second)

	// Find the permanent value
	value, found := trie.Find("permanent")
	if !found || value != "value" {
		t.Errorf("expected found=true, value='value', got found=%v, value=%v", found, value)
	}
}

func TestEmptyTreeFind(t *testing.T) {
	trie := NewTree[string]()

	// Try to find a value in an empty tree
	value, found := trie.Find("hello")
	if found || value != "" {
		t.Errorf("expected found=false, value='', got found=%v, value=%v", found, value)
	}
}

func TestOverwriteWithExpiry(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value without expiry
	oldValue, replaced := trie.Insert("key", "initial")
	if replaced || oldValue != "" {
		t.Errorf("expected replaced=false, oldValue='', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Insert a new value with expiry
	oldValue, replaced = trie.InsertWithExpiry("key", "updated", 1*time.Second)
	if !replaced || oldValue != "initial" {
		t.Errorf("expected replaced=true, oldValue='initial', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Find the updated value
	value, found := trie.Find("key")
	if !found || value != "updated" {
		t.Errorf("expected found=true, value='updated', got found=%v, value=%v", found, value)
	}

	// Wait for expiry
	time.Sleep(2 * time.Second)

	// Find the value after expiry
	value, found = trie.Find("key")
	if found || value != "" {
		t.Errorf("expected found=false, value='', got found=%v, value=%v", found, value)
	}
}

func TestPartialKeyInsertAndFind(t *testing.T) {
	trie := NewTree[string]()

	// Insert values with partial keys
	trie.Insert("app", "partial")
	trie.Insert("apple", "complete")

	// Find the value for the complete key
	value, found := trie.Find("apple")
	if !found || value != "complete" {
		t.Errorf("expected found=true, value='complete', got found=%v, value=%v", found, value)
	}

	// Find the value for the partial key
	value, found = trie.Find("app")
	if !found || value != "partial" {
		t.Errorf("expected found=true, value='partial', got found=%v, value=%v", found, value)
	}
}

func TestInsertAndRemoveMultipleValues(t *testing.T) {
	trie := NewTree[int]()

	// Insert multiple values
	keys := []string{"one", "two", "three", "four", "five"}
	for i, key := range keys {
		trie.Insert(key, i+1)
	}

	// Find and check each value
	for i, key := range keys {
		value, found := trie.Find(key)
		if !found || value != i+1 {
			t.Errorf("expected found=true, value=%d, got found=%v, value=%v", i+1, found, value)
		}
	}

	// Remove and check each value
	for i, key := range keys {
		oldValue, removed := trie.Remove(key)
		if !removed || oldValue != i+1 {
			t.Errorf("expected removed=true, oldValue=%d, got removed=%v, oldValue=%v", i+1, removed, oldValue)
		}
		value, found := trie.Find(key)
		if found || value != 0 {
			t.Errorf("expected found=false, value=0, got found=%v, value=%v", found, value)
		}
	}
}

func TestOverwriteWithoutExpiry(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value with expiry
	oldValue, replaced := trie.InsertWithExpiry("key", "initial", 1*time.Second)
	if replaced || oldValue != "" {
		t.Errorf("expected replaced=false, oldValue='', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Wait for expiry
	time.Sleep(2 * time.Second)

	// Overwrite the expired value without expiry
	oldValue, replaced = trie.Insert("key", "updated")
	if !replaced || oldValue != "initial" {
		t.Errorf("expected replaced=true, oldValue='initial', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Find the updated value
	value, found := trie.Find("key")
	if !found || value != "updated" {
		t.Errorf("expected found=true, value='updated', got found=%v, value=%v", found, value)
	}
}

func TestExpiredValueDoesNotReplaceNonExpiredValue(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value with expiry
	oldValue, replaced := trie.InsertWithExpiry("key", "initial", 1*time.Second)
	if replaced || oldValue != "" {
		t.Errorf("expected replaced=false, oldValue='', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Insert a non-expiring value before expiry
	oldValue, replaced = trie.Insert("key", "updated")
	if !replaced || oldValue != "initial" {
		t.Errorf("expected replaced=true, oldValue='initial', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Wait for the original expiry to pass
	time.Sleep(2 * time.Second)

	// Ensure the non-expiring value is still present
	value, found := trie.Find("key")
	if !found || value != "updated" {
		t.Errorf("expected found=true, value='updated', got found=%v, value=%v", found, value)
	}
}

func TestConcurrency(t *testing.T) {
	trie := NewConcurrentTree[int]() // Create a thread-safe Trie

	done := make(chan bool)

	// Concurrent inserts
	go func() {
		for i := 0; i < 1000; i++ {
			trie.Insert(fmt.Sprint(i), i)
		}
		done <- true
	}()

	// Concurrent finds
	go func() {
		for i := 0; i < 1000; i++ {
			trie.Find(fmt.Sprint(i))
		}
		done <- true
	}()

	// Concurrent removals
	go func() {
		for i := 0; i < 1000; i++ {
			trie.Remove(fmt.Sprint(i))
		}
		done <- true
	}()

	<-done
	<-done
	<-done
}

func TestNonExpiredValueDoesNotOverrideExistingValue(t *testing.T) {
	trie := NewTree[string]()

	// Insert a value without expiry
	trie.Insert("key", "value")

	// Insert a value with expiry for the same key
	oldValue, replaced := trie.InsertWithExpiry("key", "newvalue", 1*time.Second)
	if !replaced || oldValue != "value" {
		t.Errorf("expected replaced=true, oldValue='value', got replaced=%v, oldValue=%v", replaced, oldValue)
	}

	// Wait for the expiry to pass
	time.Sleep(2 * time.Second)

	// Ensure the key no longer exists
	value, found := trie.Find("key")
	if found || value != "" {
		t.Errorf("expected found=false, value='', got found=%v, value=%v", found, value)
	}
}

func BenchmarkInsert(b *testing.B) {
	trie := NewTree[string]()
	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}
}

func BenchmarkInsertWithExpiry(b *testing.B) {
	trie := NewTree[string]()
	for n := 0; n < b.N; n++ {
		trie.InsertWithExpiry(fmt.Sprint(n), "value", 10*time.Second)
	}
}

func BenchmarkFind(b *testing.B) {
	trie := NewTree[string]()
	for n := 0; n < 1000; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Find(fmt.Sprint(n % 1000))
	}
}

func BenchmarkRemove(b *testing.B) {
	trie := NewTree[string]()
	for n := 0; n < 1000; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Remove(fmt.Sprint(n % 1000))
	}
}

func BenchmarkInsertThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}
}

func BenchmarkInsertWithExpiryThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	for n := 0; n < b.N; n++ {
		trie.InsertWithExpiry(fmt.Sprint(n), "value", 10*time.Second)
	}
}

func BenchmarkFindThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	for n := 0; n < 1000; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Find(fmt.Sprint(n % 1000))
	}
}

func BenchmarkRemoveThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	for n := 0; n < 1000; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Remove(fmt.Sprint(n % 1000))
	}
}

func BenchmarkHeapAllocations(b *testing.B) {
	trie := NewTree[string]()
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprint(n), "value")
		trie.Find(fmt.Sprint(n))
		trie.Remove(fmt.Sprint(n))
	}
}

func BenchmarkHeapAllocationsThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprint(n), "value")
		trie.Find(fmt.Sprint(n))
		trie.Remove(fmt.Sprint(n))
	}
}

// Additional benchmarks

func BenchmarkLargeInsert(b *testing.B) {
	trie := NewTree[string]()
	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprintf("key-%d", n), "value")
	}
}

func BenchmarkLargeInsertThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprintf("key-%d", n), "value")
	}
}

func BenchmarkMixedOperations(b *testing.B) {
	trie := NewTree[string]()
	for n := 0; n < 1000; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprint(n%1000), "newvalue")
		trie.Find(fmt.Sprint(n % 1000))
		trie.Remove(fmt.Sprint(n % 1000))
	}
}

func BenchmarkMixedOperationsThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	for n := 0; n < 1000; n++ {
		trie.Insert(fmt.Sprint(n), "value")
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Insert(fmt.Sprint(n%1000), "newvalue")
		trie.Find(fmt.Sprint(n % 1000))
		trie.Remove(fmt.Sprint(n % 1000))
	}
}

func BenchmarkDeepTrieInsert(b *testing.B) {
	trie := NewTree[string]()
	key := "a"
	for n := 0; n < b.N; n++ {
		trie.Insert(key, "value")
		key += "a"
	}
}

func BenchmarkDeepTrieInsertThreadSafe(b *testing.B) {
	trie := NewConcurrentTree[string]()
	key := "a"
	for n := 0; n < b.N; n++ {
		trie.Insert(key, "value")
		key += "a"
	}
}
