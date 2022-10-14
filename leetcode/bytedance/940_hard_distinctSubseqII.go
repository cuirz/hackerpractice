package main

//940. 不同的子序列 II
//给定一个字符串 s，计算 s 的 不同非空子序列 的个数。因为结果可能很大，所以返回答案需要对 10^9 + 7 取余 。
//
//字符串的 子序列 是经由原字符串删除一些（也可能不删除）字符但不改变剩余字符相对位置的一个新字符串。
//
//例如，"ace" 是 "abcde" 的一个子序列，但 "aec" 不是。
//
//
//示例 1：
//
//输入：s = "abc"
//输出：7
//解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
//示例 2：
//
//输入：s = "aba"
//输出：6
//解释：6 个不同的子序列分别是 "a", "b", "ab", "ba", "aa" 以及 "aba"。
//示例 3：
//
//输入：s = "aaa"
//输出：3
//解释：3 个不同的子序列分别是 "a", "aa" 以及 "aaa"。
//jjj
//
//提示：
//
//1 <= s.length <= 2000
//s 仅由小写英文字母组成

func distinctSubseqII(s string) int {
	n := len(s)
	mod := int(1e9 + 7)
	result, cnt := 1, 0
	preCnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt = result
		index := s[i] - 'a'
		result = ((result+cnt)%mod - preCnt[index]%mod + mod) % mod
		preCnt[index] = cnt
	}
	return result - 1
}
