package leetcode

//先对nums进行非降序排序，然后以1,n,2,3,n-1,4,5,n-2的顺序进行输出新的顺序
func wiggleSort(nums []int) {
	nums = sortN2(nums)
	tem := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i%3 == 0 {
			tem[i] = tem[i/3*2]
		}
		if i%3 == 1 {
			tem[i] = tem[len(nums)-1-i/3]
		}
		if i%3 == 2 {
			tem[i] = tem[i/3*2+1]
		}
	}
	nums = tem
}

/**
Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
给定一个整数数组 nums ，对其进行重新排序，以便 nums[0] < nums[1] > nums[2] < nums[3]... .

You may assume the input array always has a valid answer.
您可以假设输入数组始终具有有效的答案。



Example 1: 示例 1：

Input: nums = [1,5,1,1,6,4]
Output: [1,6,1,5,1,4]
Explanation: [1,4,1,5,1,6] is also accepted.
Example 2: 示例 2：

Input: nums = [1,3,2,2,3,1]
Output: [2,3,1,3,1,2]


Constraints: 约束：

1 <= nums.length <= 5 * 104
0 <= nums[i] <= 5000
It is guaranteed that there will be an answer for the given input nums.
保证给定输入 nums 会有答案。


Follow Up: 跟进： Can you do it in O(n) time and/or in-place with O(1) extra space?
你能 O(n) 及时和/或用额外的空间就 O(1) 地做吗？
*/
