package hashtable

// Substring with Concatenation of All Words : 030
// 串联所有单词的子串
// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
//
// reference: https://leetcode.com/problems/substring-with-concatenation-of-all-words/
// 思路:
// 首先将words转换成字典wordCount, key为单词, value为单词出现的次数。以words所有单词总长度为窗口, 截取字符串s,
// 然后再将子串以单词长度进行分割。利用第二个 字典tmp 来统计子串单词出现的次数, 最后比较 tmp 是否与 wordCount 相等.

// FindSubstring 串联所有单词的子串
func FindSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(words) == 0 {
		return res
	}
	wordLen := len(words[0]) // 每个单词的长度
	allLen := wordLen * len(words)
	wordCount := make(map[string]int)
	for _, v := range words {
		if _, ok := wordCount[v]; ok {
			wordCount[v]++
		} else {
			wordCount[v] = 1
		}
	}

	var w, word string

	for i := 0; i <= len(s)-allLen; i++ {
		word = s[i : i+allLen]
		tmp := make(map[string]int)

		for j := 0; j < len(word); j += wordLen {
			w = word[j : j+wordLen]
			if _, ok := wordCount[w]; ok {
				tmp[w]++
			} else {
				break
			}
		}
		// 判断 tmp 和 wordCount 是否相等
		flag := false
		for k, v := range wordCount {
			if value, ok := tmp[k]; ok && v == value {
				flag = true
			} else {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}

	return res
}
