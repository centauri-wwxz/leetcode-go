package array

// Search Insert Position: 35
// 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 你可以假设数组中无重复元素。
//
// reference:https://leetcode.com/problems/search-insert-position/
// 思路: 二分查找

// SearchInsert 搜索插入位置
func SearchInsert(nums []int, target int) int {
	l := len(nums)
	if target < nums[0] {
		return 0
	}
	if target > nums[l-1] {
		return l
	}

	return search(nums, target)
}

func search(nums []int, target int) int {
	low, mid, high := 0, 0, len(nums)
	for low < high {
		mid = (low + high) / 2
		if nums[mid] < target {
			if nums[mid+1] >= target {
				return mid + 1
			}
			low = mid + 1

		} else if nums[mid] > target {
			if nums[mid-1] < target {
				return mid
			} else if nums[mid-1] == target {
				return mid - 1
			}
			high = mid
		} else {
			return mid
		}
	}
	return 0
}
