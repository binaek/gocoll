package collection

import "testing"

func TestFilter(t *testing.T) {
	coll := New[int]()
	coll.Add(1)
	coll.Add(2)
	result := coll.Filter(func(x int) bool { return x == 1 })
	if len(result.Elements()) != 1 || result.Elements()[0] != 1 {
		t.Errorf("Filter failed to filter out elements")
	}
}

func TestPartition(t *testing.T) {
	coll := From(1, 2)
	matching, notMatching := coll.Partition(func(x int) bool { return x == 1 })
	if len(matching.Elements()) != 1 || matching.Elements()[0] != 1 {
		t.Errorf("Partition failed to partition matching elements")
	}
	if len(notMatching.Elements()) != 1 || notMatching.Elements()[0] != 2 {
		t.Errorf("Partition failed to partition not matching elements")
	}
}

func TestDistinct(t *testing.T) {
	coll := From(1, 2, 1)
	result := coll.Distinct(func(a, b int) bool { return a == b })
	if len(result.Elements()) != 2 || result.Elements()[0] != 1 || result.Elements()[1] != 2 {
		t.Errorf("Distinct failed to remove duplicate elements")
	}
}

func TestTakeWhile(t *testing.T) {
	coll := From(1, 2, 3)
	result := coll.TakeWhile(func(x int) bool { return x < 3 })
	if len(result.Elements()) != 2 || result.Elements()[0] != 1 || result.Elements()[1] != 2 {
		t.Errorf("TakeWhile failed to take elements while predicate was true")
	}
}

func TestDropWhile(t *testing.T) {
	coll := From(1, 2, 3)
	result := coll.DropWhile(func(x int) bool { return x < 3 })
	if len(result.Elements()) != 1 || result.Elements()[0] != 3 {
		t.Errorf("DropWhile failed to drop elements while predicate was true")
	}
}
