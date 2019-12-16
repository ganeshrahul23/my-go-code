package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"`
}

var movies = []Movie{
	{"Avengers", 2012, true},
	{"Pan", 2014, false},
}

func main() {
	data, _ := json.MarshalIndent(movies, "", "	")
	fmt.Printf("%s\n", data)

	var movieTitles []struct{ Title string }

	json.Unmarshal(data, &movieTitles)

	fmt.Println(movieTitles)

}
