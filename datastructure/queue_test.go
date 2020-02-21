package datastructure

import (
	"fmt"
	"testing"
)

func TestQueueBT(t *testing.T) {
	queue := NewQueueBT(10)
	queue.Enqueue(&BTNode{Data: 1})
	queue.Enqueue(&BTNode{Data: 2})
	queue.Enqueue(&BTNode{Data: 3})

	for queue.Size() != 0 {
		fmt.Println(queue.Dequeue())
	}

	fmt.Printf("final size:%d\n", queue.Size())
	fmt.Println(queue)
}
