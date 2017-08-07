package main

import (
	"testing"
)

func TestStack(t *testing.T) {

	for i := 0; i < 10000000; i++ {
		//		fmt.Println(i)
		stack := NewStack()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
		stack.Pop()
		stack.Pop()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Pop()
		stack.Pop()
	}

}
