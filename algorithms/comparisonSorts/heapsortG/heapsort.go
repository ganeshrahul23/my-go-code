package main

import "fmt"

func main() {
	arr := []int{7, 5, 2, 4, 3, 9}
	HeapSort(arr)
	fmt.Println(arr)
}

func Left(i int) int {
	return (i << 2) + 1
}

func Right(i int) int {
	return (i << 2) + 2
}

func MaxHeapify(array []int, i int, length int) {
	l, r := Left(i), Right(i)
	largest := i

	if l < length && array[l] > array[i] {
		largest = l
	}
	if r < length && array[r] > array[largest] {
		largest = r
	}

	if largest != i {
		array[largest], array[i] = array[i], array[largest]
		MaxHeapify(array, largest, length)
	}
}

func BuildMaxHeap(array []int) {
	length := len(array)
	for i := length / 2; i >= 0; i-- {
		MaxHeapify(array, i, length)
	}
}

func HeapSort(array []int) {
	BuildMaxHeap(array)
	length := len(array)
	for i := length - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		MaxHeapify(array, 0, i)
	}
}
