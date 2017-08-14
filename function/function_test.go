package main

import (
	"testing"

	//	"fmt"
	"strings"

	"github.com/sycdtk/gotools/rpn"
	"github.com/sycdtk/gotools/stack"
)

func TestExp(t *testing.T) {
	al := 2
	exp := "@SUB(@ADD(1,@ADD(2,3)),@ADD(1,2))"
	exps := rpn.Parse(exp)
	s := stack.NewStack()

	for _, e := range strings.Split(exps, " ") {
		switch e {
		case "@ADD":
			stack.Push(s, ADD(s, al))
		case "@SUB":
			stack.Push(s, SUB(s, al))
		default:
			stack.Push(s, e)
		}
	}
	//		fmt.Println(stack.Pop(s))
}

func BenchmarkExp(b *testing.B) {

	exp := "@SUB(@ADD(1,@ADD(2,XX)),@ADD(1,2))"
	exps := rpn.Parse(exp)

	for i := 0; i < b.N; i++ { //use b.N for looping

		al := 2

		s := stack.NewStack()

		exps = strings.Replace(exps, "XX", "3", -1)

		for _, e := range strings.Split(exps, " ") {
			switch e {
			case "@ADD":
				stack.Push(s, ADD(s, al))
			case "@SUB":
				stack.Push(s, SUB(s, al))
			default:
				stack.Push(s, e)
			}
		}
	}
}

func BenchmarkExpC(b *testing.B) {
	exp := "@SUB(@ADD(1,@ADD(2,XX)),@ADD(1,2))"
	exps := rpn.Parse(exp)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			al := 2

			s := stack.NewStack()

			exps = strings.Replace(exps, "XX", "3", -1)

			for _, e := range strings.Split(exps, " ") {
				switch e {
				case "@ADD":
					stack.Push(s, ADD(s, al))
				case "@SUB":
					stack.Push(s, SUB(s, al))
				default:
					stack.Push(s, e)
				}
			}
		}
	})
}
