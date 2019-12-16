package main

import "fmt"

func main() {
	ar1 := [5]int{1, 2, 3, 4, 5}
	ar2 := [...]int{4, 57, 8, 9, 0}
	ar3 := [...]int{0: 1, 9: 10}

	ar4Pointer := [5]*int{0: new(int)}
	ar4Pointer[1] = new(int)
	*ar4Pointer[1] = 2
	fmt.Println(ar1, ar2, ar3, ar4Pointer)

	array1 := [2]int{1, 2}
	var array2 [2]int
	array2 = array1

	array1[0] = 4

	fmt.Println("Array 1 :", array1, "Array 2 :", array2)
}
