package linkedList

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

func (l *LinkedList) AddLast(data int) {
	if l.initHead(data) {
		return
	}
	l.findTail().next = &node{data: data, next: nil}
}

func (l *LinkedList) AddFirst(data int) {
	if l.initHead(data) {
		return
	}
	l.head = &node{data: data, next: l.head}
}

func (l *LinkedList) findTail() *node {
	tempNode := l.head
	for tempNode.next != nil {
		tempNode = tempNode.next
	}
	return tempNode
}

func (l *LinkedList) initHead(data int) bool {
	if l.head != nil {
		return false
	}
	l.head = &node{data: data, next: nil}
	return true
}

func (l *LinkedList) Iterate() {
	tempNode := l.head
	for tempNode.next != nil {
		fmt.Print(tempNode.data, "\t")
		tempNode = tempNode.next
	}
	fmt.Println(tempNode.data)
}

func (l *LinkedList) Len() int {
	if l.head == nil {
		return 0
	}
	currentNode := l.head
	length := 1
	for ; currentNode.next != nil; length++ {
		currentNode = currentNode.next
	}
	return length
}

func (l *LinkedList) Find(data int) int {
	node := l.head
	index := -1
	for ; node.next != nil && node.data != data; node = node.next {
		index++
	}
	if node.data == data {
		return index + 1
	}
	return -1
}

func (l *LinkedList) RemoveTail() {
	node := l.head
	for node.next.next != nil {
		node = node.next
	}
	node.next = nil
}

func (l *LinkedList) ReverseList() {
	if l.head == nil || l.head.next == nil {
		return
	}
	var prevNode, nextNode, currentNode *node
	currentNode = l.head
	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	l.head = prevNode
}

func (l *LinkedList) MergeSortedList(list LinkedList) {

}
