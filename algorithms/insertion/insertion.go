package main

import "fmt"

func main() {
	fmt.Println(insertion([10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}

func insertion(array [10]int) [10]int {
	l := len(array)
	for j := 1; j < l; j++ {
		key := array[j]
		i := j - 1
		for ; i >= 0 && array[i] > key; i = i - 1 {
			array[i+1] = array[i]
		}
		array[i+1] = key
	}
	return array
}
