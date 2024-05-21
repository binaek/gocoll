package collection

import "testing"

func TestFind(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	found, ok := coll.Find(func(x int) bool { return x == 2 })
	if !ok || found != 2 {
		t.Errorf("Find failed to find the element")
	}
	_, ok = coll.Find(func(x int) bool { return x == 3 })
	if ok {
		t.Errorf("Find should not have succeeded")
	}
}

func TestFindAll(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	coll.Add(1)
	result := coll.FindAll(func(x int) bool { return x == 1 })
	if len(result.Elements()) != 2 {
		t.Errorf("FindAll failed to find all matching elements")
	}
}

func TestAll(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(1)
	if !coll.All(func(x int) bool { return x == 1 }) {
		t.Errorf("All reported false when all elements matched")
	}
	coll.Add(2)
	if coll.All(func(x int) bool { return x == 1 }) {
		t.Errorf("All reported true when not all elements matched")
	}
}

func TestAny(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	if !coll.Any(func(x int) bool { return x == 2 }) {
		t.Errorf("Any reported false when there was a matching element")
	}
	if coll.Any(func(x int) bool { return x == 3 }) {
		t.Errorf("Any reported true when there was no matching element")
	}
}

func TestNone(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	if !coll.None(func(x int) bool { return x == 3 }) {
		t.Errorf("None reported false when there were no matching elements")
	}
}

func TestFindIndex(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	idx := coll.FindIndex(func(x int) bool { return x == 2 })
	if idx != 1 {
		t.Errorf("FindIndex returned incorrect index, got %d", idx)
	}
}
