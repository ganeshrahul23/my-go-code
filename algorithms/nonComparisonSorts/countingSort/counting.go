package main

import "fmt"

func main() {
	fmt.Println(countingSort([]int{1, 1, 2, 3, 5, 0, 3, 4}, 5))
}

func countingSort(array []int, k int) []int {
	C := make([]int, k+1)
	l := len(array)
	B := make([]int, l)

	for _, v := range array {
		C[v] = C[v] + 1
	}
	for j := 1; j <= k; j++ {
		C[j] = C[j] + C[j-1]
	}

	//This loop makes the counting sort not stable
	//for _, v := range array {
	//	C[v] = C[v] - 1
	//	B[C[v]] = v
	//}

	for i := l - 1; i >= 0; i-- {
		C[array[i]] = C[array[i]] - 1
		B[C[array[i]]] = array[i]
	}

	return B
}
