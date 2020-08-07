package main

import (
	"bytes"
)

//49. 字母异位词分组
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。

//思路  字典表统计，给特征码
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	n := len(strs)
	tb := make(map[string]int)
	for i := 0; i < n; i++ {
		dic := [26]int{}
		for _, v := range strs[i] {
			dic[v-'a'] ++
		}
		finger := comToFinger(dic, ",")
		if tb[finger] == 0 {
			tb[finger] = len(result) + 1
			result = append(result, []string{strs[i]})
		} else {
			index := tb[finger] - 1
			result[index] = append(result[index],strs[i])
		}
	}
	return result

}

func comToFinger(a [26]int, sep string) string {
	var buff bytes.Buffer
	n := len(a)
	for i := 0; i < n; i++ {
		buff.WriteByte(byte(a[i]))
		buff.WriteByte(byte(a[i] >> 8))
		buff.WriteByte(byte(a[i] >> 16))
		buff.WriteByte(byte(a[i] >> 24))
		if i < n-1 {
			buff.WriteString(sep)
		}
	}
	return buff.String()
}
