interface Queue<T> {
  enqueue(item: T): void
  dequeue(): T
  peek(): T
}

interface Stack<T> {
  push(item: T): void
  pop(item: T): T
  head: T
}