## Double Ended Queue implementation for Golang ##

This is a classical ring-buffer based deque (double ended queue) implementation.

The API is inspired by C++ [`std::deque`](https://en.cppreference.com/w/cpp/container/deque) with separated access and pop operations (back and front).

### Example ###

_For full code see [BFS](examples/bfs)_

```Go
// Tree models a classical binary tree
type Tree struct {
	Key   int
	Left  *Tree
	Right *Tree
}

// BFS traverse a tree in breadth first order printing keys, level by level
func BFS(t *Tree) error {
	if t == nil {
		return nil
	}
	q := deque.NewDeque[*Tree]()
	q.PushBack(t)
	q.PushBack(nil)
	for !q.IsEmpty() {
		cur, err := q.Front()
		if err != nil {
			return err
		}
		err = q.PopFront()
		if err != nil {
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
