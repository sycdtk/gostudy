package main

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		fmt.Println(i)
		s := NewStack()
		Push(s, 1000)
		Push(s, 56566)
		Push(s, 34)
		Push(s, 99)
		Pop(s)
		Pop(s)
		Push(s, 756)
		Push(s, 889)
		Push(s, 34)
		Push(s, 99)
		Pop(s)
		Pop(s)
		Pop(s)
		Pop(s)
		Pop(s)
		Pop(s)

	}

}
