package main

import "container/heap"

//786. 第 K 个最小的素数分数
//给你一个按递增顺序排序的数组 arr 和一个整数 k 。数组 arr 由 1 和若干 素数 组成，且其中所有整数互不相同。
//
//对于每对满足 0 < i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。
//
//那么第k个最小的分数是多少呢? 以长度为 2 的整数数组返回你的答案, 这里answer[0] == arr[i]且answer[1] == arr[j] 。
//
//
//示例 1：
//
//输入：arr = [1,2,3,5], k = 3
//输出：[2,5]
//解释：已构造好的分数,排序后如下所示:
//1/5, 1/3, 2/5, 1/2, 3/5, 2/3
//很明显第三个最小的分数是 2/5
//示例 2：
//
//输入：arr = [1,7], k = 1
//输出：[1,7]
//
//
//提示：
//
//2 <= arr.length <= 1000
//1 <= arr[i] <= 3 * 10^4
//arr[0] == 1
//arr[i] 是一个 素数 ，i > 0
//arr 中的所有数字 互不相同 ，且按 严格递增 排序
//1 <= k <= arr.length * (arr.length - 1) / 2

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	h := make(hp, n-1)
	for i := 1; i < len(arr); i++ {
		h[i-1] = frac{arr[0], arr[i], 0, i}
	}
	heap.Init(&h)
	for i := 1; i < k; i++ {
		top := heap.Pop(&h).(frac)
		if top.l+1 < top.r {
			heap.Push(&h, frac{arr[top.l+1], arr[top.r], top.l + 1, top.r})
		}
	}
	return []int{h[0].x, h[0].y}
}

type frac struct {
	x, y, l, r int
}
type hp []frac

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(l, r int) bool {
	return h[l].x*h[r].y < h[l].y*h[r].x
}

func (h hp) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(frac))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
