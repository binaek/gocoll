package gocoll

import "testing"

func TestReduce(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	result := Reduce(*coll, 0, func(acc int, x int) int { return acc + x })
	if result != 3 {
		t.Errorf("Reduce failed, expected 3, got %d", result)
	}
}

func TestMin(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	if min := coll.Min(func(a, b int) bool { return a < b }); min != 1 {
		t.Errorf("Min did not return the smallest element, got %d", min)
	}
}

func TestMax(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	if max := coll.Max(func(a, b int) bool { return a < b }); max != 2 {
		t.Errorf("Max did not return the largest element, got %d", max)
	}
}

func TestGroupBy(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	coll.Add(3)
	groups := GroupBy(*coll, func(x int) int { return x % 2 })
	if len(groups) != 2 {
		t.Errorf("GroupBy failed, expected 2 groups, got %d", len(groups))
	}
	if len(groups[0].Elements()) != 1 {
		t.Errorf("GroupBy failed, expected 1 element in group 0, got %d", len(groups[0].Elements()))
	}
	if len(groups[1].Elements()) != 2 {
		t.Errorf("GroupBy failed, expected 2 elements in group 1, got %d", len(groups[1].Elements()))
	}
}
