package main

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{name: "found at start", haystack: "sadbutsad", needle: "sad", want: 0},
		{name: "found later", haystack: "ksadbutsad", needle: "sad", want: 1},
		{name: "not found", haystack: "leetcode", needle: "leeto", want: -1},
		{name: "empty needle", haystack: "abc", needle: "", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Fatalf("strStr() = %d, want %d", got, tt.want)
			}
		})
	}
}
