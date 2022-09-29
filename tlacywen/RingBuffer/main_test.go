package main

import "testing"

func TestMain(t *testing.T) {
	buffer := &RingBuffer[int]{}

	buffer.Push(5)
	if *buffer.Pop() != 5 {
		t.Fail()
	}

	if buffer.Pop() != nil {
		t.Fail()
	}

	buffer.Push(42)
	buffer.Push(9)
	if *buffer.Pop() != 42 {
		t.Fail()
	}
	if *buffer.Pop() != 9 {
		t.Fail()
	}
	if buffer.Pop() != nil {
		t.Fail()
	}

	buffer.Push(42)
	buffer.Push(9)
	buffer.Push(12)
	if *buffer.Get(2) != 12 {
		t.Fail()
	}
	if *buffer.Get(1) != 9 {
		t.Fail()
	}
	if *buffer.Get(0) != 42 {
		t.Fail()
	}
}
