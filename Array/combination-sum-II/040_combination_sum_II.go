package array

import "sort"

// Combination Sum II: 40
// 组合总和 II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明：
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
//
// reference:https://leetcode.com/problems/combination-sum-ii/
// 思路: 回溯法, 类似 Combination Sum, 这里的注意重复值

// CombinationSum2 组合总和 II
func CombinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtrack(&res, []int{}, candidates, target, 0)
	return res
}

func backtrack(res *[][]int, temp, candidates []int, remain, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		t := make([]int, len(temp))
		copy(t, temp)
		*res = append(*res, t)
	} else {
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] { // 防止重复
				continue
			}

			temp = append(temp, candidates[i])
			backtrack(res, temp, candidates, remain-candidates[i], i+1)
			temp = temp[:len(temp)-1]
		}
	}
}
