package array

import (
	"sort"
)

// 3Sum Closest: 16
// 最接近的三数之和
// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，
// 使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
//
// reference:https://leetcode.com/problems/3sum-closest/
// 思路:
// 与 3Sum 类似, 排序之后遍历数组, a+b+c 最接近 target, 并返回 a+b+c之和, 即返回 a+b+c 与 target
// 绝对值最小的那个结果

// ThreeSumClosest 最接近的三数之和
func ThreeSumClosest(nums []int, target int) int {
	// sort slice
	sort.Ints(nums)
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return 0 - x
	}
	sum := 0
	res := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum = nums[i] + nums[start] + nums[end]
			if sum == target {
				return sum
			} else if sum < target {
				start++
			} else {
				end--
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
		}
	}

	return res
}
