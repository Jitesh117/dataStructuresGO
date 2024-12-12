package searching

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []int{4, 2, 7, 1, 9, 3}
	tests := []struct {
		name   string
		target int
		want   int
	}{
		{"Found at index 0", 4, 0},
		{"Found at index 4", 9, 4},
		{"Not found", 10, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LinearSearch(arr, tt.target)
			if got != tt.want {
				t.Errorf("LinearSearch() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{4, 2, 7, 1, 9, 3}
	tests := []struct {
		name   string
		target int
		want   int
	}{
		{"Found at index 0", 4, 0},
		{"Found at index 4", 9, 4},
		{"Not found", 10, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(arr, tt.target)
			if got != tt.want {
				t.Errorf("BinarySearch() = %d, want %d", got, tt.want)
			}
		})
	}
}
