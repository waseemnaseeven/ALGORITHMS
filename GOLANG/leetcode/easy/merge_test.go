package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "merge basic",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "merge empty nums2",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "merge empty nums1",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums1 := append([]int(nil), tt.nums1...)
			merge(nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(nums1, tt.want) {
				t.Fatalf("merge() = %v, want %v", nums1, tt.want)
			}
		})
	}
}
