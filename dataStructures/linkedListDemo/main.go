package main

import (
	"github.com/ganeshrahul23/dataStructures/linkedList"
)

func main() {
	ll := linkedList.LinkedList{}
	ll.AddFirst(100)
	ll.AddFirst(200)
	ll.AddFirst(300)
	ll.Iterate()
	ll.ReverseList()
	ll.Iterate()
}
