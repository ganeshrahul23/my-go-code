package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	array    [100]int
	top      int
	max_size int
	init     bool
}

func (s *Stack) newStack() {
	if !s.init {
		s.top = -1
		s.max_size = 100
		s.init = true
	}
}

func (s *Stack) Push(elem int) error {
	s.newStack()
	if s.top < s.max_size-1 {
		s.top++
		s.array[s.top] = elem
		return nil
	}
	return errors.New("Stack Overflow")
}

func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return 0, errors.New("Stack Underflow")
	}
	s.top--
	return s.array[s.top+1], nil
}

func (s *Stack) PrintArray() {
	if s.top == -1 {
		return
	}
	for i := 0; i <= s.top; i++ {
		fmt.Print(s.array[i], "\t")
	}
	fmt.Println()
}

func (s *Stack) Peek() (int, error) {
	if s.top == -1 {
		return 0, errors.New("Stack Underflow")
	}
	return s.array[s.top], nil
}
