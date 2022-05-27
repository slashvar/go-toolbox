package utils

type Option[T any] interface {
	HasValue() bool
	Value() T
}

// OptionImplem model an optional value
type OptionImplem[T any] struct {
	content T
}

// NilOption returns an empty optional value
func NilOption[T any]() *OptionImplem[T] {
	return nil
}

// NewOption returns an optional value containing v
func NewOption[T any](v T) *OptionImplem[T] {
	return &OptionImplem[T]{content: v}
}

// HasValue returns true if the optional has a value
func (o *OptionImplem[T]) HasValue() bool {
	return o != nil
}

// Value returns the content of the optional value, panic if empty
func (o *OptionImplem[T]) Value() T {
	if o == nil {
		panic("request value of an empty optional value")
	}
	return o.content
}

// OptionalMap returns an option value result from f(o.Value()) if o is not empty, returns an empty optional otherwise
func OptionalMap[T1, T2 any](o Option[T1], f func(T1) T2) Option[T2] {
	if o.HasValue() {
		return NewOption(f(o.Value()))
	}
	return NilOption[T2]()
}

// OptionalFlatMap returns the result of f(o.Value()) if o has a value, returns an empty optional otherwise
// Similar to OptionalMap but for function that already returns an optional type
func OptionalFlatMap[T1, T2 any](o Option[T1], f func(T1) Option[T2]) Option[T2] {
	if o.HasValue() {
		return f(o.Value())
	}
	return NilOption[T2]()
}

// OptionalElse calls f if o is empty and returns o in all cases
func OptionalElse[T any](o Option[T], f func()) Option[T] {
	if !o.HasValue() {
		f()
	}
	return o
}

// OptionalDo calls f(o.Value()) if o is not empty and returns o in all cases
func OptionalDo[T any](o Option[T], f func(T)) Option[T] {
	if o.HasValue() {
		f(o.Value())
	}
	return o
}
