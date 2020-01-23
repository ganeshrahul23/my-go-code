package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//insertion(arr)
	//selection(arr)
	//bubble(arr)
	//mergeSort(arr, 0, len(arr) - 1)
	//quicksort(arr, 0, len(arr) - 1, true)
	fmt.Println(arr)
}

func bubble(array []int) {
	l := len(array)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func selection(array []int) {
	l := len(array)
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}

func insertion(array []int) {
	l := len(array)
	for j := 1; j < l; j++ {
		i := j - 1
		key := array[j]
		for ; i >= 0 && array[i] > key; i-- {
			array[i+1] = array[i]
		}
		array[i+1] = key
	}
}

func mergeSort(array []int, p int, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(array, p, q)
		mergeSort(array, q+1, r)
		merge(array, p, q, r)
	}
}

func merge(array []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q

	L := make([]int, n1)
	R := make([]int, n2)

	i, j, k := 0, 0, p

	for ; i < n1; i++ {
		L[i] = array[i+p]
	}

	for ; j < n2; j++ {
		R[j] = array[j+q+1]
	}

	i, j = 0, 0

	for ; i < n1 && j < n2; k++ {
		if L[i] <= R[j] {
			array[k] = L[i]
			i++
		} else {
			array[k] = R[j]
			j++
		}
	}
	for ; i < n1; k++ {
		array[k] = L[i]
		i++
	}
	for ; j < n2; k++ {
		array[k] = R[j]
		j++
	}
}

func quicksort(array []int, p, q int, random bool) {
	if p < q {
		r := partition(array, p, q, random)
		quicksort(array, p, r-1, random)
		quicksort(array, r+1, q, random)
	}
}

func partition(array []int, p, q int, random bool) int {
	if random {
		randIndex := rand.Intn(q-p) + p
		array[randIndex], array[q] = array[q], array[randIndex]
	}
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
