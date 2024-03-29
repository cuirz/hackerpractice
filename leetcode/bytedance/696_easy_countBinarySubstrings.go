package main

//696. 计数二进制子串
//给定一个字符串s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
//
//重复出现的子串要计算它们出现的次数。
//
//示例 1 :
//
//输入: "00110011"
//输出: 6
//解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
//
//请注意，一些重复出现的子串要计算它们出现的次数。
//
//另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
//示例 2 :
//
//输入: "10101"
//输出: 4
//解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
//注意：
//
//s.length在1到50,000之间。
//s只包含“0”或“1”字符。

//思路
func countBinarySubstrings(s string) int {
	result := 0
	swap := [2]int{0, 0}
	var pre byte
	for i := 0; i < len(s); i++ {
		if pre == s[i] {
			swap[0]++
		} else {
			pre = s[i]
			swap[0], swap[1] = 1, swap[0]
		}
		if swap[0] <= swap[1] {
			result++
		}

	}
	return result
}
