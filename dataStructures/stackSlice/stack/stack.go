package stack

import "errors"

//A Int Stack
//Use @NewStack to create a empty stack
type stack struct {
	slice []int
}

//Function to Initialize
//and return a int stack
func NewStack() *stack {
	return &stack{slice: make([]int, 0)}
}

func (s *stack) Push(elems ...int) {
	s.slice = append(s.slice, elems...)
}

func (s *stack) Pop() (int, error) {
	l := len(s.slice) - 1
	if l == -1 {
		return 0, errors.New("Stack Underflow")
	}
	elem := s.slice[l]
	s.slice = s.slice[:l]
	return elem, nil
}

func (s *stack) Peek() (int, error) {
	l := len(s.slice) - 1
	if l == -1 {
		return 0, errors.New("Empty Stack")
	}
	return s.slice[l], nil
}
