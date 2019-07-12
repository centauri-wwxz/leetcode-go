package array

import (
	"sort"
)

// 3Sum: 15
// 三数之和
// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// reference: https://leetcode.com/problems/3sum
// 思路:
// 类似 2Sum, 得出: a + b = -c, 将数组排序后, 从首尾标记 i、k , 在 i、k 之间遍历数组, 下标为 j, 如果
// nums[j]+ nums[k] = -nums[i], 则可将 {i,j,k}对应的值 放入结果集. 注意去重.

// ThreeSum 三数之和
func ThreeSum(nums []int) [][]int {
	// sort slice
	sort.Ints(nums)
	i, j, k := 0, 0, 0

	target := 0
	res := make([][]int, 0)
	for i = 0; i < len(nums)-2; i++ {
		// 如果最左边的大于或者最右边的小于 0, 无解, 跳出循环
		if nums[i] > 0 || nums[len(nums)-1] < 0 {
			break
		}
		// i 去重
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			j = i + 1
			k = len(nums) - 1
			target = 0 - nums[i]

			for j < k {
				if nums[j]+nums[k] == target {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					// j,k 去重
					for j < k && nums[j] == nums[j+1] {
						j++
					}
					for j < k && nums[k] == nums[k-1] {
						k--
					}
					j++
					k--
				} else if nums[j]+nums[k] < target {
					j++
				} else {
					k--
				}
			}
		}
	}

	return res
}
