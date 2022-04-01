package main

import "sort"

//954. 二倍数对数组
//给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <=i < len(arr) / 2，都有 arr[2 * i + 1] = 2 * arr[2 * i]”时，返回 true；否则，返回 false。
//
//
//
//示例 1：
//
//输入：arr = [3,1,3,6]
//输出：false
//示例 2：
//
//输入：arr = [2,1,2,6]
//输出：false
//示例 3：
//
//输入：arr = [4,-2,2,-4]
//输出：true
//解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
//
//
//提示：
//
//0 <= arr.length <= 3 * 10^4
//arr.length 是偶数
//-10^5 <= arr[i] <= 10^5

func canReorderDoubled(arr []int) bool {
	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) < abs(arr[j])
	})

	for _, v := range arr {
		if dic[v] > dic[v*2] {
			return false
		}
		dic[v*2] -= dic[v]
		dic[v] = 0
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	println(canReorderDoubled([]int{2, 1, 2, 1, 1, 1, 2, 2}))
}
