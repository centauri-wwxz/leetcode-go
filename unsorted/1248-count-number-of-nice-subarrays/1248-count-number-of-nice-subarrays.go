package unsorted

// Count Number of Nice Subarrays: 1248
// 给你一个整数数组 nums 和一个整数 k。
// 如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
// 请返回这个数组中「优美子数组」的数目。
//
// reference: https://leetcode.com/problems/count-number-of-nice-subarrays/
// 思路:
// 动态规划, 每遇到奇数, 记录它前面的偶数个数Ni,若遇到第 k 个奇数时,
// 此时的数目为前一个奇数的偶数个数 Nk-1; (后续遇到的偶数数目都是 Nk-1).

func numberOfSubarrays(nums []int, k int) int {
	dp := make([]int, 0)
	var cnt, ret int
	for i := 0; i < len(nums); i++ {
		cnt++
		if nums[i]&1 == 1 {
			dp = append(dp, cnt)
			cnt = 0
		}
		if len(dp) >= k {
			ret += dp[len(dp)-k]
		}
	}
	return ret
}
