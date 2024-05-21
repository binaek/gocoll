package collection

// Filter filters the elements of a collection based on a predicate.
func (collection *Collection[T]) Filter(f Predicate[T]) *Collection[T] {
	result := New[T]()
	for _, e := range collection.Elements() {
		if f(e) {
			result.Add(e)
		}
	}
	return result
}

// Partition partitions the elements of a collection based on a predicate.
func (collection *Collection[T]) Partition(f Predicate[T]) (*Collection[T], *Collection[T]) {
	matching := New[T]()
	notMatching := New[T]()
	for _, e := range collection.Elements() {
		if f(e) {
			matching.Add(e)
		} else {
			notMatching.Add(e)
		}
	}
	return matching, notMatching
}

func (collection *Collection[T]) Distinct(e Equality[T]) *Collection[T] {
	result := New[T]()
	for _, elem := range collection.Elements() {
		contains := false
		for _, res := range result.Elements() {
			if e(elem, res) {
				contains = true
				break
			}
		}
		if !contains {
			result.Add(elem)
		}
	}
	return result
}

// TakeWhile returns a collection with elements until the predicate is false.
func (collection *Collection[T]) TakeWhile(f Predicate[T]) *Collection[T] {
	result := New[T]()
	for _, e := range collection.Elements() {
		if f(e) {
			result.Add(e)
		} else {
			break
		}
	}
	return result
}

// DropWhile returns a collection without elements until the predicate is false.
func (collection *Collection[T]) DropWhile(f Predicate[T]) *Collection[T] {
	result := New[T]()
	shouldAdd := false
	for _, e := range collection.Elements() {
		if !shouldAdd && f(e) {
			continue
		}
		shouldAdd = true
		result.Add(e)
	}
	return result
}
