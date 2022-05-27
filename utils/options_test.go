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
