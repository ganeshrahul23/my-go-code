package main

import "fmt"

func main() {
	arr := []int{4, 0, 3, 9, 1}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
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
		L[i] = array[p+i]
	}

	for ; j < n2; j++ {
		R[j] = array[q+j+1]
	}
	for i, j = 0, 0; i < n1 && j < n2; k++ {
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
