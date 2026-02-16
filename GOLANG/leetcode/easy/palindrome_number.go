package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	rev := 0

	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}

	return original == rev
}
