package gocoll

// Predicate defines a function type that determines if a given element of type T satisfies a condition.
type Predicate[T any] func(T) bool

// Equality defines a function type that determines if two elements of type T are equal.
type Equality[T any] func(T, T) bool

// Mapper defines a function type that maps an element of type T to a new element of type U.
type Mapper[T any, U any] func(T) U

// FlatMapper defines a function type that maps an element of type T to a Collection of elements of type U.
type FlatMapper[T any, U any] func(T) *Collection[U]

// Reducer defines a function type that accumulates a value V using an element of type T.
type Reducer[T any, V any] func(V, T) V

// GroupKey defines a function type that categorizes an element of type T into a group of type G.
type GroupKey[T any, G comparable] func(T) G

// NotFinder returns a Finder function that negates the result of the given Finder function f.
func NotFinder[T any](f Predicate[T]) Predicate[T] {
	return func(e T) bool {
		return !f(e)
	}
}
