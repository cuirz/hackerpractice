package main

//1345. 跳跃游戏 IV
//给你一个整数数组arr，你一开始在数组的第一个元素处（下标为 0）。
//
//每一步，你可以从下标i跳到下标：
//
//i + 1满足：i + 1 < arr.length
//i - 1满足：i - 1 >= 0
//j满足：arr[i] == arr[j]且i != j
//请你返回到达数组最后一个元素的下标处所需的最少操作次数。
//
//注意：任何时候你都不能跳到数组外面。
//
//
//
//示例 1：
//
//输入：arr = [100,-23,-23,404,100,23,23,23,3,404]
//输出：3
//解释：那你需要跳跃 3 次，下标依次为 0 --> 4 --> 3 --> 9 。下标 9 为数组的最后一个元素的下标。
//示例 2：
//
//输入：arr = [7]
//输出：0
//解释：一开始就在最后一个元素处，所以你不需要跳跃。
//示例 3：
//
//输入：arr = [7,6,9,6,9,6,9,7]
//输出：1
//解释：你可以直接从下标 0 处跳到下标 7 处，也就是数组的最后一个元素处。
//示例 4：
//
//输入：arr = [6,1,9]
//输出：2
//示例 5：
//
//输入：arr = [11,22,7,7,7,7,7,7,7,22,13]
//输出：3
//
//
//提示：
//
//1 <= arr.length <= 5 * 10^4
//-10^8 <= arr[i] <= 10^8

func minJumps(arr []int) int {
	dic := make(map[int][]int)
	for i, v := range arr {
		dic[v] = append(dic[v], i)
	}
	type pair struct {
		index, depth int
	}
	queue := make([]pair, 0)
	n := len(arr)
	queue = append(queue, pair{0, 0})
	visited := make([]bool, n)
	visited[0] = true
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.index == n-1 {
			return top.depth
		}
		if list, ok := dic[arr[top.index]]; ok {
			for _, v := range list {
				if !visited[v] {
					queue = append(queue, pair{v, top.depth + 1})
					visited[v] = true
				}
			}
		}
		delete(dic, arr[top.index])
		if top.index > 0 && !visited[top.index-1] {
			queue = append(queue, pair{top.index - 1, top.depth + 1})
			visited[top.index-1] = true
		}
		if !visited[top.index+1] {
			queue = append(queue, pair{top.index + 1, top.depth + 1})
			visited[top.index+1] = true
		}

	}
	return 0
}
