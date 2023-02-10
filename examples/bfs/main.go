package main

import (
	"fmt"
	"github.com/slashvar/go-toolbox/deque"
	"log"
)

/*
 * Using Deque[T] to traverse a tree using with a breadth first order (BFS)
 */

// Tree models a classical binary tree
type Tree struct {
	Key   int
	Left  *Tree
	Right *Tree
}

// toDot prints the tree in graphviz format
func toDot(t *Tree) {
	fmt.Println("digraph T {")
	if t != nil {
		toDotRec(t)
	}
	fmt.Println("}")
}

func toDotRec(t *Tree) {
	if t.Left != nil {
		fmt.Printf("\t%d -> %d;\n", t.Key, t.Left.Key)
		toDotRec(t.Left)
	}
	if t.Right != nil {
		fmt.Printf("\t%d -> %d;\n", t.Key, t.Right.Key)
		toDotRec(t.Right)
	}
}

// buildTreeRec builds a perfect tree (recursive helper)
func buildTreeRec(n int, k int) *Tree {
	t := Tree{Key: k, Left: nil, Right: nil}
	if n > 0 {
		t.Left = buildTreeRec(n-1, 2*k+1)
		t.Right = buildTreeRec(n-1, 2*k+2)
	}
	return &t
}

// buildTree builds a perfect tree of depth n uniformly filled
func buildTree(n int) *Tree {
	if n < 0 {
		return nil
	}
	return buildTreeRec(n, 0)
}

// BFS traverse a tree in breadth first order printing keys, level by level
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

func main() {
	t := buildTree(3)
	toDot(t)
	fmt.Println("")
	if err := BFS(t); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
}
