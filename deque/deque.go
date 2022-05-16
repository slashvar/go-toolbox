package deque

// Deque[T] describes a double-ended queue of elements of type T
type Deque[T any] struct {
	buffer   []T
	first    int
	length   int
	capacity int
}

// NewDeque[T] creates an empty Deque[T]
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		buffer:   nil,
		first:    0,
		length:   0,
		capacity: 0,
	}
}

// NotEnoughElementsError is returned when trying to get or pop an element from an empty Deque
type NotEnoughElementsError struct{}

// Error implements error interface
func (e NotEnoughElementsError) Error() string {
	return "Deque is empty"
}

// IsEmpty returns true if and only if the deque contains no element
func (q *Deque[T]) IsEmpty() bool {
	return q.length == 0
}

// Size returns the number of elements in the deque
func (q *Deque[T]) Size() int {
	return q.length
}

// grow doubles the underlying storage capacity (or fix it to 1 if it was 0)
func (q *Deque[T]) grow() {
	newCapacity := q.capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}
	q.resize(newCapacity)
}

// resize resizes the internal buffer, keep only the n first elements
func (q *Deque[T]) resize(n int) {
	newBuffer := make([]T, n)
	for i := 0; i < n && i < q.length; i++ {
		newBuffer[i] = q.buffer[(q.first+i)%q.capacity]
	}
	q.buffer = newBuffer
	q.capacity = n
	q.first = 0
	if n < q.length {
		q.length = n
	}
}

// Back returns a pointer to the element at the back of the queue
// returns NotEnoughElementsError if the deque is empty
func (q *Deque[T]) Back() (*T, error) {
	if q.IsEmpty() {
		return nil, NotEnoughElementsError{}
	}
	return &q.buffer[(q.first+q.length-1)%q.capacity], nil
}

// Front returns a pointer to the element at the front of the queue
// returns NotEnoughElementsError if the deque is empty
func (q *Deque[T]) Front() (*T, error) {
	if q.IsEmpty() {
		return nil, NotEnoughElementsError{}
	}
	return &q.buffer[q.first], nil
}

// PushBack inserts an element at the back of the deque
func (q *Deque[T]) PushBack(e T) {
	if q.length == q.capacity {
		q.grow()
	}
	q.buffer[(q.length+q.first)%q.capacity] = e
	q.length++
}

// PushFront inserts an element at the front of the deque
func (q *Deque[T]) PushFront(e T) {
	if q.length == q.capacity {
		q.grow()
	}
	if q.first == 0 {
		q.first = q.capacity - 1
	} else {
		q.first = q.first - 1
	}
	q.buffer[q.first] = e
	q.length++
}

// PopBack removes the element at the back
// returns NotEnoughElementsError if the deque is empty
func (q *Deque[T]) PopBack() error {
	if q.IsEmpty() {
		return NotEnoughElementsError{}
	}
	q.length--
	return nil
}

// PopFront removes the element at the front
// returns NotEnoughElementsError if the deque is empty
func (q *Deque[T]) PopFront() error {
	if q.IsEmpty() {
		return NotEnoughElementsError{}
	}
	q.first = (q.first + 1) % q.capacity
	q.length--
	return nil
}

// Clear removes all elements from the queue
func (q *Deque[T]) Clear() {
	q.length = 0
	q.first = 0
}

// nextPowerTwo returns the smallest power of two greater than n
func nextPowerTwo(n int) int {
	if n <= 0 {
		return 1
	}
	n--
	p := 1
	for ; n > 1; n = n >> 1 {
		p++
	}
	return 1 << p
}

// ShrinkToFit shrinks the internal buffer to at least q.Size()
func (q *Deque[T]) ShrinkToFit() {
	if q.capacity < 4*q.length {
		return
	}
	q.resize(nextPowerTwo(2 * q.length))
}

// Get returns nth element if it exists
func (q *Deque[T]) Get(n int) (*T, error) {
	if n >= q.length {
		return nil, NotEnoughElementsError{}
	}
	return &q.buffer[(q.first+n)%q.capacity], nil
}
