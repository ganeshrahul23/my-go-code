package main

import "fmt"

func sliceToString(slice []string) string {
	return fmt.Sprintf("%q", slice)
}

func addToMap(key []string, value string, m *map[string]string) {
	if *m == nil {
		*m = make(map[string]string)
	}
	(*m)[sliceToString(key)] = value
}

func getFromMap(key []string, m *map[string]string) (string, bool) {
	value, ok := (*m)[sliceToString(key)]
	return value, ok
}

func main() {
	var m map[string]string
	slice1 := []string{"A", "B", "C"}
	addToMap(slice1, "First 3", &m)
	slice1 = []string{"A"}
	addToMap(slice1, "First", &m)
	fmt.Println(getFromMap([]string{"A", "B"}, &m))
}
