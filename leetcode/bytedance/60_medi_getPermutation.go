package main

import (
	"strconv"
	"strings"
)

//60. 第k个排列
//给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
//
//按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//"123"
//"132"
//"213"
//"231"
//"312"
//"321"
//给定 n 和 k，返回第 k 个排列。
//
//说明：
//
//给定 n 的范围是 [1, 9]。
//给定 k 的范围是[1,  n!]。
//示例 1:
//
//输入: n = 3, k = 3
//输出: "213"
//示例 2:
//
//输入: n = 4, k = 9
//输出: "2314"

//思路
// k % (n-1)!
func getPermutation(n int, k int) string {
	result := make([]string, 0)
	presum := make([]int, n+1)
	queue := make([]int, n)
	for i := 1; i <= n; i++ {
		queue[i-1] = i
	}
	stack := func(i int) string {
		if i < 0 {
			i = len(queue) - 1
		}
		res := strconv.Itoa(queue[i])
		tmp := queue[:i]
		queue = append(tmp, queue[i+1:]...)
		return res
	}
	presum[0], presum[1] = 1, 1
	for i := 2; i <= n; i++ {
		presum[i] = presum[i-1] * i
	}
	index := k
	for i := n; i > 0; i-- {
		if index > 0{
			result = append(result, stack((index-1)/presum[i-1]))
		}else{
			result = append(result, stack(-1))
		}
		index = index % presum[i-1]
	}
	return strings.Join(result, "")
}

func main() {
	println(getPermutation(4, 6))
}
