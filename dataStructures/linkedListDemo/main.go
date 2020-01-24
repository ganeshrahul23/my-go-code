package main

import (
	"fmt"
	"github.com/ganeshrahul23/dataStructures/linkedList"
)

func main() {
	ll := linkedList.LinkedList{}
	ll.AddFirst(100)
	ll.AddFirst(200)
	ll.AddFirst(300)
	ll.Iterate()
	ll.RemoveTail()
	ll.Iterate()
	fmt.Println(ll.Find(200))
	fmt.Println(ll.Find(300))
	fmt.Println(ll.Find(100))
	fmt.Println(ll.Find(1))
}
