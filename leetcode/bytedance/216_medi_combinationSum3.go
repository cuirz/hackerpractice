package main

//216. 组合总和 III
//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
//
//说明：
//
//所有数字都是正整数。
//解集不能包含重复的组合。 
//示例 1:
//
//输入: k = 3, n = 7
//输出: [[1,2,4]]
//示例 2:
//
//输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]

func combinationSum3(k int, n int) [][]int {
	var dfs func(s,target,dep int) [][]int
	dfs = func(s,target,dep int) [][]int {
		var result [][]int
		for i := s; i < 10; i++ {
			if i > target || dep > k{
				break
			}else if dep == k && i == target{
				result = append(result,[]int{i})
				break
			}
			res := dfs(i+1,target-i,dep+1)
			for _, v := range res {
				result = append(result, append([]int{i}, v...))
			}
		}
		return result
	}
	return dfs(1,n,1)

}
