package main

import "fmt"

func main() {
	fmt.Println(insertion10([10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}

func insertion10(array [10]int) [10]int {
	for j := 1; j < 10; j++ {
		key := array[j]
		i := j - 1
		for ; i >= 0 && array[i] > key; i-- {
			array[i+1] = array[i]
		}
		array[i+1] = key
	}
	return array
}

func insertionReverse10(array [10]int) [10]int {
	for j := 1; j < 10; j++ {
		key := array[j]
		i := j - 1
		for ; i >= 0 && array[i] < key; i-- {
			array[i+1] = array[i]
		}
		array[i+1] = key
	}
	return array
}
