package datastructure

import (
	"fmt"
	"testing"
)

func TestNormal(t *testing.T) {
	stack1 := NewStackInt(10)
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)
	stack1.Push(4)

	for stack1.Size() != 0 {
		fmt.Println(stack1.Pop())
	}
}
