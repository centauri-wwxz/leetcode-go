package array

import "sort"

// Combination Sum: 39
// 组合总和
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明：
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
//
// reference:https://leetcode.com/problems/combination-sum/
// 思路: 回溯法

// CombinationSum 组合总和
func CombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtrackForCombination(&res, []int{}, candidates, target, 0)
	return res
}

// backtrackForCombination
// res 结果集
// temp 保存结果的临时切片
// candidates 原切片
// remain target-临时结果的剩余值, 小于 0 时表示不存在, 等于 0 时表示结果刚好是 temp
// start 内嵌循环的起始值
func backtrackForCombination(res *[][]int, temp, candidates []int, remain, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		t := make([]int, len(temp))
		copy(t, temp) // 切片的底层是数组, 这边使用临时变量避免修改切片的值
		*res = append(*res, t)
	} else {
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i]) // 保存元素到临时结果
			backtrackForCombination(res, temp, candidates, remain-candidates[i], i)
			temp = temp[:len(temp)-1] // 如果失败, 则回溯到上一个临时结果切片
		}
	}

}
