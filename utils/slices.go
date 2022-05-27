package utils

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
	var r []T
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}
