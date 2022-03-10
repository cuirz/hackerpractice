package main

import (
	"container/heap"
	"sort"
)

//480. 滑动窗口中位数
//中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
//
//例如：
//
//[2,3,4]，中位数是3
//[2,3]，中位数是 (2 + 3) / 2 = 2.5
//给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。
//
//
//
//示例：
//
//给出nums = [1,3,-1,-3,5,3,6,7]，以及k = 3。
//
//窗口位置                      中位数
//---------------               -----
//[1  3  -1] -3  5  3  6  7       1
//1 [3  -1  -3] 5  3  6  7      -1
//1  3 [-1  -3  5] 3  6  7      -1
//1  3  -1 [-3  5  3] 6  7       3
//1  3  -1  -3 [5  3  6] 7       5
//1  3  -1  -3  5 [3  6  7]      6
//因此，返回该滑动窗口的中位数数组[1,-1,-1,3,5,6]。
//
//
//
//提示：
//
//你可以假设k始终有效，即：k 始终小于输入的非空数组的元素个数。
//与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。

//思路 双优先队列+延迟删除
func medianSlidingWindow(nums []int, k int) []float64 {
	left, right := &PQueue{minus: true}, &PQueue{minus: false}
	delayDic := make(map[int]int)
	balance := func() {
		if left.size > right.size+1 {
			right.push(-left.pop())
			left.prune(delayDic)
		} else if left.size < right.size {
			left.push(-right.pop())
			right.prune(delayDic)
		}
	}
	inserts := func(x int) {
		if left.size == 0 || x <= -left.top() {
			left.push(-x)
		} else {
			right.push(x)
		}
		balance()
	}
	remove := func(x int) {
		delayDic[x]++
		if x <= -left.top() {
			left.size--
			if left.top() == -x {
				left.prune(delayDic)
			}
		} else {
			right.size--
			if right.top() == x {
				right.prune(delayDic)
			}
		}
		balance()
	}
	getMid := func() float64 {
		if k&1 > 0 {
			return float64(-left.top())
		}
		return float64(right.top()-left.top()) / 2
	}

	n := len(nums)
	for i := 0; i < k; i++ {
		inserts(nums[i])
	}
	result := make([]float64, 0)
	result = append(result, getMid())
	for i := k; i < n; i++ {
		inserts(nums[i])
		remove(nums[i-k])
		result = append(result, getMid())
	}
	return result
}

type PQueue struct {
	sort.IntSlice
	size  int
	minus bool
}

func (pq *PQueue) Push(v interface{}) {
	pq.IntSlice = append(pq.IntSlice, v.(int))
}

func (pq *PQueue) Pop() interface{} {
	q := pq.IntSlice
	v := pq.IntSlice[q.Len()-1]
	pq.IntSlice = q[:q.Len()-1]
	return v

}

func (pq *PQueue) push(x int) {
	pq.size++
	heap.Push(pq, x)
}

func (pq *PQueue) pop() int {
	pq.size--
	return heap.Pop(pq).(int)
}

func (pq *PQueue) top() int {
	if len(pq.IntSlice) > 0 {
		return pq.IntSlice[0]
	}
	return 0
}

func (pq *PQueue) prune(delay map[int]int) {
	for pq.Len() > 0 {
		x := pq.top()
		if pq.minus {
			x = -x
		}
		if sum, has := delay[x]; has {
			if sum > 1 {
				delay[x]--
			} else {
				delete(delay, x)
			}
			heap.Pop(pq)
		} else {
			break
		}
	}

}

func main() {
	medianSlidingWindow([]int{9, 7, 0, 3, 9, 8, 6, 5, 7, 6}, 2)
	//[9,7,0,3,9,8,6,5,7,6]
	//	2
	stack := &PQueue{}
	n := 10
	for i := 0; i < n; i++ {
		stack.push(i)
	}
	print(stack)
}
