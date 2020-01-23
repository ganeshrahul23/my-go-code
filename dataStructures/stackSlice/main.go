package main

import (
	"fmt"
	"github.com/ganeshrahul23/dataStructures/stackSlice/stack"
)

func main() {
	st := stack.NewStack()
	arr := []int{10, 45, 67, 43}

	for _, elem := range arr {
		fmt.Print(elem, "\t")
		st.Push(elem)
	}
	fmt.Println()

	for elem, err := st.Pop(); err == nil; elem, err = st.Pop() {
		fmt.Print(elem, "\t")
	}
	fmt.Println()
}
