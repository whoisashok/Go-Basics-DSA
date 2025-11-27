package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"found middle", []int{1, 3, 5, 7, 9}, 7, 3},
		{"found first", []int{1, 3, 5, 7, 9}, 1, 0},
		{"found last", []int{1, 3, 5, 7, 9}, 9, 4},
		{"not found", []int{1, 3, 5, 7, 9}, 2, -1},
		{"empty", []int{}, 1, -1},
		{"single match", []int{42}, 42, 0},
		{"single not match", []int{42}, 7, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := BinarySearch(tc.arr, tc.target)
			if got != tc.want {
				t.Fatalf("BinarySearch(%v, %d) = %d; want %d", tc.arr, tc.target, got, tc.want)
			}
		})
	}
}
