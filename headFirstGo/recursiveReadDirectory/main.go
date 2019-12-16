package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Give the folder name")
	}
	dirPath := os.Args[1]
	err := scanDirectory(dirPath)
	if err != nil {
		log.Fatal(err)
	}
}

func scanDirectory(path string) error {
	fmt.Println("The Dir Path is", path)
	contents, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, content := range contents {
		if content.IsDir() {
			filePath := filepath.Join(path, content.Name())
			err := scanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("The File Name is :", content.Name())
		}
	}
	return nil
}
