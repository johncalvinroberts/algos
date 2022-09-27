package main

import "fmt"

/*
Stack data structure made with Go generics
*/
type Node[T comparable] struct {
	value T
	prev  *Node[T]
}

type Stack[T comparable] struct {
	head   *Node[T]
	length int
}

func (s *Stack[T]) pop() *Node[T] {
	n := s.head
	if n != nil {
		s.length--
		s.head = n.prev
	} else {
		fmt.Println("The stack is empty.")
	}
	return n
}

func (s *Stack[T]) push(value T) {
	n := &Node[T]{value: value}
	s.length++
	if s.head == nil {
		s.head = n
		return
	}
	n.prev = s.head
	s.head = n
}

func (s *Stack[T]) Display() {
	fmt.Printf("LinkedList -- %d length: \n", s.length)
	n := s.head
	for n != nil {
		if n.prev == nil {
			fmt.Printf("%v\n", n.value)
		} else {
			fmt.Printf("%v -> ", n.value)
		}
		n = n.prev
	}
}

func main() {
	s := &Stack[int]{}
	s.push(111)
	s.push(1112)
	s.push(22)
	s.push(33)
	s.Display()
	s.pop()
	s.Display()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.Display()
}
