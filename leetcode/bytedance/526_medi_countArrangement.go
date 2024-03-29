package main

//526. 优美的排列
//假设有从 1 到 N 的N个整数，如果从这N个数字中成功构造出一个数组，使得数组的第 i位 (1 <= i <= N) 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：
//
//第i位的数字能被i整除
//i 能被第 i 位上的数字整除
//现在给定一个整数 N，请问可以构造多少个优美的排列？
//
//示例1:
//
//输入: 2
//输出: 2
//解释:
//
//第 1 个优美的排列是 [1, 2]:
//第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
//第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除
//
//第 2 个优美的排列是 [2, 1]:
//第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
//第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
//说明:
//
//N 是一个正整数，并且不会超过15。

//思路 回溯
func countArrangement(n int) int {
	var result int
	visited := make([]bool, n+1)
	nums := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				nums[i] = append(nums[i], j)
			}
		}
	}
	var backTrace func(index int)
	backTrace = func(index int) {
		if index > n {
			result++
			return
		}
		for _, x := range nums[index] {
			if !visited[x] {
				visited[x] = true
				backTrace(index + 1)
				visited[x] = false
			}
		}
	}
	backTrace(1)
	return result
}
