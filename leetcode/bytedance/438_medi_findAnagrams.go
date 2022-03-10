package main

//438. 找到字符串中所有字母异位词
//给定一个字符串s和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。
//
//字符串只包含小写英文字母，并且字符串s和 p的长度都不超过 20100。
//
//说明：
//
//字母异位词指字母相同，但排列不同的字符串。
//不考虑答案输出的顺序。
//示例1:
//
//输入:
//s: "cbaebabacd" p: "abc"
//
//输出:
//[0, 6]
//
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
//示例 2:
//
//输入:
//s: "abab" p: "ab"
//
//输出:
//[0, 1, 2]
//
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。

//思路 滑动窗口 字母字典表 统计
func findAnagrams(s string, p string) []int {
	dic := make([]int, 26)
	for _, v := range p {
		dic[v-'a']++
	}

	result := make([]int, 0)
	size := len(p)
	n := len(s)
	dicS := make([]int, 26)
	dp := make([][26]int, n+1)
	for i := 0; i < len(s); i++ {
		dicS[s[i]-'a']++
		isValid := 0
		for j := 0; j < 26; j++ {
			dp[i+1][j] = dicS[j]
			// 对比长度，s 中 p长度开始 对比
			if i >= size-1 {
				// dp 前缀和 减掉前i-size+1和 就是 dp size 部分和
				if dp[i+1][j]-dp[i-size+1][j] == dic[j] {
					isValid++
				}
			}
		}
		if isValid == 26 {
			result = append(result, i-size+1)
		}
	}
	return result

}

func main() {
	println(findAnagrams("cbaebabacd",
		"abc"))
}
