package main

import "github.com/ganeshrahul23/dataStructures/linkedList"

func main() {
	//ll := linkedList.LinkedList{}
	//ll.AddFirst(100)
	//ll.AddFirst(200)
	//ll.AddFirst(300)
	//ll.Iterate()
	//ll.ReverseList()
	//ll.Iterate()

	ll1 := linkedList.LinkedList{}
	ll1.AddLast(1)
	ll1.AddLast(3)
	ll1.AddLast(5)
	ll1.AddLast(7)

	ll2 := linkedList.LinkedList{}
	ll2.AddLast(2)
	ll2.AddLast(4)
	ll2.AddLast(6)

	ll1.MergeSortedList(ll2)
	ll1.Iterate()
}
