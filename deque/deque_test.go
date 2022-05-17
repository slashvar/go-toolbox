package deque

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmpty(t *testing.T) {
	type testCases struct {
		name     string
		expected bool
		deque    *Deque[int]
	}
	cases := []testCases{
		{name: "empty deque", expected: true, deque: NewDeque[int]()},
		{name: "non empty deque", expected: false, deque: func() *Deque[int] {
			q := NewDeque[int]()
			q.buffer = append(q.buffer, 0)
			q.capacity = len(q.buffer)
			q.length = 1
			return q
		}()},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.deque.IsEmpty())
		})
	}
}

func buildDeque(first, capacity int, elem []int) *Deque[int] {
	q := NewDeque[int]()
	q.buffer = make([]int, capacity)
	q.first = first
	q.capacity = capacity
	q.length = len(elem)
	for i, e := range elem {
		q.buffer[(i+first)%capacity] = e
	}
	return q
}

func TestGrow(t *testing.T) {
	type testCases struct {
		name     string
		elem     []int
		first    int
		capacity int
	}
	cases := []testCases{
		{name: "grow empty", elem: []int{}, first: 0, capacity: 0},
		{name: "grow first=0", elem: []int{1, 2}, first: 0, capacity: 2},
		{name: "grow not full", elem: []int{1, 2}, first: 0, capacity: 3},
		{name: "grow shifted", elem: []int{1, 2}, first: 1, capacity: 2},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDeque(tt.first, tt.capacity, tt.elem)
			q.grow()
			if tt.capacity == 0 {
				require.Equal(t, 1, q.capacity)
			} else {
				require.Equal(t, tt.capacity*2, q.capacity)
			}
			require.Equal(t, 0, q.first)
			require.Equal(t, len(tt.elem), q.length)
			for i := 0; i < q.length; i++ {
				require.Equal(t, tt.elem[i], q.buffer[i])
			}
		})
	}
}

func TestPushBack(t *testing.T) {
	type testCases struct {
		name     string
		elem     []int
		first    int
		capacity int
		toPush   []int
	}
	cases := []testCases{
		{name: "PushBack one element in empty deque", elem: []int{}, first: 0, capacity: 0, toPush: []int{1}},
		{name: "PushBack two elements in empty deque", elem: []int{}, first: 0, capacity: 0, toPush: []int{1, 2}},
		{name: "PushBack some elements in empty deque", elem: []int{}, first: 0, capacity: 0, toPush: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "PushBack some elements in a deque", elem: []int{1, 2, 3}, first: 0, capacity: 4, toPush: []int{4, 5}},
		{name: "PushBack some elements in a deque", elem: []int{1, 2, 3}, first: 3, capacity: 4, toPush: []int{4, 5}},
		{name: "PushBack some elements in a full deque", elem: []int{1, 2, 3, 4}, first: 0, capacity: 4, toPush: []int{5, 6}},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDeque(tt.first, tt.capacity, tt.elem)
			for _, e := range tt.toPush {
				q.PushBack(e)
			}
			require.Equal(t, len(tt.elem)+len(tt.toPush), q.length)
			require.Equal(t, len(tt.elem)+len(tt.toPush), q.Size())
			elem := append([]int{}, tt.elem...)
			elem = append(elem, tt.toPush...)
			for i, e := range elem {
				require.Equal(t, e, q.buffer[(q.first+i)%q.capacity])
			}
		})
	}
}

func TestPushFront(t *testing.T) {
	type testCases struct {
		name     string
		elem     []int
		first    int
		capacity int
		toPush   []int
	}
	cases := []testCases{
		{name: "PushBack one element in empty deque", elem: []int{}, first: 0, capacity: 0, toPush: []int{1}},
		{name: "PushBack two elements in empty deque", elem: []int{}, first: 0, capacity: 0, toPush: []int{1, 2}},
		{name: "PushBack some elements in empty deque", elem: []int{}, first: 0, capacity: 0, toPush: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{name: "PushBack some elements in a deque", elem: []int{1, 2, 3}, first: 0, capacity: 4, toPush: []int{4, 5}},
		{name: "PushBack some elements in a deque", elem: []int{1, 2, 3}, first: 3, capacity: 4, toPush: []int{4, 5}},
		{name: "PushBack some elements in a full deque", elem: []int{1, 2, 3, 4}, first: 0, capacity: 4, toPush: []int{5, 6}},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDeque(tt.first, tt.capacity, tt.elem)
			for _, e := range tt.toPush {
				q.PushFront(e)
			}
			require.Equal(t, len(tt.elem)+len(tt.toPush), q.length)
			require.Equal(t, len(tt.elem)+len(tt.toPush), q.Size())
			var elem []int
			for i := len(tt.toPush); i > 0; i-- {
				elem = append(elem, tt.toPush[i-1])
			}
			elem = append(elem, tt.elem...)
			for i, e := range elem {
				require.Equal(t, e, q.buffer[(q.first+i)%q.capacity])
			}
		})
	}
}

func TestPopBack(t *testing.T) {
	type testCases struct {
		name     string
		elem     []int
		first    int
		capacity int
	}
	cases := []testCases{
		{name: "Back/PopBack on empty deque", elem: []int{}, first: 0, capacity: 0},
		{name: "Back/PopBack with one element", elem: []int{1}, first: 0, capacity: 1},
		{name: "Back/PopBack with some elements", elem: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, first: 0, capacity: 16},
		{name: "Back/PopBack with some elements not starting at 0", elem: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, first: 10, capacity: 16},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDeque(tt.first, tt.capacity, tt.elem)
			if len(tt.elem) == 0 {
				ref := NotEnoughElementsError{}
				_, err := q.Back()
				require.Error(t, err)
				require.True(t, errors.Is(err, ref))
				err = q.PopBack()
				require.Error(t, err)
				require.True(t, errors.Is(err, ref))

			}
			var popElem []int
			for i := 0; i < len(tt.elem); i++ {
				e, err := q.Back()
				require.NoError(t, err)
				popElem = append(popElem, e)
				require.NoError(t, q.PopBack())
			}
			require.Equal(t, len(tt.elem), len(popElem))
			require.True(t, q.IsEmpty())
			for i, e := range tt.elem {
				require.Equal(t, e, popElem[len(popElem)-i-1])
			}
		})
	}
}

func TestPopFront(t *testing.T) {
	type testCases struct {
		name     string
		elem     []int
		first    int
		capacity int
	}
	cases := []testCases{
		{name: "Front/PopFront on empty deque", elem: []int{}, first: 0, capacity: 0},
		{name: "Front/PopFront with one element", elem: []int{1}, first: 0, capacity: 1},
		{name: "Front/PopFront with some elements", elem: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, first: 0, capacity: 16},
		{name: "Front/PopFront with some elements not starting at 0", elem: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, first: 10, capacity: 16},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDeque(tt.first, tt.capacity, tt.elem)
			if len(tt.elem) == 0 {
				ref := NotEnoughElementsError{}
				_, err := q.Front()
				require.Error(t, err)
				require.True(t, errors.Is(err, ref))
				err = q.PopFront()
				require.Error(t, err)
				require.True(t, errors.Is(err, ref))

			}
			var popElem []int
			for i := 0; i < len(tt.elem); i++ {
				e, err := q.Front()
				require.NoError(t, err)
				popElem = append(popElem, e)
				require.NoError(t, q.PopFront())
			}
			require.Equal(t, len(tt.elem), len(popElem))
			require.True(t, q.IsEmpty())
			for i, e := range tt.elem {
				require.Equal(t, e, popElem[i])
			}
		})
	}
}

func children(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func TestBFS(t *testing.T) {
	var tree []int
	for i := 0; i < 63; i++ {
		tree = append(tree, i)
	}
	q := NewDeque[int]()
	q.PushBack(0)
	var res []int
	for !q.IsEmpty() {
		cur, err := q.Front()
		require.NoError(t, err)
		require.NoError(t, q.PopFront())
		res = append(res, tree[cur])
		left, right := children(cur)
		if left < len(tree) {
			q.PushBack(left)
		}
		if right < len(tree) {
			q.PushBack(right)
		}
	}
	require.Equal(t, tree, res)
}

func TestBFS2(t *testing.T) {
	var tree []int
	for i := 0; i < 63; i++ {
		tree = append(tree, i)
	}
	q := NewDeque[int]()
	q.PushFront(0)
	var res []int
	for !q.IsEmpty() {
		cur, err := q.Back()
		require.NoError(t, err)
		require.NoError(t, q.PopBack())
		res = append(res, tree[cur])
		left, right := children(cur)
		if left < len(tree) {
			q.PushFront(left)
		}
		if right < len(tree) {
			q.PushFront(right)
		}
	}
	require.Equal(t, tree, res)
}

func isSorted(v []int) bool {
	i := 0
	for ; i < len(v)-1 && v[i] <= v[i+1]; i++ {
		continue
	}
	return i == len(v)-1
}

func TestDFS(t *testing.T) {
	tree := []int{1, 2, 5, 3, 4, 6, 7}
	q := NewDeque[int]()
	q.PushBack(0)
	var res []int
	for !q.IsEmpty() {
		cur, err := q.Back()
		require.NoError(t, err)
		require.NoError(t, q.PopBack())
		res = append(res, tree[cur])
		left, right := children(cur)
		if right < len(tree) {
			q.PushBack(right)
		}
		if left < len(tree) {
			q.PushBack(left)
		}
	}
	require.True(t, isSorted(res))
}

func TestDFS2(t *testing.T) {
	tree := []int{1, 2, 5, 3, 4, 6, 7}
	q := NewDeque[int]()
	q.PushFront(0)
	var res []int
	for !q.IsEmpty() {
		cur, err := q.Front()
		require.NoError(t, err)
		require.NoError(t, q.PopFront())
		res = append(res, tree[cur])
		left, right := children(cur)
		if right < len(tree) {
			q.PushFront(right)
		}
		if left < len(tree) {
			q.PushFront(left)
		}
	}
	require.Equal(t, len(tree), len(res))
	require.True(t, isSorted(res))
}

func TestMixed(t *testing.T) {
	q := NewDeque[int]()
	for i := 0; i < 10; i++ {
		q.PushFront(0)
		q.PushBack(1)
	}
	require.False(t, q.IsEmpty())
	require.Equal(t, 20, q.Size())
	var front []int
	var back []int
	for !q.IsEmpty() {
		f, err := q.Front()
		require.NoError(t, err)
		front = append(front, f)
		require.NoError(t, q.PopFront())
		b, err := q.Back()
		require.NoError(t, err)
		back = append(back, b)
		require.NoError(t, q.PopBack())
	}
	require.True(t, q.IsEmpty())
	require.Equal(t, 0, q.Size())
	require.Equal(t, 10, len(front))
	require.Equal(t, 10, len(back))
	for _, e := range front {
		require.Equal(t, 0, e)
	}
	for _, e := range back {
		require.Equal(t, 1, e)
	}
}

func TestNextPowerTwo(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	expected := []int{1, 2, 2, 4, 4, 8, 8, 8, 8, 16, 16, 16, 16, 16, 16, 16, 16}
	for i, v := range input {
		r := nextPowerTwo(v)
		require.Equal(t, expected[i], r)
	}
}

func TestShrinkToFit(t *testing.T) {
	type testCases struct {
		name     string
		elem     []int
		first    int
		capacity int
	}
	cases := []testCases{
		{name: "ShrinkToFit empty deque", elem: []int{}, first: 0, capacity: 0},
		{name: "ShrinkToFit nominal case", elem: []int{1}, first: 0, capacity: 4},
		{name: "ShrinkToFit shifted", elem: []int{1, 2}, first: 1, capacity: 8},
		{name: "ShrinkToFit no effect (full)", elem: []int{1, 2}, first: 0, capacity: 2},
		{name: "ShrinkToFit no effect (partial)", elem: []int{1, 2}, first: 0, capacity: 4},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDeque(tt.first, tt.capacity, tt.elem)
			q.ShrinkToFit()
			require.Equal(t, len(tt.elem), q.length)
			require.True(t, q.length <= q.capacity)
			for i := 0; i < q.length; i++ {
				require.Equal(t, tt.elem[i], q.buffer[i])
			}
		})
	}
}

func TestGet(t *testing.T) {
	type testCases struct {
		name     string
		elem     []int
		first    int
		capacity int
		howMany  int
	}
	cases := []testCases{
		{name: "Get with one element", elem: []int{1}, first: 0, capacity: 1, howMany: 1},
		{name: "Get with some elements", elem: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, first: 0, capacity: 16, howMany: 10},
		{name: "Get with some elements not starting at 0", elem: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, first: 10, capacity: 16, howMany: 10},
		{name: "Get on empty deque", elem: []int{}, first: 0, capacity: 0, howMany: 1},
		{name: "Get too far", elem: []int{1}, first: 0, capacity: 2, howMany: 10},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDeque(tt.first, tt.capacity, tt.elem)
			for i, e := range tt.elem {
				g, err := q.Get(i)
				require.NoError(t, err)
				require.Equal(t, e, *g)
			}
			for i := len(tt.elem); i < tt.howMany; i++ {
				_, err := q.Get(i)
				require.Error(t, err)
				ref := NotEnoughElementsError{}
				require.True(t, errors.Is(err, ref))
			}
		})
	}
}
