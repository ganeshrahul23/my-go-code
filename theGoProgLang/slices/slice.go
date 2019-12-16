package main

import "fmt"

func main() {
	array := [10]int{}
	for i := range array {
		array[i] = i
	}

	slice1 := array[:5]

	slice2 := array[:5]

	slice3 := array[5:]

	fmt.Println(slice1, slice2, slice3)
	//fmt.Println(&slice1, &slice2, &slice3)

	slice1 = append(slice1, 10, 11, 12, 13, 14)
	slice1 = append(slice1, 15, 16)
	fmt.Println(slice1, slice2, slice3)
	//fmt.Println(&slice1, &slice2, &slice3)
	fmt.Println(array)

	fmt.Println("Append Visualization")

	slice1 = make([]int, 2)
	slice1[0] = -2
	slice1[1] = -1
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
		fmt.Printf("cap : %d\tslice : %v\n", cap(slice1), slice1)
	}
}
