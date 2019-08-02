package array

// First Missing Positive: 41
// 缺失的第一个正数
// 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
// 示例:
// 输入: [3,4,-1,1]
// 输出: 2
//
// reference: https://leetcode.com/problems/first-missing-positive/
// 思路: 遍历一次数组把大于等于1的和小于数组大小的值放到原数组对应位置，然后再遍历一次数组查当前下标是否和值对应，
// 如果不对应那这个下标就是答案，否则遍历完都没出现那么答案就是数组长度加1。

// FirstMissingPositive 缺失的第一个正数
func FirstMissingPositive(nums []int) int {
	// 将数组的值交换到数组对应的位置
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] == i+1 || nums[i] > len(nums) {
			continue
		} else if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i-- // 替换之后还得从当前位置开始
		}
	}

	// 如果下标不对应就是答案
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1

	// 另一种想法:
	// 如果下标对应了, 就是匹配上了, 计数器加一, 最后返回匹配的数加一(正数应该在的位置)
	// i := 0
	// for i < len(nums) && nums[i] == i+1 {
	// 	i++
	// }
	// return i+1
}
