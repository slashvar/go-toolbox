package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	type testCases struct {
		name     string
		opt      Option[int]
		hasValue bool
		value    int
	}
	cases := []testCases{
		{
			name:     "nil option",
			opt:      NilOption[int](),
			hasValue: false,
		},
		{
			name:     "some value",
			opt:      NewOption(42),
			hasValue: true,
			value:    42,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.hasValue, tt.opt.HasValue())
			if tt.hasValue {
				require.Equal(t, tt.value, tt.opt.Value())
			}
		})
	}
}

func TestOptionalMap(t *testing.T) {
	identity := func(n int) int { return n }
	t.Run("call on nil option", func(t *testing.T) {
		var v Option[int] = NilOption[int]()
		r1 := OptionalMap(v, identity)
		require.False(t, r1.HasValue())
		r2 := OptionalMap(OptionalMap(v, identity), identity)
		require.False(t, r2.HasValue())
	})
	t.Run("call on option", func(t *testing.T) {
		var v Option[int] = NewOption(42)
		r1 := OptionalMap(v, identity)
		require.True(t, r1.HasValue())
		require.Equal(t, 42, r1.Value())
		r2 := OptionalMap(OptionalMap(v, identity), identity)
		require.True(t, r2.HasValue())
		require.Equal(t, 42, r2.Value())
	})
}

func TestOptionalFlatMap(t *testing.T) {
	identity := func(n int) Option[int] { return NewOption(n) }
	t.Run("call on nil option", func(t *testing.T) {
		var v Option[int] = NilOption[int]()
		r1 := OptionalFlatMap(v, identity)
		require.False(t, r1.HasValue())
		r2 := OptionalFlatMap(OptionalFlatMap(v, identity), identity)
		require.False(t, r2.HasValue())
	})
	t.Run("call on option", func(t *testing.T) {
		var v Option[int] = NewOption(42)
		r1 := OptionalFlatMap(v, identity)
		require.True(t, r1.HasValue())
		require.Equal(t, 42, r1.Value())
		r2 := OptionalFlatMap(OptionalFlatMap(v, identity), identity)
		require.True(t, r2.HasValue())
		require.Equal(t, 42, r2.Value())
	})
}

func TestElse(t *testing.T) {
	t.Run("OptionalElse call on non nil optional", func(t *testing.T) {
		b := false
		f := func() { b = true }
		r := OptionalElse[int](NewOption(42), f)
		require.False(t, b)
		require.True(t, r.HasValue())
	})
	t.Run("OptionalElse call on nil optional", func(t *testing.T) {
		b := false
		f := func() { b = true }
		r := OptionalElse[int](NilOption[int](), f)
		require.True(t, b)
		require.False(t, r.HasValue())
	})
}

func TestDo(t *testing.T) {
	t.Run("OptionalDo call on non nil optional", func(t *testing.T) {
		b := false
		f := func(int) { b = true }
		r := OptionalDo[int](NewOption(42), f)
		require.True(t, b)
		require.True(t, r.HasValue())
	})
	t.Run("OptionalDo call on nil optional", func(t *testing.T) {
		b := false
		f := func(int) { b = true }
		r := OptionalDo[int](NilOption[int](), f)
		require.False(t, b)
		require.False(t, r.HasValue())
	})
}
