package unsorted

// Check If It Is a Good Array:1250
// 给你一个正整数数组 nums，你需要从中任选一些子集，然后将子集中每一个数乘以一个任意整数，并求出他们的和。
// 假如该和结果为 1，那么原数组就是一个「好数组」，则返回 True；否则请返回 False。
//
// reference: https://leetcode.com/problems/check-if-it-is-a-good-array/
// 思路:
// 数学知识: 如果多个数乘以任意整数的值为 1, 则它们的最大公约数为 1.

func isGoodArray(nums []int) bool {
	d := nums[0]
	for i := 1; i < len(nums); i++ {
		d = gcd(d, nums[i])
	}
	return d == 1
}

// gcd(a, b) = gcd(b, a mod b)
func gcd(x, y int) int {
	if y > x {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
