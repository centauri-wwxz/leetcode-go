package array

// Remove Element: 27
// 移除元素
// 给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
// reference: https://leetcode-cn.com/problems/remove-element/
// 思路:
// 类似 026, 使用双指针, 遇到不同于 val 则交换 i 和 j 的值, 同时 i 也是计数器.
// 注意: 原地操作

// RemoveElement 移除元素
func RemoveElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
