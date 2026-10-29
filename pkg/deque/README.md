# deque

Ring-buffer double-ended queue (deque) with generics.

The API follows C++ [`std::deque`](https://en.cppreference.com/w/cpp/container/deque): access and removal are separate calls.

```go
import "github.com/slashvar/go-toolbox/pkg/deque"
```

## API

`Deque[T]` is not safe for concurrent use. Create one with `deque.New[T]()`.

| Method          | Description                                                       |
|-----------------|-------------------------------------------------------------------|
| `PushBack(e)`   | Appends `e` at the back                                           |
| `PushFront(e)`  | Prepends `e` at the front                                         |
| `Back()`        | Returns the back element                                          |
| `Front()`       | Returns the front element                                         |
| `Get(n)`        | Returns the `n`th element from the front                          |
| `PopBack()`     | Removes the back element                                          |
| `PopFront()`    | Removes the front element                                         |
| `Size()`        | Returns the number of elements                                    |
| `IsEmpty()`     | Reports whether the deque holds no element                        |
| `Clear()`       | Removes all elements and keeps the allocated buffer               |
| `ShrinkToFit()` | Shrinks the buffer when its capacity is at least 4 times the size |

`Back`, `Front`, `Get`, `PopBack` and `PopFront` return `NotEnoughElementsError` when the deque is empty or `n` is out of range.

Pushes run in amortized constant time. The buffer doubles when full.

## Example

The full program is in [examples/bfs](../../examples/bfs).

```go
// BFS traverses a tree in breadth-first order, printing keys level by level
func BFS(t *Tree) error {
	if t == nil {
		return nil
	}
	q := deque.New[*Tree]()
	q.PushBack(t)
	q.PushBack(nil)
	for !q.IsEmpty() {
		cur, err := q.Front()
		if err != nil {
			return err
		}
		if err = q.PopFront(); err != nil {
			return err
		}
		if cur == nil {
			fmt.Println("")
			if !q.IsEmpty() {
				q.PushBack(nil)
			}
			continue
		}
		fmt.Printf("%d ", cur.Key)
		if cur.Left != nil {
			q.PushBack(cur.Left)
		}
		if cur.Right != nil {
			q.PushBack(cur.Right)
		}
	}
	return nil
}
```

A `nil` entry marks the end of a level.
