package unsorted

import (
	"bytes"
)

// Minimum Remove to Make Valid Parentheses:1249
// 给你一个由 '('、')' 和小写字母组成的字符串 s。
// 你需要从字符串中删除最少数目的 '(' 或者 ')' （可以删除任意位置的括号)，使得剩下的「括号字符串」有效。
// 请返回任意一个合法字符串。
// 有效「括号字符串」应当符合以下 任意一条 要求：
// 空字符串或只包含小写字母的字符串
// 可以被写作 AB（A 连接 B）的字符串，其中 A 和 B 都是有效「括号字符串」
// 可以被写作 (A) 的字符串，其中 A 是一个有效的「括号字符串」
//
// reference: https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// 思路:
// 成对的匹配可以使用栈思想, push '(', 遇到 ')' pop.

func minRemoveToMakeValid(s string) string {
	// stack := newStack()
	pre := make([]int, 0)
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '(' {
			pre = append(pre, i)
		} else if b[i] == ')' {
			if len(pre) == 0 {
				b[i] = ' '
			} else {
				pre = pre[1:]
			}
		}
	}
	for i := 0; i < len(pre); i++ {
		b[pre[i]] = ' '
	}
	// todo: replace " " with "" cost too much
	return string(bytes.Replace(b, []byte(" "), []byte(""), -1))
}

func newStack() *stack {
	return &stack{val: make([]int, 0)}
}

type stack struct {
	val []int
}

func (s *stack) push(i int) {
	s.val = append(s.val, i)
}

func (s *stack) pop() int {
	v := s.val[0]
	s.val = s.val[1:]
	return v
}
func (s *stack) isEmpty() bool {
	return len(s.val) == 0
}
