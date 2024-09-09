package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) Push(ele int) {
	s.elements = append(s.elements, ele)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		fmt.Println("栈为空")
		return -1
	}
	top := s.elements[len(s.elements)-1]

	// 移除栈顶元素
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		fmt.Println("栈为空")
		return -1
	}
	return s.elements[len(s.elements)-1]
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Printf("pop-> %d\n", stack.Pop())
	fmt.Printf("peek-> %d\n", stack.Peek())
}
