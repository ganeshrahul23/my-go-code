package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{0, 1, 9, 3, 4, 5, 6, 7, 90, 6, 70}
	fmt.Println(radixSort(arr))
}

func radixSort(array []int) []int {
	input := array
	maxLength := strconv.FormatInt(findMax(input), 2)
	for i := range maxLength {
		input = countingSort(input, i)
		fmt.Println(input)
	}
	return input
}

func countingSort(array []int, bit int) []int {
	C := []int{0, 0}

	l := len(array)
	B := make([]int, l)

	for _, v := range array {
		radix := (v >> bit) & 1
		C[radix] = C[radix] + 1
	}

	C[1] = C[1] + C[0]

	for i := l - 1; i >= 0; i-- {
		radix := (array[i] >> bit) & 1
		C[radix] = C[radix] - 1
		B[C[radix]] = array[i]
	}
	return B
}

func findMax(array []int) int64 {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return int64(max)
}
