package collection

// Filter filters the elements of a collection based on a predicate.
func (collection *Collection[T]) Filter(f Predicate[T]) *Collection[T] {
	result := []T{}
	for _, e := range collection.Elements() {
		if f(e) {
			result = append(result, e)
		}
	}
	return From(result...)
}

// Partition partitions the elements of a collection based on a predicate.
func (collection *Collection[T]) Partition(f Predicate[T]) (*Collection[T], *Collection[T]) {
	leftPartition, rightPartition := []T{}, []T{}
	for _, e := range collection.Elements() {
		if f(e) {
			leftPartition = append(leftPartition, e)
			continue
		}
		rightPartition = append(rightPartition, e)
	}
	return From(leftPartition...), From(rightPartition...)
}

func (collection *Collection[T]) Distinct(e Equality[T]) *Collection[T] {
	result := []T{}
	for _, elem := range collection.Elements() {
		contains := false
		for _, res := range result {
			if e(elem, res) {
				contains = true
				break
			}
		}
		if !contains {
			result = append(result, elem)
		}
	}
	return From(result...)
}

// TakeWhile returns a collection with elements until the predicate is false.
func (collection *Collection[T]) TakeWhile(f Predicate[T]) *Collection[T] {
	result := []T{}
	for _, e := range collection.Elements() {
		fe := f(e)
		if fe {
			result = append(result, e)
			continue
		}
		break
	}
	return From(result...)
}

// DropWhile returns a collection without elements until the predicate is false.
func (collection *Collection[T]) DropWhile(f Predicate[T]) *Collection[T] {
	result := []T{}
	for _, e := range collection.Elements() {
		if f(e) {
			continue
		}
		result = append(result, e)
	}
	return From(result...)
}
