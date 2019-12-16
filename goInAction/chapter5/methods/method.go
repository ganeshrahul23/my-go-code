package main

import "fmt"

type Temperature float64

func (t Temperature) Print0() {
	fmt.Println("The Temperature 1 is : ", t)
}

func (t *Temperature) Print1() {
	fmt.Println("The Temperature 2 is : ", *t)
}
func main() {
	t := Temperature(23.0)
	t.Print0()
	t.Print1()
	tp := &t
	tp.Print0()

	t1 := new(Temperature)
	*t1 = 34.0
	t1.Print0()
	t1.Print1()
	Temperature.Print0(Temperature(34.0))
	(*Temperature).Print1(new(Temperature))
}
