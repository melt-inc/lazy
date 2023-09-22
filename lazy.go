package lazy

import "sync"

type ValueProducer[T any] func() T
type ValueProducerWithError[T any] func() (T, error)

// New returns a function that will call fn only once, and return the same value
// on subsequent calls
func New[T any](fn ValueProducer[T]) ValueProducer[T] {
	var once sync.Once
	var v T
	return func() T {
		once.Do(func() { v = fn() })
		return v
	}
}

// NewErrorable returns a function that will call fn only once, and return the
// same value and error on subsequent calls
func NewErrorable[T any](fn ValueProducerWithError[T]) ValueProducerWithError[T] {
	var once sync.Once
	var v T
	var err error
	return func() (T, error) {
		once.Do(func() { v, err = fn() })
		return v, err
	}
}

// Must returns a function that will call fn only once, and panic if fn returns
// an error
func Must[T any](fn ValueProducerWithError[T]) ValueProducer[T] {
	var once sync.Once
	var v T
	var err error
	return func() T {
		once.Do(func() { v, err = fn() })
		if err != nil {
			panic(err)
		}
		return v
	}
}
