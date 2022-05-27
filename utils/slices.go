package utils

import (
	"golang.org/x/exp/constraints"
)

// Map converts a slice of InType into a slice of OutType using function f
func Map[InType, OutType any](in []InType, f func(InType) OutType) []OutType {
	r := make([]OutType, 0, len(in))
	for _, e := range in {
		r = append(r, f(e))
	}
	return r
}

// Iter applies f to all elements of s
func Iter[T any](s []T, f func(T)) {
	for _, e := range s {
		f(e)
	}
}

// Accumulate traverses s and accumulates results of function f with initial value init
// Accumulate(f, init, {a1, ..., bn}) -> f( ... f(f(init, a1), a2) ..., an)
func Accumulate[InType, OutType any](f func(OutType, InType) OutType, init OutType, s []InType) OutType {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

// ForAll returns true if f applies to all elements of s returns true
func ForAll[T any](f func(T) bool, s []T) bool {
	r := true
	for i := 0; i < len(s) && r; i++ {
		r = r && f(s[i])
	}
	return r
}

// Exists returns true if f returns true for at least one elements of s
func Exists[T any](f func(T) bool, s []T) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

// Filter returns all the elements of s for which f returns true
func Filter[T any](f func(T) bool, s []T) []T {
	acc := func(r []T, e T) []T {
		if f(e) {
			return append(r, e)
		}
		return r
	}
	return Accumulate(acc, []T{}, s)
}

// OptionalFilter returns the slice made of f(e) for all e in s such that f(e).HasValue() is true
func OptionalFilter[T1, T2 any](f func(T1) Option[T2], s []T1) []T2 {
	acc := func(r []T2, e T1) []T2 {
		_ = OptionalDo(f(e), func(x T2) { r = append(r, x) })
		return r
	}
	return Accumulate(acc, []T2{}, s)
}

// Reverse reverse elements in s
func Reverse[T any](s []T) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

// Insert inserts x as position pos in s
// if pos >= len(s) insert at the end of s
// panic if pos is less than 0
func Insert[T any](x T, pos int, s []T) []T {
	if pos < 0 {
		panic("inserting at a negative position")
	}
	s = append(s, x)
	for i := len(s) - 1; i > pos; i-- {
		s[i], s[i-1] = s[i-1], s[i]
	}
	return s
}

// LowerBound returns the position of the first elements that is not smaller than x
// LowerBound returns a coherent result only if s is sorted
func LowerBound[T constraints.Ordered](x T, s []T) int {
	left, right := 0, len(s)
	for left < right {
		mid := left + (right-left)/2
		if s[mid] == x {
			return mid
		}
		if x < s[mid] {
			right = mid
			continue
		}
		left = mid + 1
	}
	return left
}

// Partition moves elements to form a partition and returns the position of the pivot
// Let p := Partition(x, s) then for all i < p s[i] < x and for all i >= p, s[i] >= x
func Partition[T constraints.Ordered](x T, s []T) int {
	pivot := 0
	for i := 0; i < len(s); i++ {
		if s[i] < x {
			s[i], s[pivot] = s[pivot], s[i]
			pivot++
		}
	}
	return pivot
}

// PartitionFilter moves elements to form a partition w.r.t to f
// Let p := PartitionFilter(f, s) then for all i < p, f([i]) is true and for all i >= p f(s[i]) is false
func PartitionFilter[T any](f func(T) bool, s []T) int {
	pivot := 0
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			s[i], s[pivot] = s[pivot], s[i]
			pivot++
		}
	}
	return pivot
}

// Combine combines elements of s1 and s2 using zipper function
// returned slice of the length of the smallest input
func Combine[T1, T2, T3 any](s1 []T1, s2 []T2, zipper func(T1, T2) T3) []T3 {
	var r []T3
	for i := 0; i < len(s1) && i < len(s2); i++ {
		r = append(r, zipper(s1[i], s2[i]))
	}
	return r
}
