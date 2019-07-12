package array

import "math"

// Container With Most Water: 11
// 盛最多水的容器
// 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，
// 垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。
//
// reference: https://leetcode.com/problems/container-with-most-water
// 思路:
// 容器的面积是底*高, 底为数组下标之差, 高为两元素值的最小值.
// 从首尾开始遍历, 记录面积, 哪边的高越小则舍弃, 并且让下标值向中间靠拢.

// MaxArea 盛最多水的容器
func MaxArea(height []int) int {
	l, r, maxArea := 0, len(height)-1, float64(0)
	for l < r {
		maxArea = math.Max(maxArea, float64(r-l)*math.Min(float64(height[l]), float64(height[r])))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return int(maxArea)
}
