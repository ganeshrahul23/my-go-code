package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("file", "foo", "fileName")
	noOfLines := flag.Int("n", -1, "No of Files")
	flag.Parse()
	if *fileName == "foo" {
		return
	}
	file, _ := os.Open(*fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
		if *noOfLines != -1 && i == *noOfLines {
			break
		}
	}
}
