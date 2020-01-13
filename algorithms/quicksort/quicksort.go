package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{4, 0, 3, 9, 1, 11, 45, -1, -45, 100, 103}
	randomizedQuicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func randomizedPartition(array []int, p, q int) int {
	i := rand.Intn(q-p) + p
	array[i], array[q] = array[q], array[i]
	return partition(array, p, q)
}

func partition(array []int, p, q int) int {
	pivot := array[p]
	i := p
	for j := p + 1; j <= q; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[p], array[i] = array[i], array[p]
	return i
}

func quicksort(array []int, p, q int) {
	if p < q {
		r := partition(array, p, q)
		quicksort(array, p, r-1)
		quicksort(array, r+1, q)
	}
}

func randomizedQuicksort(array []int, p, q int) {
	if p < q {
		r := randomizedPartition(array, p, q)
		randomizedQuicksort(array, p, r-1)
		randomizedQuicksort(array, r+1, q)
	}
}
