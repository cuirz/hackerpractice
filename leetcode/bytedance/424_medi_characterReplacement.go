package main

//424. 替换后的最长重复字符
//给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换k次。在执行上述操作后，找到包含重复字母的最长子串的长度。
//
//注意：字符串长度 和 k 不会超过104。
//
//
//
//示例 1：
//
//输入：s = "ABAB", k = 2
//输出：4
//解释：用两个'A'替换为两个'B',反之亦然。
//示例 2：
//
//输入：s = "AABABBA", k = 1
//输出：4
//解释：
//将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
//子串 "BBBB" 有最长重复字母, 答案为 4。

//思路 双指针滑动窗口
func characterReplacement(s string, k int) int {
	var count [26]int
	sum, left := 0, 0
	for i, v := range s {
		count[v-'A']++
		sum = max(sum, count[v-'A'])
		if i+1-left-sum > k {
			count[s[left]-'A']--
			left++
		}
	}
	return len(s) - left

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
