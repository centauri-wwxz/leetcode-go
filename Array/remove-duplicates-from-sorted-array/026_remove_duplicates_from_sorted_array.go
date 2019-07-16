package array

// Remove Duplicates from Sorted Array: 26
// 删除排序数组中的重复项
// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
// reference:https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// 思路:
// 交换元素的值, 得到新的不重复数组的长度, 也就是原数组不相同元素的个数, 使用一个额外的计数器即可完成.
// 原地操作: 切片不可以拼装或者重组, 只能交换索引的值.

// RemoveDuplicates 删除排序数组中的重复项
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0 // index and counter
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
