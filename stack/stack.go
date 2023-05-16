package stack

import "fmt"

type Stack struct {
	Elements []interface{}
	Top      int64
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(element interface{}) {
	s.Elements = append(s.Elements, element)
	s.Top++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}
	s.Top--
	ans := s.Elements[s.Top]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return ans, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}
	return s.Elements[s.Top-1], nil
}

func (s *Stack) Len() int64 {
	return s.Top
}

func (s *Stack) IsEmpty() bool {
	return s.Top == 0
}
