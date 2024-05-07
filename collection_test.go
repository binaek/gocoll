package gocoll

import "testing"

func TestAdd(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	if len(coll.Elements()) != 1 {
		t.Errorf("Expected length of 1, got %d", len(coll.Elements()))
	}
	if coll.Elements()[0] != 1 {
		t.Errorf("Expected element '1', got %d", coll.Elements()[0])
	}
}

func TestConcat(t *testing.T) {
	coll1 := New[int]()
	coll2 := New[int]()
	coll1.Add(1)
	coll2.Add(2)
	coll1.Concat(coll2)
	if len(coll1.Elements()) != 2 {
		t.Errorf("Expected length of 2, got %d", len(coll1.Elements()))
	}
}

func TestRemove(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	coll.Remove(func(x int) bool { return x == 1 })
	if len(coll.Elements()) != 1 || coll.Elements()[0] != 2 {
		t.Errorf("Failed to remove element")
	}
}
