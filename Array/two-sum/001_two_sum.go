package array

// Two Sum : 001
// 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// reference: https://leetcode.com/problems/two-sum/
// 思路:
// 若结果的其中一个值为a, 另一个必为 target-a; 使用 map 减少复杂度

// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {
	result := make([]int, 2)
	if len(nums) < 2 {
		return result
	}
	tmpMap := make(map[int]int)
	for i, v := range nums {
		if _, ok := tmpMap[target-v]; ok {
			result[0] = tmpMap[target-v]
			result[1] = i
			break
		}
		tmpMap[v] = i // 便于获取 index
	}

	return result
}
