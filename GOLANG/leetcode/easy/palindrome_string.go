package main

import "unicode"

func isPalindromeString(s string) bool {
	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	rev := make([]rune, len(cleaned))
	copy(rev, cleaned)

	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}

	for i := range cleaned {
		if cleaned[i] != rev[i] {
			return false
		}
	}
	return true
}
