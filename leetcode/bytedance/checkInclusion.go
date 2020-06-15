package main

import (
	"strings"
)

// 统计 字母数量多少就能对比
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	c1 := make([]int, 26)
	c2 := make([]int, 26)

	for _, w := range s1 {
		c1[w-'a']++
	}

	for i, w := range s2 {
		if i >= len(s1) {
			c2[s2[i-len(s1)]-'a']--
		}
		c2[w-'a']++
		if Compare(c1, c2) {
			return true
		}
	}
	return false
}
func Compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func checkInclusion2(s1 string, s2 string) bool {
	result := new([][]rune)
	dfs(result, make([]rune, 0), []rune(s1), make([]bool, len(s1)))
	//println(fmt.Sprintf("%v", result))
	for _, w := range *result {
		if strings.Contains(s2, string(w)) {
			return true
		}
	}
	return false
}
func dfs(result *[][]rune, path []rune, res []rune, pick []bool) {
	if len(path) == len(res) {
		*result = append(*result, append(make([]rune, 0), path...))
		return
	}
	for i, w := range res {
		if !pick[i] {
			path = append(path, w)
			pick[i] = true
			dfs(result, path, res, pick)
			pick[i] = false
			path = path[:len(path)-1]
		}
	}
}

// path 是 添加到 result时必须拷贝一份，要不然传进去的是相同内存会被修改
func dfs2(result *[][]string, path []string, res []string, pick []bool) {
	if len(path) == len(res) {
		*result = append(*result, append(make([]string, 0), path...))
		return
	}
	for i, w := range res {
		if !pick[i] {
			path = append(path, w)
			pick[i] = true
			dfs2(result, path, res, pick)
			pick[i] = false
			path = path[:len(path)-1]
		}
	}
}

func main() {
	//result := new([][]string)
	//dfs(result, make([]string, 0), []string{"1", "2", "3", "4"}, make([]bool, 5))
	//println(fmt.Sprintf("%v", result))

	print(checkInclusion("acd", "dcda"))
}
