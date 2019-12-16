package main

import (
	"fmt"
	"github.com/ganeshrahul23/headFirstGo/automatedTesting/join"
)

func main() {
	phrases := []string{"apple", "orange"}
	fmt.Println(join.JoinWithCommas(phrases))
	phrases = append(phrases, "banana")
	fmt.Println(join.JoinWithCommas(phrases))

}
