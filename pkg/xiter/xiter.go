package xiter

import (
	"iter"
	"slices"
)

func Slice[Slice ~[]E, E any](s Slice) Seq[E] {
	return Seq[E](slices.Values(s))
}

type Seq[T any] iter.Seq[T]

// Map returns an iterator over f applied to seq.
func (s Seq[T]) Map[U any](f func(T) U) Seq[U] {
	return func(yield func(U) bool) {
		for v := range s {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func (s Seq[T]) Iter() iter.Seq[T] {
	return iter.Seq[T](s)
}

func (s Seq[T]) Collect() []T {
	return slices.Collect(s.Iter())
}

func (s Seq[T]) All(f func(T) bool) bool {
	for v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}
