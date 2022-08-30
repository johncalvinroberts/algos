package main

import "fmt"

/*
	Implements a tree using generics in go
*/

type orderable interface {
	~int8 | ~int16 | ~int32 | ~float32 | ~float64
}
type Node[T orderable] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func (n Node[T]) insert(val T) {
	defer func() {
		n.print()
		fmt.Println("-")
	}()
	if val > n.val {
		x := n.val
		n.val = val
		n.insert(x)
	}

	if n.right == nil {
		n.right = &Node[T]{val: val}
		return
	}

	if n.left == nil && val < n.right.val {
		n.left = &Node[T]{val: val}
	}

	if n.right != nil && n.right.val == val {
		// node already exists on right side
		return
	}

	if (n.right != nil && val > n.right.val) || (n.left != nil && val > n.left.val) {
		n.right.insert(val)
		return
	}

}

func (n Node[T]) print() {
	fmt.Printf("n: %v, left: %v, right: %v\n", n.val, n.left, n.right)
	if n.right != nil {
		n.right.print()
	}

	if n.left != nil {
		n.left.print()
	}

}

func main() {
	t := &Node[int8]{val: 11}
	t.insert(12)
	t.insert(9)
	t.insert(10)
	t.print()
}
