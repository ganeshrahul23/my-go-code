package main

import (
	"fmt"
	"github.com/ganeshrahul23/dataStructures/queueSlice/queue"
)

func main() {
	qu := queue.NewQueue()
	arr := []int{10, 45, 67, 43}

	for _, elem := range arr {
		fmt.Print(elem, "\t")
		qu.Enqueue(elem)
	}
	fmt.Println()

	for elem, err := qu.Dequeue(); err == nil; elem, err = qu.Dequeue() {
		fmt.Print(elem, "\t")
	}
	fmt.Println()
}
