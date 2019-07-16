package array

import (
	"sort"
)

// Next Permutation: 31
// 下一个排列
// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
// 必须原地修改，只允许使用额外常数空间。
// 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1
//
// reference:
// 思路:
// 先从数组尾部遍历, 找到第一个 nums[i] > nums[i-1] 时的 数i-1, 然后再从尾部遍历, 找到第一个大于 nums[i-1]的,
// 交换两个的值, 最后将 nums[i:] 反转

// NextPermutation 下一个排列
func NextPermutation(nums []int) {
	// min := 0
	changed := false
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			changed = true
			// 得到 i-1
			// 查找 len(nums)-1 ~ i 里第一个大于 nums[i-1] 的
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					// swap i-1 and min
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}

			// reverse the nums[i:]
			for k := i; k <= (len(nums)-1+i)/2; k++ {
				nums[k], nums[len(nums)-1+i-k] = nums[len(nums)-1+i-k], nums[k]
			}
			break
		}
	}
	if !changed {
		sort.Ints(nums)
	}
}
