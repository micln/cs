package cs

import (
	"fmt"
)

type stack struct {
	data []interface{}
	top  int
}

func (s *stack) Empty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *stack) Top() interface{} {
	return s.data[s.top]
}

func (s *stack) String() string {
	r := ""
	fmt.Sprintf(r, "%s", s.data[s.top])
	return r
}

func (s *stack) Pop() {
	if !s.Empty() {
		s.top--
	}
}

func (s *stack) Push(t interface{}) {
	s.top++
	s.data = append(s.data, t)
}

func Stack() *stack {
	var s *stack
	return s
}
