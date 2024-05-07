package gocoll

// Find returns the first element that satisfies the predicate.
func (c Collection[T]) Find(f Finder[T]) (T, bool) {
	for _, e := range c.Elements() {
		if f(e) {
			return e, true
		}
	}
	var zeroValue T
	return zeroValue, false
}

// FindAll returns all elements that satisfy the predicate.
func (c Collection[T]) FindAll(f Finder[T]) *Collection[T] {
	result := New[T]()
	for _, e := range c.Elements() {
		if f(e) {
			result.Add(e)
		}
	}
	return result
}

// All returns true if all elements satisfy the predicate.
func (c Collection[T]) All(f Finder[T]) bool {
	for _, e := range c.Elements() {
		if !f(e) {
			return false
		}
	}
	return true
}

// Any returns true if any element satisfies the predicate.
func (c Collection[T]) Any(f Finder[T]) bool {
	for _, e := range c.Elements() {
		if f(e) {
			return true
		}
	}
	return false
}

// None returns true if no element satisfies the predicate.
func (c Collection[T]) None(f Finder[T]) bool {
	return !c.Any(f)
}

// FindIndex returns the index of the first element that satisfies the predicate.
func (c Collection[T]) FindIndex(f Finder[T]) int {
	for i, e := range c.Elements() {
		if f(e) {
			return i
		}
	}
	return -1
}
