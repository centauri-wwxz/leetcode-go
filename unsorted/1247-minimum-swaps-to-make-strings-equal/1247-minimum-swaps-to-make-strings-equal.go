package unsorted

// Minimum Swaps to Make Strings Equal:1247
// 有两个长度相同的字符串 s1 和 s2，且它们其中 只含有 字符 "x" 和 "y"，你需要通过「交换字符」的方式使这两个字符串相同。
// 每次「交换字符」的时候，你都可以在两个字符串中各选一个字符进行交换。
// 交换只能发生在两个不同的字符串之间，绝对不能发生在同一个字符串内部。也就是说，我们可以交换 s1[i] 和 s2[j]，但不能交换 s1[i] 和 s1[j]。
// 最后，请你返回使 s1 和 s2 相同的最小交换次数，如果没有方法能够使得这两个字符串相同，则返回 -1 。
//
// reference: https://leetcode.com/problems/minimum-swaps-to-make-strings-equal/
// 思路:
// 1. 如果字符相同, 则跳过;
// 2. 如果字符串s1 中的 x 和 y 个数(Nx,Ny) 相加不为偶数, 则无法交换;
// 3. Nx 如果是偶数, 则交换的次数为 (Nx+Ny)/2; 否则, (Nx+Ny)/2+1;.

func minimumSwap(s1 string, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	var nx, ny int
	for i, c := range s1 {
		if byte(c) == s2[i] {
			continue
		}
		if c == 'x' {
			nx++
		} else {
			ny++
		}
	}
	if (nx+ny)&1 == 0 {
		if nx&1 == 1 {
			return (nx+ny)/2 + 1
		}
		return (nx + ny) / 2
	}

	return -1
}
