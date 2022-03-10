package main

//77. 组合
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//
//输入:n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	queue := make([]int, k)
	var dfs func(s, dep int)
	dfs = func(s, dep int) {
		if n-s < k-1-dep {
			return
		}
		queue[dep] = s
		if dep == k-1 {
			tmp := make([]int, k)
			copy(tmp, queue)
			result = append(result, tmp)
			return
		}
		for i := s; i <= n; i++ {
			dfs(i+1, dep+1)
		}
	}
	for i := 1; i <= n; i++ {
		dfs(i, 0)
	}
	return result
}
