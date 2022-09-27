package main

import "fmt"

/*

A queue using golang generics

*/

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Queue[T comparable] struct {
	length int
	tail   *Node[T]
	head   *Node[T]
}

func (q *Queue[T]) enqueue(value T) {
	n := &Node[T]{value: value}
	if q.length == 0 {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.length++
}

func (q *Queue[T]) deque() *Node[T] {
	var ret *Node[T]
	if q.head != nil {
		ret = q.head
		q.head = ret.next
		q.length--
	} else {
		fmt.Println("The queue is empty.")
	}
	return ret
}

func (q *Queue[T]) peek() T {
	return q.head.value
}

func (q *Queue[T]) Display() {
	fmt.Printf("Queue -- %d length: \n", q.length)
	n := q.head
	for n != nil {
		if n.next == nil {
			fmt.Printf("%v\n", n.value)
		} else {
			fmt.Printf("%v -> ", n.value)
		}
		n = n.next
	}
}

func main() {
	fmt.Println("---- MAKING AN INT QUEUE ----")
	queue := &Queue[int]{length: 0, head: &Node[int]{value: 100}}
	queue.enqueue(120)
	fmt.Println(queue.peek())
	queue.enqueue(100)
	queue.enqueue(110101)
	queue.enqueue(10)
	queue.enqueue(12)
	queue.Display() // should show length as 5, 120 at the beginning, 12 at the end
	r := queue.deque()
	fmt.Printf("Giving service to %v\n", r)
	queue.Display() // should show length as 4, 100 at the beginning, 120 at the end
	queue.enqueue(1)
	queue.Display() // should show length as 5, 100 at the beginning, 1 at the end
}
