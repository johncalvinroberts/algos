package main

import "fmt"

func foo(n int) int {
  if n == 1 {
    return 1
  }

  return n + foo(n - 1)
}

func main() {
  fmt.Println(foo(5))
}
