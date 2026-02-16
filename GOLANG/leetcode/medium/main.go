package main 

import (
	"fmt"
)

/*

	Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

	Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

	Return k after placing the final result in the first k slots of nums.

	Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.


	Example 1:
	Input: nums = [1,1,1,2,2,3]
	Output: 5, nums = [1,1,2,2,3,_]
	Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).

	Example 2:
	Input: nums = [0,0,1,1,1,1,2,3,3]
	Output: 7, nums = [0,0,1,1,2,3,3,_,_]
	Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).

*/

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	k := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

/*
	Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

	Example 1:

	Input: nums = [1,2,3,4,5,6,7], k = 3
	Output: [5,6,7,1,2,3,4]
	Explanation:
	rotate 1 steps to the right: [7,1,2,3,4,5,6]
	rotate 2 steps to the right: [6,7,1,2,3,4,5]
	rotate 3 steps to the right: [5,6,7,1,2,3,4]

	Example 2:

	Input: nums = [-1,-100,3,99], k = 2
	Output: [3,99,-1,-100]
	Explanation: 
	rotate 1 steps to the right: [99,-1,-100,3]
	rotate 2 steps to the right: [3,99,-1,-100]

*/

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	len := len(nums)
	if len == 0 {
		return
	}

	k %= len

	// reverse tout
	reverse(nums, 0, len -1)
	// reverse les k premiers
	reverse(nums, 0, k-1)
	// reverse le reste
	reverse(nums, k, len-1)
}


func main() {
	fmt.Println(removeDuplicates([...]int{1,1,1,5,2,3,4,8,5,4}))
}