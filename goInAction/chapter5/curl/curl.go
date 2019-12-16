package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func init() {
	if len(os.Args) != 2 {
		switch osName := runtime.GOOS; osName {
		case "windows":
			fmt.Println("Usage : curl.exe <url>")
		default:
			fmt.Println("Usage : ./curl <url>")
		}
		os.Exit(-1)
	}
}
func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, resp.Body)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
