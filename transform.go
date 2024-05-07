package gocoll

// Map applies a function to each element of a collection and returns a new collection with the results.
func Map[T, U any](collection *Collection[T], f Mapper[T, U]) *Collection[U] {
	result := New[U]()
	for _, e := range collection.Elements() {
		result.Add(f(e))
	}
	return result
}

// FlatMap processes each element into a list of elements, then flattens the lists into a single list.
func FlatMap[T, U any](collection *Collection[T], f FlatMapper[T, U]) *Collection[U] {
	result := New[U]()
	for _, e := range collection.Elements() {
		u := f(e)
		result.Concat(u)
	}
	return result
}
