package gocoll

type Collection[T any] struct {
	elements []T
}

func New[T any]() *Collection[T] {
	return &Collection[T]{}
}

func (c *Collection[T]) Add(element T) {
	c.elements = append(c.elements, element)
}

func (c *Collection[T]) Concat(collection Collection[T]) {
	c.elements = append(c.elements, collection.elements...)
}

func (c *Collection[T]) Remove(f Finder[T]) {
	for i, e := range c.elements {
		if f(e) {
			c.elements = append(c.elements[:i], c.elements[i+1:]...)
			return
		}
	}
}

func (c *Collection[T]) Contains(f Finder[T]) bool {
	for _, e := range c.elements {
		if f(e) {
			return true
		}
	}
	return false
}

func (c *Collection[T]) Count() int {
	return len(c.elements)
}

func (c *Collection[T]) Elements() []T {
	return c.elements
}

func (c *Collection[T]) Clear() {
	c.elements = []T{}
}

func (c *Collection[T]) IsEmpty() bool {
	return len(c.elements) == 0
}
