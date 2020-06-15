package main

import (
	"sort"
)

//767. 重构字符串
//给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
//
//若可行，输出任意可行的结果。若不可行，返回空字符串。
//
//示例 1:
//
//输入: S = "aab"
//输出: "aba"
//示例 2:
//
//输入: S = "aaab"
//输出: ""

//思路： 统计字母，排序，然后交替填坑,先打印偶数坑，再填奇数坑
//  当 最多字符数大于 (len+1)/2 时 无解

func reorganizeString(S string) string {
	n := len(S)
	dic := make(map[byte]int)
	array := make([]byte, 0)
	result := make([]byte, n)

	for i, _ := range S {
		c := S[i]
		if dic[c] == 0 {
			array = append(array, c)
		}
		dic[c]++
		if dic[c] > (n+1)/2 {
			return ""
		}
	}

	sort.Slice(array, func(i, j int) bool {
		return dic[array[i]] > dic[array[j]]
	})

	index := 0
	for i := 0; i < n; i += 2 {
		c := array[index]
		result[i] = c
		dic[c]--
		if dic[c] < 1 {
			index++
		}
	}
	for i := 1; i < n; i += 2 {
		c := array[index]
		result[i] = c
		dic[c]--
		if dic[c] < 1 {
			index++
		}

	}

	return string(result)

}
func main() {
	println(reorganizeString("bbbbbxbwhbbbbmsybtbbbbbkncyybnjbhxbbrxibcjybb"))
	//"bybybybybxbxbxbcbcbhbhbjbjbnbnbibkbmbrbsbtbwb"
}
