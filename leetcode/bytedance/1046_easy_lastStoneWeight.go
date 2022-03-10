package main

import (
	"container/heap"
	"sort"
)

//1046. 最后一块石头的重量
//有一堆石头，每块石头的重量都是正整数。
//
//每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为x 和y，且x <= y。那么粉碎的可能结果如下：
//
//如果x == y，那么两块石头都会被完全粉碎；
//如果x != y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
//最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
//
//
//
//示例：
//
//输入：[2,7,4,1,8,1]
//输出：1
//解释：
//先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
//再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
//接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
//最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
//
//
//提示：
//
//1 <= stones.length <= 30
//1 <= stones[i] <= 1000

func lastStoneWeight(stones []int) int {
	q := &pq{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0

}

type pq struct {
	sort.IntSlice
}

func (h *pq) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *pq) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *pq) Pop() interface{} {
	x := h.IntSlice
	v := x[len(x)-1]
	h.IntSlice = x[:len(x)-1]
	return v
}
func (h *pq) push(x int) { heap.Push(h, x) }
func (h *pq) pop() int   { return heap.Pop(h).(int) }
