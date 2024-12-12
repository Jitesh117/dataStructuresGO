package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array with duplicates",
			input:    []int{3, 1, 2, 1, 3},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "Array with one element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-3, -1, -2, -5, -4},
			expected: []int{-5, -4, -3, -2, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := BubbleSort(
				append([]int{}, tt.input...),
			)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v; want %v", tt.input, output, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := InsertionSort(
				append([]int{}, tt.input...),
			)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("InsertionSort(%v) = %v; want %v", tt.input, output, tt.expected)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := QuickSort(
				append([]int{}, tt.input...),
			)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("QuickSort(%v) = %v; want %v", tt.input, output, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MergeSort(
				append([]int{}, tt.input...),
			)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("MergeSort(%v) = %v; want %v", tt.input, output, tt.expected)
			}
		})
	}
}
