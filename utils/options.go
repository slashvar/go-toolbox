package utils

type Option[T any] interface {
	HasValue() bool
	Value() T
}

// OptionImplem model an optional value
type OptionImplem[T any] struct {
	content T
}

func NilOption[T any]() *OptionImplem[T] {
	return nil
}

func NewOption[T any](v T) *OptionImplem[T] {
	return &OptionImplem[T]{content: v}
}

func (o *OptionImplem[T]) HasValue() bool {
	return o != nil
}

func (o *OptionImplem[T]) Value() T {
	return o.content
}

func OptionalMap[T1, T2 any](o Option[T1], f func(T1) T2) Option[T2] {
	if o.HasValue() {
		return NewOption(f(o.Value()))
	}
	return NilOption[T2]()
}

func OptionalFlatMap[T1, T2 any](o Option[T1], f func(T1) Option[T2]) Option[T2] {
	if o.HasValue() {
		return f(o.Value())
	}
	return NilOption[T2]()
}

func OptionalElse[T any](o Option[T], f func()) Option[T] {
	if !o.HasValue() {
		f()
	}
	return o
}
