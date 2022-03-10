package main

//389. 找不同
//给定两个字符串 s 和 t，它们只包含小写字母。
//
//字符串t由字符串s随机重排，然后在随机位置添加一个字母。
//
//请找出在 t 中被添加的字母。
//
//
//
//示例 1：
//
//输入：s = "abcd", t = "abcde"
//输出："e"
//解释：'e' 是那个被添加的字母。
//示例 2：
//
//输入：s = "", t = "y"
//输出："y"
//示例 3：
//
//输入：s = "a", t = "aa"
//输出："a"
//示例 4：
//
//输入：s = "ae", t = "aea"
//输出："a"
//
//
//提示：
//
//0 <= s.length <= 1000
//t.length == s.length + 1
//s 和 t 只包含小写字母

func findTheDifference(s string, t string) byte {
	var dic [26]int
	for _, v := range s {
		dic[v-'a']++
	}
	for _, v := range t {
		dic[v-'a']--
	}
	for i, v := range dic {
		if v < 0 {
			return byte(i + 'a')
		}
	}
	return 0
}
