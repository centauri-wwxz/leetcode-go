package hasttable

// Longest Substring Without Repeating Characters: 003
// 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// reference: https://leetcode.com/problems/longest-substring-without-repeating-characters/
// 思路:
// 两个指针, 一个记录未重复地方的位置, 一个用来遍历. 使用 map 来判断是否重复.

// LengthOfLongestSubstring 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	var max, i int
	tmp := make(map[byte]int)
	for j := 0; j < len(s); j++ {
		if value, ok := tmp[s[j]]; ok {
			i = larger(i, value+1)
		}
		tmp[s[j]] = j
		max = larger(max, j-i+1)
	}
	return max
}

func larger(i, j int) int {
	if i > j {
		return i
	}
	return j
}
