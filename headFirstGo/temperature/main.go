package main

import (
	"fmt"
	"github.com/ganeshrahul23/headFirstGo/keyboard"
)

func main() {
	temperature, err := keyboard.GetFloat("Enter the Temperature :")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The Temperature is", temperature)
}
