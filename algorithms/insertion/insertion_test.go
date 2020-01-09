package main

import "testing"

func TestInsertion10(t *testing.T) {
	tests := []struct {
		name     string
		unsorted [10]int
		sorted   [10]int
	}{
		{
			name:     "Test One",
			unsorted: [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			sorted:   [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Test Two",
			unsorted: [10]int{0, 3, 7, 2, 1, 9, 8, 5, 6, 4},
			sorted:   [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dummy := insertion10(test.unsorted)
			if dummy != test.sorted {
				t.Error("Expected ", test.sorted, " but got", dummy)
			}
		})
	}
}

func TestInsertionReverse10(t *testing.T) {
	tests := []struct {
		name     string
		unsorted [10]int
		sorted   [10]int
	}{
		{
			name:     "Test Three",
			unsorted: [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			sorted:   [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name:     "Test Four",
			unsorted: [10]int{0, 3, 7, 2, 1, 9, 8, 5, 6, 4},
			sorted:   [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dummy := insertionReverse10(test.unsorted)
			if dummy != test.sorted {
				t.Error("Expected ", test.sorted, " but got", dummy)
			}
		})
	}
}
