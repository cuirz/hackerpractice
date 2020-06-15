package main

import (
	"bytes"
)

//336. 回文对
//给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。
//
//示例 1:
//
//输入: ["abcd","dcba","lls","s","sssll"]
//输出: [[0,1],[1,0],[3,2],[2,4]]
//解释: 可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
//
//示例 2:
//
//输入: ["bat","tab","cat"]
//输出: [[0,1],[1,0]]
//解释: 可拼接成的回文串为 ["battab","tabbat"]o

// 思路
// 情况一： 字符数一样情况，字符串反转，然后从哈希表中寻找， [abcd]  [dcba]
// 情况二： 前缀时 [abc] , [ba] | [abcc] [ba] [cba]
// 情况三： 后缀   [xxabc]  [cbax] [cba]

func palindromePairs(words []string) [][]int {
	dic := make(map[string]int)
	for i, v := range words {
		dic[v] = i
	}

	result := make([][]int, 0)
	for i, v := range words {
		temp := make([]string, 0)
		// 情况1 反转
		temp = append(temp, revers(v))
		//情况2 前缀
		temp = append(temp, prefix(v)...)
		for _, word := range temp {
			if index, ok := dic[word]; ok && i != index {
				result = append(result, []int{i, index})
			}
		}
		//情况3 后缀
		// 制作v的可能的回文串列表,然后跟dic哈希表对比
		for _, word := range suffix(v) {
			if index, ok := dic[word]; ok && i != index {
				result = append(result, []int{index, i})
			}
		}
	}
	return result
}

func revers(s string) string {
	buff := bytes.Buffer{}
	for i := len(s) - 1; i > -1; i-- {
		buff.WriteByte(s[i])
	}
	return buff.String()
}

func prefix(s string) []string {
	// 找回文串
	result := make([]string, 0)
	n := len(s)
	for i := n - 1; i > -1; i-- {
		if check(s[i:]) {
			result = append(result, revers(s[:i]))
		}
	}
	return result
}

func suffix(s string) []string {
	// 找回文串
	result := make([]string, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		if check(s[:i+1]) {
			result = append(result, revers(s[i+1:]))
		}
	}
	return result
}

func check(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
