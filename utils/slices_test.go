package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func buildSlice(n, step int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = i * step
	}
	return r
}

func TestMap(t *testing.T) {
	type testCases[T, U any] struct {
		input  []T
		output []U
		f      func(T) U
		name   string
	}
	cases := []testCases[int, int]{
		{
			input:  []int{},
			output: []int{},
			f:      func(i int) int { return i },
			name:   "empty input",
		},
		{
			input:  buildSlice(10, 1),
			output: buildSlice(10, 1),
			f:      func(i int) int { return i },
			name:   "identity map",
		},
		{
			input:  buildSlice(10, 1),
			output: buildSlice(10, 2),
			f:      func(i int) int { return 2 * i },
			name:   "double map",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			res := Map(tt.input, tt.f)
			require.Equal(t, tt.output, res)
		})
	}
}

func TestIter(t *testing.T) {
	for i := 0; i < 100; i++ {
		in := buildSlice(i, 1)
		acc := 0
		f := func(n int) { acc += n }
		Iter(in, f)
		require.Equal(t, i*(i-1)/2, acc)
	}
}

func TestAccumulate(t *testing.T) {
	for i := 0; i < 100; i++ {
		in := buildSlice(i, 1)
		f := func(a, i int) int { return a + i }
		r := Accumulate(f, 0, in)
		require.Equal(t, i*(i-1)/2, r)
	}
	for i := 0; i < 100; i++ {
		in := buildSlice(i, 1)
		f := func(a, i int) int { return a + i }
		r := Accumulate(f, i, in)
		require.Equal(t, i+i*(i-1)/2, r)
	}
}

func TestForAll(t *testing.T) {
	type testCases struct {
		name     string
		input    []int
		f        func(int) bool
		expected bool
	}
	cases := []testCases{
		{
			name:     "empty slice",
			input:    []int{},
			f:        func(int) bool { return false },
			expected: true,
		},
		{
			name:     "All even",
			input:    buildSlice(10, 2),
			f:        func(n int) bool { return n%2 == 0 },
			expected: true,
		},
		{
			name:     "No all even",
			input:    buildSlice(10, 1),
			f:        func(n int) bool { return n%2 == 0 },
			expected: false,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := ForAll(tt.f, tt.input)
			require.Equal(t, tt.expected, r)
		})
	}
}

func TestExists(t *testing.T) {
	type testCases struct {
		name     string
		input    []int
		f        func(int) bool
		expected bool
	}
	cases := []testCases{
		{
			name:     "empty slice",
			input:    []int{},
			f:        func(int) bool { return true },
			expected: false,
		},
		{
			name:     "One multiple of 7",
			input:    buildSlice(10, 3)[1:],
			f:        func(n int) bool { return n%7 == 0 },
			expected: true,
		},
		{
			name:     "No multiple of 7",
			input:    buildSlice(7, 3)[1:],
			f:        func(n int) bool { return n%7 == 0 },
			expected: false,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := Exists(tt.f, tt.input)
			require.Equal(t, tt.expected, r)
		})
	}
}

func TestFilter(t *testing.T) {
	type testCases struct {
		input  []int
		output []int
		f      func(int) bool
		name   string
	}
	cases := []testCases{
		{
			input:  []int{},
			output: []int{},
			f:      func(i int) bool { return i%2 == 0 },
			name:   "empty slice",
		},
		{
			input:  buildSlice(10, 1),
			output: buildSlice(5, 2),
			f:      func(i int) bool { return i%2 == 0 },
			name:   "slice of even",
		},
	}
	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := Filter(tt.f, tt.input)
			require.True(t, eq(tt.output, r))
		})
	}
}

func TestReverse(t *testing.T) {
	equalRev := func(s1, s2 []int) bool {
		if len(s1) != len(s2) {
			return false
		}
		for i := range s1 {
			if s1[i] != s2[len(s2)-i-1] {
				return false
			}
		}
		return true
	}
	type testCase struct {
		name string
		in   []int
	}
	cases := []testCase{
		{
			name: "reverse empty slice",
			in:   []int{},
		},
		{
			name: "reverse even length slice",
			in:   buildSlice(10, 1),
		},
		{
			name: "reverse odd length slice",
			in:   buildSlice(11, 1),
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := make([]int, len(tt.in))
			copy(r, tt.in)
			Reverse(r)
			require.True(t, equalRev(tt.in, r))
		})
	}
}

func TestOptionalFilter(t *testing.T) {
	type testCases struct {
		input  []int
		output []int
		f      func(int) Option[int]
		name   string
	}
	even := func(i int) Option[int] {
		if i%2 == 0 {
			return NewOption(i / 2)
		}
		return NilOption[int]()
	}
	cases := []testCases{
		{
			input:  []int{},
			output: []int{},
			f:      even,
			name:   "empty slice",
		},
		{
			input:  buildSlice(10, 1),
			output: buildSlice(5, 1),
			f:      even,
			name:   "slice of half value",
		},
	}
	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := OptionalFilter(tt.f, tt.input)
			require.True(t, eq(tt.output, r))
		})
	}
}

func TestInsert(t *testing.T) {
	type testCase struct {
		name  string
		in    []int
		value int
		pos   int
	}
	cases := []testCase{
		{
			name:  "insert in empty slice",
			in:    []int{},
			value: 0,
			pos:   0,
		},
		{
			name:  "insert in front",
			in:    buildSlice(10, 1),
			value: 42,
			pos:   0,
		},
		{
			name:  "insert at back",
			in:    buildSlice(10, 1),
			value: 42,
			pos:   9,
		},
		{
			name:  "insert in the middle",
			in:    buildSlice(10, 1),
			value: 42,
			pos:   5,
		},
		{
			name:  "insert too far",
			in:    buildSlice(10, 1),
			value: 42,
			pos:   20,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := Insert(tt.value, tt.pos, tt.in)
			require.Equal(t, len(tt.in)+1, len(r))
			if tt.pos < len(r) {
				require.Equal(t, tt.value, r[tt.pos])
				return
			}
			require.Equal(t, tt.value, r[len(r)-1])
		})
	}
}
