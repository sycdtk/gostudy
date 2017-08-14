package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sycdtk/gotools/rpn"
	"github.com/sycdtk/gotools/stack"
)

func main() {
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
	fmt.Println(stack.Pop(s))
}

func ADD(s *stack.Stack, n int) string {
	args := stack.NewStack()

	for i := 0; i < n; i++ {
		stack.Push(args, stack.Pop(s))
	}

	value := add(GetFloat64(stack.Pop(args)), GetFloat64(stack.Pop(args)))

	return strconv.FormatFloat(value, 'f', -1, 64)
}

func SUB(s *stack.Stack, n int) string {
	args := stack.NewStack()

	for i := 0; i < n; i++ {
		stack.Push(args, stack.Pop(s))
	}

	value := sub(GetFloat64(stack.Pop(args)), GetFloat64(stack.Pop(args)))

	return strconv.FormatFloat(value, 'f', -1, 64)
}

func add(v1, v2 float64) float64 {
	return v1 + v2
}

func sub(v1, v2 float64) float64 {
	return v1 - v2
}

func GetFloat64(v string) float64 {
	v22, _ := strconv.ParseFloat(v, 64)
	return v22
}
