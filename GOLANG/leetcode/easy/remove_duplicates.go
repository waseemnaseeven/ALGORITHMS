package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	k := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
