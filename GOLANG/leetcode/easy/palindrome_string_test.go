package main

import "testing"

func TestIsPalindromeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "panama", s: "A man, a plan, a canal: Panama", want: true},
		{name: "not palindrome", s: "race a car", want: false},
		{name: "blank", s: " ", want: true},
		{name: "mixed", s: "No 'x' in Nixon", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeString(tt.s); got != tt.want {
				t.Fatalf("isPalindromeString() = %t, want %t", got, tt.want)
			}
		})
	}
}
