package main

/*
	Implements a tree using generics in go
*/

type Node[T cancompare] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func (n Node[T]) addChild(val T) {
	if T.compare(val) {
		n.right = &Node[T]{val: val}
	} else {
		n.left = &Node[T]{val: val}
	}
}

type cancompare interface {
	compare() bool
}

type IntNode struct{}

func (i IntNode) compare(x int, y int) bool {
	return x > y
}

type StringNode struct{}

func (s StringNode) compare(x string, y string) bool {
	return len(x) > len(y)
}

func main() {
	// TODO
}
