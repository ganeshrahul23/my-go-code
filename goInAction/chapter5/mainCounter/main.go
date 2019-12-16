package main

import (
	"fmt"
	"github.com/ganeshrahul23/goInAction/chapter5/counters"
)

func main() {
	c := counters.New(10)
	fmt.Printf("Type is %T\tValue is %v\n", c, c)
}
