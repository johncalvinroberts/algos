package main

/*

A simple singly-linked list with integers as values

*/

import "fmt"

type Node struct {
	next *Node
	key  int
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) length() int {
	sum := 0
	for n := l.head; n != nil; n = n.next {
		sum++
	}
	return sum
}

func (l *LinkedList) Append(key int) {
	node := &Node{key: key}
	list := l.head
	for list.next != nil {
		list = list.next
	}
	list.next = node
}

func (l *LinkedList) Prepend(key int) {
	node := &Node{key: key}
	list := l.head
	node.next = list
	l.head = node
}

func (l *LinkedList) InsertAt(where int, key int) {
	node := &Node{key: key}
	list := l.head
	for {
		if list.key == where {
			node.next = list.next
			list.next = node
			break
		}
		list = list.next
	}
}

func (l *LinkedList) RemoveAt(where int) {
	list := l.head
	for {
		if list.next.key == where {
			list.next = list.next.next
			break
		}
		list = list.next
	}
}

func (l *LinkedList) Get(key int) Node {
	list := l.head
	for {
		if list.key == key {
			return *list
		} else {
			list = list.next
		}
	}
}

func (l *LinkedList) Display() {
	fmt.Printf("LinkedList -- %d length: \n", l.length())
	list := l.head
	for list != nil {
		if list.next == nil {
			fmt.Printf("%d\n", list.key)
		} else {
			fmt.Printf("%d -> ", list.key)
		}
		list = list.next
	}
}

func main() {
	head := &Node{key: 420}
	l := LinkedList{head: head}
	l.Append(10)
	l.Append(200)
	l.Append(201)
	l.Prepend(101)
	l.Prepend(222222)
	l.Display()
	l.InsertAt(10, 123)
	l.Display()
	l.RemoveAt(10)
	l.Display()
	fmt.Print(l.Get(420))
}
