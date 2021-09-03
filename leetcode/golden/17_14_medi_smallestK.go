package main

import (
	"container/heap"
	"sort"
)

//17.14. 最小K个数
//设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
//示例：
//
//输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
//提示：
//
//0 <= len(arr) <= 100000
//0 <= k <= min(100000, len(arr))

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	stack := &hp{arr[:k]}
	heap.Init(stack)

	for i := k; i < len(arr); i++ {
		if stack.IntSlice[0] > arr[i] {
			stack.IntSlice[0] = arr[i]
			heap.Fix(stack, 0)
		}
	}
	return stack.IntSlice
}

type hp struct {
	sort.IntSlice
}

func (this *hp) Less(i, j int) bool {
	return this.IntSlice[i] > this.IntSlice[j]
}

func (hp) Push(v interface{}) {
}
func (hp) Pop() (_ interface{}) {
	return
}
