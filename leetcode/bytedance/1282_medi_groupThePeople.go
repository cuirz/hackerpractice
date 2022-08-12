package main

//1282. 用户分组
//有n个人被分成数量未知的组。每个人都被标记为一个从 0 到 n - 1 的唯一ID。
//
//给定一个整数数组 groupSizes ，其中groupSizes[i]是第 i 个人所在的组的大小。例如，如果groupSizes[1] = 3，则第 1 个人必须位于大小为 3 的组中。
//
//返回一个组列表，使每个人 i 都在一个大小为groupSizes[i]的组中。
//
//每个人应该恰好只出现在一个组中，并且每个人必须在一个组中。如果有多个答案，返回其中任何一个。可以保证给定输入至少有一个有效的解。
//
//
//
//示例 1：
//
//输入：groupSizes = [3,3,3,3,3,1,3]
//输出：[[5],[0,1,2],[3,4,6]]
//解释：
//第一组是 [5]，大小为 1，groupSizes[5] = 1。
//第二组是 [0,1,2]，大小为 3，groupSizes[0] = groupSizes[1] = groupSizes[2] = 3。
//第三组是 [3,4,6]，大小为 3，groupSizes[3] = groupSizes[4] = groupSizes[6] = 3。
//其他可能的解决方案有 [[2,1,6],[5],[0,4,3]] 和 [[5],[0,6,2],[4,3,1]]。
//示例 2：
//
//输入：groupSizes = [2,1,3,3,3,2]
//输出：[[1],[0,5],[2,3,4]]
//
//
//提示：
//
//groupSizes.length == n
//1 <= n<= 500
//1 <=groupSizes[i] <= n

//思路 哈希表
func groupThePeople(groupSizes []int) [][]int {
	idxs := map[int]int{}
	index := 0
	var ok bool
	var result [][]int
	for i, v := range groupSizes {
		ok = false
		if index, ok = idxs[v]; !ok {
			index = len(result)
			idxs[v] = index
			result = append(result, []int{})
			result[index] = append(result[index], i)
		} else {
			result[index] = append(result[index], i)
		}
		if len(result[index]) == v {
			delete(idxs, v)
		}
	}
	return result
}
