package gocoll

type Finder[T any] func(T) bool
type Mapper[T comparable, U comparable] func(T) U
type FlatMapper[T comparable, U comparable] func(T) Collection[U]
type Reducer[T any, V any] func(V, T) V
type Grouper[T any, G comparable] func(T) G

func NotFinder[T any](f Finder[T]) Finder[T] {
	return func(e T) bool {
		return !f(e)
	}
}
