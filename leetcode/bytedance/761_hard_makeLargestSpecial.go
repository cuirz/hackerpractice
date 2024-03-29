package main

import (
	"sort"
	"strings"
)

//
//761. 特殊的二进制序列
//特殊的二进制序列是具有以下两个性质的二进制序列：
//
//0 的数量与 1 的数量相等。
//二进制序列的每一个前缀码中 1 的数量要大于等于 0 的数量。
//给定一个特殊的二进制序列S，以字符串形式表示。定义一个操作 为首先选择S的两个连续且非空的特殊的子串，然后将它们交换。（两个子串为连续的当且仅当第一个子串的最后一个字符恰好为第二个子串的第一个字符的前一个字符。)
//
//在任意次数的操作之后，交换后的字符串按照字典序排列的最大的结果是什么？
//
//示例 1:
//
//输入: S = "11011000"
//输出: "11100100"
//解释:
//将子串 "10" （在S[1]出现） 和 "1100" （在S[3]出现）进行交换。
//这是在进行若干次操作后按字典序排列最大的结果。
//说明:
//
//S的长度不超过50。
//S保证为一个满足上述定义的特殊 的二进制序列。

//思路 分治
// 1 看成 （
// 0 堪称  ）
func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	left, cnt := 0, 0
	subSet := sort.StringSlice{}
	for i, v := range s {
		if v == '1' {
			cnt++
		} else if cnt--; cnt == 0 {
			subSet = append(subSet, "1"+makeLargestSpecial(s[left+1:i])+"0")
			left = i + 1
		}
	}
	sort.Sort(sort.Reverse(subSet))
	return strings.Join(subSet, "")
}
