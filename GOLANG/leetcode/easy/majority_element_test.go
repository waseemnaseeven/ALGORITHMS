package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "single", nums: []int{3}, want: 3},
		{name: "odd", nums: []int{3, 2, 3}, want: 3},
		{name: "even", nums: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Fatalf("majorityElement() = %d, want %d", got, tt.want)
			}
		})
	}
}
