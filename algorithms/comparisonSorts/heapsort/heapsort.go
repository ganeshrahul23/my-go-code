package main

import "fmt"

func main() {
	arr := []int{4, 0, 3, 9, 1}
	HeapSort(arr)
	fmt.Println(arr)
}

func Left(index int) int {
	return (index * 2) + 1
}

func Right(index int) int {
	return (index * 2) + 2
}

func Parent(index int) int {
	return ((index + 1) / 2) - 1
}

func HeapSort(array []int) {
	BuildMaxHeap(array)
	for length := len(array); length > 1; length-- {
		HeapExtractMax(array, length)
	}
}

func BuildMaxHeap(array []int) {
	l := len(array)
	for i := l / 2; i >= 0; i-- {
		MaxHeapify(array, i, l)
	}
}

func HeapExtractMax(array []int, length int) {
	lastIndex := length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	MaxHeapify(array, 0, lastIndex)
}

func MaxHeapify(array []int, index, length int) {
	max := index
	l, r := Left(index), Right(index)

	if l < length && array[l] > array[max] {
		max = l
	}
	if r < length && array[r] > array[max] {
		max = r
	}
	if max != index {
		array[index], array[max] = array[max], array[index]
		MaxHeapify(array, max, length)
	}
}
