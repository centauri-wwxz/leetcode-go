package array

import (
	"math"
)

// Median of Two Sorted Arrays : 004
// 寻找两个有序数组的中位数
// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
// 你可以假设 nums1 和 nums2 不会同时为空。
//
// reference: https://leetcode.com/problems/median-of-two-sorted-arrays/
// 思路:
// 每个数组都有一个中位数区分点 partitionX 和 partitionY, 区分点两边的数要满足:
// leftX <= rightY && leftY <= rightX
// 如果不满足, 则需要移动标点
// 2, 3, 4, 5       ===> leftX:3 rightX:4
// 1, 2, 3, 4, 6    ===> leftY:2 rightY:3

// FindMedianSortedArrays 寻找两个有序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	x, y := len(nums1), len(nums2)
	leftX, leftY, rightX, rightY := 0, 0, 0, 0
	low := 0
	high := x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y)/2 - partitionX

		if partitionX == 0 {
			leftX = math.MinInt64
		} else {
			leftX = nums1[partitionX-1]
		}
		if partitionX == x {
			rightX = math.MaxInt64
		} else {
			rightX = nums1[partitionX]
		}

		if partitionY == 0 {
			leftY = math.MinInt64
		} else {
			leftY = nums2[partitionY-1]
		}
		if partitionY == y {
			rightY = math.MaxInt64
		} else {
			rightY = nums2[partitionY]
		}

		if leftX <= rightY && leftY <= rightX {
			if (x+y)%2 == 0 {
				return (math.Max(float64(leftX), float64(leftY)) + math.Min(float64(rightX), float64(rightY))) / 2
			}
			return math.Min(float64(rightX), float64(rightY))
		} else if leftX > rightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return math.SmallestNonzeroFloat64
}
