package lazy

import "sync"

type ValueProducer[T any] func() T
type ValueProducerWithError[T any] func() (T, error)

func New[T any](fn ValueProducer[T]) ValueProducer[T] {
	var once sync.Once
	var v T
	return func() T {
		once.Do(func() { v = fn() })
		return v
	}
}

func NewErrorable[T any](fn ValueProducerWithError[T]) ValueProducerWithError[T] {
	var once sync.Once
	var v T
	var err error
	return func() (T, error) {
		once.Do(func() { v, err = fn() })
		return v, err
	}
}

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
