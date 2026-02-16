package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "profit", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "no profit", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "single", prices: []int{1}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Fatalf("maxProfit() = %d, want %d", got, tt.want)
			}
		})
	}
}
