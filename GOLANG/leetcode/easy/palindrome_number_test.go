package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{name: "palindrome", x: 121, want: true},
		{name: "negative", x: -121, want: false},
		{name: "ends with zero", x: 10, want: false},
		{name: "single digit", x: 7, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Fatalf("isPalindrome() = %t, want %t", got, tt.want)
			}
		})
	}
}
