package main

import (
	"container/heap"
	"sort"
)

//295. 数据流的中位数
//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
//例如，
//
//[2,3,4]的中位数是 3
//
//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
//设计一个支持以下两种操作的数据结构：
//
//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
//double findMedian() - 返回目前所有元素的中位数。
//示例：
//
//addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3)
//findMedian() -> 2
//进阶:
//
//如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
//如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

//思路 双优先队列

type MedianFinder struct {
	pqMin, pqMax hp
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	pqMin, pqMax := &this.pqMin, &this.pqMax
	if pqMin.Len() == 0 || num <= -pqMin.IntSlice[0] {
		heap.Push(pqMin, -num)
		if pqMax.Len()+1 < pqMin.Len() {
			heap.Push(pqMax, -heap.Pop(pqMin).(int))
		}

	} else {
		heap.Push(pqMax, num)
		if pqMax.Len() > pqMin.Len() {
			heap.Push(pqMin, -heap.Pop(pqMax).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.pqMin.Len() > this.pqMax.Len() {
		return float64(-this.pqMin.IntSlice[0])
	}
	return float64(this.pqMax.IntSlice[0]-this.pqMin.IntSlice[0]) / 2

}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))

}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
