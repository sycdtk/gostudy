package main

const (
	MAX_SIZE = 10 //栈最大值
	EMPTY    = -1 //空值
)

type Element int64 //栈元素类型

type Stack struct {
	Array []Element //数组
	Top   int       //栈顶指指针
}

//判断栈是否满
func IsFull(s *Stack) bool {
	return s.Top >= (MAX_SIZE - 1)
}

//判断栈是否空
func IsEmpty(s *Stack) bool {
	return s.Top == -1
}

//压栈
func Push(s *Stack, data Element) {
	if !IsFull(s) {
		s.Top++
		s.Array[s.Top] = data
	}
}

func Pop(s *Stack) Element {
	if !IsEmpty(s) {
		data := s.Array[s.Top]
		s.Array[s.Top] = 0
		s.Top--
		return data
	}
	return EMPTY
}

//创建栈
func NewStack() *Stack {
	return &Stack{Array: make([]Element, MAX_SIZE), Top: -1}
}

func main() {
	s := NewStack()
	Push(s, 1000)
	Push(s, 56566)
	Push(s, 7876)
	Push(s, 235556)
	Pop(s)
	Pop(s)

}
