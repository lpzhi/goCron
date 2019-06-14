package main

import (
	"fmt"
	"goCron/prepare/queue"
)

func main() {
	q := queue.Queue{3}
	q.Push(6)
	q.Push(8)
	fmt.Println(q)
	r := q.Pop()
	fmt.Println(r)
}
