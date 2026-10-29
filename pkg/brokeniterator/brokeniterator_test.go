package brokeniterator

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

type sliceIterator struct {
	items []int
	calls int
}

func (it *sliceIterator) Next() (int, bool) {
	it.calls++
	x := it.items[0]
	it.items = it.items[1:]
	return x, len(it.items) > 0
}

func TestSeq(t *testing.T) {
	var got []int
	for x := range Seq[int](&sliceIterator{items: []int{1, 2, 3}}) {
		got = append(got, x)
	}
	require.Equal(t, []int{1, 2, 3}, got)
}

func TestSeqBreakStopsCalling(t *testing.T) {
	it := &sliceIterator{items: []int{1, 2, 3, 4}}
	for x := range Seq[int](it) {
		if x == 2 {
			break
		}
	}
	require.Equal(t, 2, it.calls)
}

func TestSeqRuntimeFrames(t *testing.T) {
	pcs := make([]uintptr, 16)
	n := runtime.Callers(0, pcs)
	count := 0
	for f := range Seq[runtime.Frame](runtime.CallersFrames(pcs[:n])) {
		require.NotEmpty(t, f.Function)
		count++
	}
	require.Positive(t, count)
}
