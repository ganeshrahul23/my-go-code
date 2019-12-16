package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Give the folder name")
	}
	dirName := os.Args[1]

	contents, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, content := range contents {
		if content.IsDir() {
			fmt.Println("Directory is ", content.Name())
		} else {
			fmt.Println("File is ", content.Name())
		}
	}

}
