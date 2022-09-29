package main

type Node[T comparable] struct {
	value T
}

type RingBuffer[T comparable] struct {
	head Node[T]
	tail Node[T]
	size int
}

func (r *RingBuffer[T]) Push(value T) {
	// TODO
}

func (r *RingBuffer[T]) Pop() *T {
	// TODO
	return nil
}
func (r *RingBuffer[T]) Get(value T) *T {
	// TODO
	return nil
}

func main() {

}
