// Package brokeniterator adapts Next-style iterators, such as runtime.Frames, to iter.Seq.
package brokeniterator

import "iter"

// BrokenIterator is an iterable type using the Next idiom: Next returns the current value and whether more elements follow.
type BrokenIterator[T any] interface {
	Next() (T, bool)
}

// Seq returns a sequence usable in for-range loops from a BrokenIterator.
func Seq[T any](b BrokenIterator[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var x T
		more := true
		for more {
			x, more = b.Next()
			if !yield(x) {
				return
			}
		}
	}
}
