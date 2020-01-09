package main

import "fmt"

func main() {
	arr := []int{4, 0, 3, 9, 1}
	h := Heap{}
	h.HeapSort(arr)

	fmt.Println(arr)
}

type Heap struct {
}

func (*Heap) Left(root int) int {
	return (root * 2) + 1
}

func (*Heap) Right(root int) int {
	return (root * 2) + 2
}

func (h *Heap) BuildHeap(array []int) {
	l := len(array)
	for i := l / 2; i >= 0; i-- {
		h.Heapify(array, i, l)
	}
}

func (h *Heap) Heapify(array []int, root, length int) {
	var max = root
	var l, r = h.Left(root), h.Right(root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		h.Heapify(array, max, length)
	}
}

func (heap *Heap) RemoveTop(array []int, length int) {
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	heap.Heapify(array, 0, lastIndex)
}

func (heap *Heap) HeapSort(array []int) {
	heap.BuildHeap(array)

	for length := len(array); length > 1; length-- {
		heap.RemoveTop(array, length)
	}
}
