package gocoll

// Reduce applies a function to each element of a collection and returns a single value.
func Reduce[T any, U any](collection Collection[T], f Reducer[T, U], initial U) U {
	result := initial
	for _, e := range collection.Elements() {
		result = f(result, e)
	}
	return result
}

// Min returns the minimum element of a collection based on a comparator function.
func (collection *Collection[T]) Min(f func(T, T) bool) T {
	min := collection.elements[0]
	for _, e := range collection.Elements() {
		if f(e, min) {
			min = e
		}
	}
	return min
}

// Max returns the maximum element of a collection based on a comparator function.
func (collection *Collection[T]) Max(f func(T, T) bool) T {
	max := collection.elements[0]
	for _, e := range collection.Elements() {
		if f(max, e) {
			max = e
		}
	}
	return max
}

// GroupBy groups the elements of a collection based on a grouper function.
func GroupBy[T any, U comparable](collection Collection[T], f Grouper[T, U]) map[U]*Collection[T] {
	groups := make(map[U]*Collection[T])
	for _, e := range collection.Elements() {
		key := f(e)
		if _, ok := groups[key]; !ok {
			groups[key] = New[T]()
		}
		groups[key].Add(e)
	}
	return groups
}
