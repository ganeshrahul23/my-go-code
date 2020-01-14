package main

import (
	"reflect"
	"testing"
)

func TestLeft(t *testing.T) {
	tests := []struct {
		name string
		root int
		left int
	}{
		{
			"Left 0",
			0,
			1,
		},
		{
			"Left 1",
			1,
			3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			left := Left(test.root)
			if left != test.left {
				t.Error(
					"For", test.root,
					"expected", test.left,
					"got", left,
				)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name  string
		root  int
		right int
	}{
		{
			"Right 0",
			0,
			2,
		},
		{
			"Right 1",
			1,
			4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			right := Right(test.root)
			if right != test.right {
				t.Error(
					"For", test.root,
					"expected", test.right,
					"got", right,
				)
			}
		})
	}
}

func TestMaxHeapify(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
	}{
		{
			"Test One",
			[]int{0, 45, 34, 99, 98, 23, 67},
			[]int{0, 45, 67, 99, 98, 23, 34},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			MaxHeapify(test.array, 2, len(test.array))
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Error("Expected ", test.expected, "but got", test.array)
			}
		})
	}
}

func TestBuildMaxHeap(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		max   int
	}{
		{
			"Test One",
			[]int{0, 45, 34, 99, 98, 23, 67},
			99,
		},
		{
			"Test Two",
			[]int{100, 45, 34, 103, 98, 23, 101},
			103,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			BuildMaxHeap(test.array)
			if test.array[0] != test.max {
				t.Error(
					"expected", test.max,
					"got", test.array[0],
				)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
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
			HeapSort(test.unsorted)
			if !reflect.DeepEqual(test.unsorted, test.sorted) {
				t.Error("Expected ", test.sorted, " but got", test.unsorted)
			}
		})
	}
}
