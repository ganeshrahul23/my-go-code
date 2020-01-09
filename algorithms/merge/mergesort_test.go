package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		unsorted []int
		sorted   []int
	}{
		{
			name:     "Test One",
			unsorted: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			sorted:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Test Two",
			unsorted: []int{0, 3, 7, 2, 1, 9, 8, 5, 6, 4},
			sorted:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mergeSort(test.unsorted, 0, len(test.unsorted)-1)
			if !reflect.DeepEqual(test.unsorted, test.sorted) {
				t.Error("Expected ", test.sorted, " but got", test.unsorted)
			}
		})
	}
}
