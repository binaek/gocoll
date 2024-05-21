package collection

import "testing"

func TestMap(t *testing.T) {
	coll := From(1, 2)
	result := Map(coll, func(x int) int { return (x * 2) })
	if len(result.Elements()) != 2 || result.Elements()[0] != 2 || result.Elements()[1] != 4 {
		t.Errorf("Map failed to apply function to elements")
	}
}

func TestFlatMap(t *testing.T) {
	coll := From(1, 2)
	result := FlatMap(coll, func(x int) *Collection[int] { return From(x, x) })
	if len(result.Elements()) != 4 || result.Elements()[0] != 1 || result.Elements()[1] != 1 || result.Elements()[2] != 2 || result.Elements()[3] != 2 {
		t.Errorf("FlatMap failed to apply function to elements")
	}
}
