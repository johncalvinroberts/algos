package main

import (
  "fmt"
  "math/rand"
  "time"
)

type BinaryNode[T comparable] struct {
  value T
  left  *BinaryNode[T]
  right *BinaryNode[T]
}

func BuildTree(depth int) *BinaryNode[int] {
  if depth == 0 {
    return nil
  }
  rand.Seed(time.Now().UnixNano())
  min := 0
  max := 10
  fmt.Println(rand.Intn(max-min+1) + min)
  value := rand.Int()
  root := &BinaryNode[int]{value: value}
  root.left = &BinaryNode[int]{value: value + 1}
  root.right = &BinaryNode[int]{value: value + 2}
  root.right.left = &BinaryNode[int]{value: value + 3}
  root.right.right = &BinaryNode[int]{value: value + 4}
  root.left.left = &BinaryNode[int]{value: 5}
  root.left.right = &BinaryNode[int]{value: value + 6}
  root.left.right.right = &BinaryNode[int]{value: value + 7}
  root.left.right.right.right = &BinaryNode[int]{value: value + 8}
  root.left.right.right.right = &BinaryNode[int]{value: value + 9}
  return root
}

func main() {
  tree := BuildTree(10)
  // fmt.Println(BreadthFirstSearch[int](100, tree, nil))
}
