package main

/*
	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	You can return the answer in any order.

	Example 1:
	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	
	Example 2:
	Input: nums = [3,2,4], target = 6
	Output: [1,2]
	
	Example 3:
	Input: nums = [3,3], target = 6
	Output: [0,1]
*/

// 33 ms
func twoSum3(nums []int, target int) []int {
    ret := make([]int, 0, 2)
    n := len(nums)
    for i := 0; i < n ; i++{
        for j := i + 1; j < n; j++ {
            if nums[i] + nums[j] == target {
                ret = append(ret, i, j)
            }
        }
    } 
    return ret
}

// 23 ms
func twoSum2(nums []int, target int) []int {
    n := len(nums)
    for i := 0; i < n ; i++{
        for j := i + 1; j < n; j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    } 
    return nil
}

func twoSum(nums []int, target int) []int {
	// i = 0, v = 2
	// complement := 9 - 2 = 7
	// → “If I see a 7 somewhere, 2 + 7 = 9.”
    seen := make(map[int]int)

    for i, v := range nums {
        need := target - v

        if j, ok := seen[need]; ok {
            return []int{j, i}
        }

        seen[v] = i
    }

    return nil
}