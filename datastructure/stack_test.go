package datastructure

import (
	"fmt"
	"testing"
)

func TestStackInt(t *testing.T) {
	stack1 := NewStackInt(10)
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)
	stack1.Push(4)

	for stack1.Size() != 0 {
		fmt.Println(stack1.Pop())
	}
}

func TestStackBT(t *testing.T) {
	stack := NewStackBT(10)
	stack.Push(&BTNode{data: 1})
	stack.Push(&BTNode{data: 2})
	stack.Push(&BTNode{data: 3})

	for stack.Size() != 0 {
		fmt.Println(stack.Pop())
	}
}
