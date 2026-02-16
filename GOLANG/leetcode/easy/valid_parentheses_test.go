package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "simple", s: "()", want: true},
		{name: "mixed", s: "()[]{}", want: true},
		{name: "invalid", s: "(]", want: false},
		{name: "unclosed", s: "([", want: false},
		{name: "wrong order", s: "([)]", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Fatalf("isValid() = %t, want %t", got, tt.want)
			}
		})
	}
}
