package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.String())
	fmt.Println(v)

}
