package main

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, x := range nums {
		if count == 0 {
			candidate = x
		}
		if x == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
