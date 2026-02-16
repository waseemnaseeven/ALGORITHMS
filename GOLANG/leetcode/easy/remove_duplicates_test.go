package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "basic", nums: []int{1, 1, 2}, k: 2, want: []int{1, 2}},
		{name: "longer", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, k: 5, want: []int{0, 1, 2, 3, 4}},
		{name: "single", nums: []int{1}, k: 1, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			if got := removeDuplicates(nums); got != tt.k {
				t.Fatalf("removeDuplicates() = %d, want %d", got, tt.k)
			}
			if !reflect.DeepEqual(nums[:tt.k], tt.want) {
				t.Fatalf("removeDuplicates() nums[:k] = %v, want %v", nums[:tt.k], tt.want)
			}
		})
	}
}
