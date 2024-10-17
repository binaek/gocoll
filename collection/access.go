package collection

// Find returns the first element that satisfies the predicate, along with a boolean indicating whether an element was found.
// If no element is found, the zero value of the element type is returned with a false boolean.
func (c Collection[T]) Find(f Predicate[T]) (T, bool) {
	for _, e := range c.Elements() {
		if f(e) {
			return e, true
		}
	}
	return ZeroValue[T](), false
}

// FindAll returns all elements that satisfy the predicate as a new Collection.
func (c *Collection[T]) FindAll(f Predicate[T]) *Collection[T] {
	result := []T{}
	for _, e := range c.Elements() {
		if f(e) {
			result = append(result, e)
		}
	}
	return From(result...)
}

// All returns true if all elements satisfy the predicate.
func (c *Collection[T]) All(f Predicate[T]) bool {
	for _, e := range c.Elements() {
		if !f(e) {
			return false
		}
	}
	return true
}

// Any returns true if any element satisfies the predicate.
func (c *Collection[T]) Any(f Predicate[T]) bool {
	for _, e := range c.Elements() {
		if f(e) {
			return true
		}
	}
	return false
}

// None returns true if no element satisfies the predicate.
func (c Collection[T]) None(f Predicate[T]) bool {
	return !c.Any(f)
}

// FindIndex returns the index of the first element that satisfies the predicate or -1 if no element matches.
func (c Collection[T]) FindIndex(f Predicate[T]) int {
	for i, e := range c.Elements() {
		if f(e) {
			return i
		}
	}
	return -1
}
