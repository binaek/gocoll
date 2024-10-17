package collection

// Collection represents a generic collection of elements of any type.
type Collection[T any] struct {
	elements []T
}

// New creates and returns a new instance of Collection.
func New[T any]() *Collection[T] {
	return &Collection[T]{}
}

func From[T any](elements ...T) *Collection[T] {
	return &Collection[T]{elements: elements}
}

// Add appends an element to the collection.
func (c *Collection[T]) Add(element T) {
	c.elements = append(c.elements, element)
}

// Concat merges another collection into the current collection.
func (c *Collection[T]) Concat(collection *Collection[T]) {
	c.elements = append(c.elements, collection.elements...)
}

// Remove deletes the first element that matches the Finder condition.
func (c *Collection[T]) Remove(f Predicate[T]) {
	for i, e := range c.elements {
		if f(e) {
			c.elements = append(c.elements[:i], c.elements[i+1:]...)
			return
		}
	}
}

// Count returns the number of elements in the collection.
func (c *Collection[T]) Count() int {
	return len(c.elements)
}

// Elements returns a slice of all elements in the collection.
func (c *Collection[T]) Elements(opts ...ElementsOption) []T {
	config := &elementsConfig{}
	for _, opt := range opts {
		opt(config)
	}
	if config.noNil && c.elements == nil {
		return []T{}
	}
	return c.elements
}

// Clear removes all elements from the collection.
func (c *Collection[T]) Clear() {
	c.elements = []T{}
}

// IsEmpty checks if the collection is empty.
func (c *Collection[T]) IsEmpty() bool {
	return len(c.elements) == 0
}

type ElementsOption func(*elementsConfig)

type elementsConfig struct {
	noNil bool
}

// NoNil prevents the Elements method from returning nil.
// If the collection is nil, an empty slice is returned instead.
func NoNil() ElementsOption {
	return func(c *elementsConfig) {
		c.noNil = true
	}
}
