package collection

// Map applies a function to each element of a collection and returns a new collection with the results.
func Map[T, U any](collection *Collection[T], f Mapper[T, U]) *Collection[U] {
	result := []U{}
	for _, e := range collection.Elements() {
		result = append(result, f(e))
	}
	return From(result...)
}

// FlatMap processes each element into a list of elements, then flattens the lists into a single list.
func FlatMap[T, U any](collection *Collection[T], f FlatMapper[T, U]) *Collection[U] {
	result := []U{}
	for _, e := range collection.Elements() {
		u := f(e)
		result = append(result, u.elements...)
	}
	return From(result...)
}
