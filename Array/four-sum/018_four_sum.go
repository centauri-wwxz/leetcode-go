package array

import "sort"

// 4Sum: 18
//  四数之和
// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。
// reference:
// 思路:
// 与 3Sum 类似, 都是在指定的区间内移动下标. 分解成 3Sum -> 2Sum, 注意去重和边界值判断

// FourSum 四数之和
func FourSum(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)
	if n < 4 {
		return res
	}
	// sort slice
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		// 如果每次遍历的前四个大于 target, 跳出循环
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 如果当前i 与后三者之和小于 target, i++
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		// i 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 3Sum
		for j := i + 1; j < n-2; j++ {
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// 如果 i,当前j 与后两者之和小于 target, i++
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			// j 去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 2Sum
			start := j + 1
			end := n - 1
			for start < end {
				if nums[i]+nums[j]+nums[start]+nums[end] == target {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					for start < end && nums[start] == nums[start+1] {
						start++
					}
					for start < end && nums[end] == nums[end-1] {
						end--
					}
					start++
					end--
				} else if nums[i]+nums[j]+nums[start]+nums[end] < target {
					start++
				} else {
					end--
				}
			}

		}

	}

	return res
}
