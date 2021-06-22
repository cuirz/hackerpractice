package swordoffer

import "sort"

//剑指 Offer 38. 字符串的排列
//输入一个字符串，打印出该字符串中字符的所有排列。
//
// 
//
//你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
// 
//
//示例:
//
//输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
// 
//
//限制：
//
//1 <= s 的长度 <= 8

func permutation(s string) []string {
	array := []byte(s)
	n := len(s)
	dic := make([]byte, 0)
	result := make([]string, 0)
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	visited := make([]bool, n)
	var dfs func(x int)
	dfs = func(x int) {
		if x == n {
			result = append(result, string(dic))
			return
		}
		for i, _ := range visited {
			if visited[i] || (i > 0 && !visited[i-1] && array[i-1] == array[i]){
				continue
			}
				visited[i] = true
			dic = append(dic, array[i])
			dfs(x + 1)
			dic = dic[:len(dic)-1]
			visited[i] = false
		}
	}
	dfs(0)
	return result
}
