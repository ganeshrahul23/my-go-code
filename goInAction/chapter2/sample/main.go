package main

import (
	_ "github.com/ganeshrahul23/goInAction/chapter2/sample/matchers"
	"github.com/ganeshrahul23/goInAction/chapter2/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("pravin")
}
