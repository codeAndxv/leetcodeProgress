package leetcode

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		j := nums[i]
		for j > 0 && j <= len(nums) && nums[j-1] != j {
			tem := nums[j-1]
			nums[j-1] = j
			j = tem
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	if nums[len(nums)-1] == len(nums) {
		return len(nums) + 1
	} else {
		return len(nums)
	}
}

/**
41. First Missing Positive
41.第一个缺失的阳性
提示
困难
1.9K
相关企业
Given an unsorted integer array nums, return the smallest missing positive integer.
给定一个未排序的整数数组 nums ，返回最小的缺失正整数。

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.
您必须实现按 O(n) 时运行并使用辅助空间的 O(1) 算法。



Example 1: 示例 1：

Input: nums = [1,2,0]
Output: 3
Explanation: The numbers in the range [1,2] are all in the array.
Example 2: 示例 2：

Input: nums = [3,4,-1,1]
Output: 2
Explanation: 1 is in the array but 2 is missing.
Example 3: 例3：

Input: nums = [7,8,9,11,12]
Output: 1
Explanation: The smallest positive integer 1 is missing.


Constraints: 约束：

1 <= nums.length <= 10^5
-23^1 <= nums[i] <= 23^1 - 1*/
