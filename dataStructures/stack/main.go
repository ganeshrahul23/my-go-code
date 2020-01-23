package main

import (
	"fmt"
	"github.com/ganeshrahul23/dataStructures/stack/stack"
	"log"
)

func main() {
	s := stack.Stack{}

	err := s.Push(100)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Push(102)
	if err != nil {
		log.Fatal(err)
	}

	s.PrintArray()

	elem, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(elem)

	s.PrintArray()

	err = s.Push(102)
	if err != nil {
		log.Fatal(err)
	}

	elem, err = s.Peek()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(elem)

	s.Pop()
	s.Pop()
	pop, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pop)
	s.PrintArray()
}
