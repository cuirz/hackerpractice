package main

import (
	"container/heap"
	"sort"
)

//630. 课程表 III
//这里有 n 门不同的在线课程，他们按从 1 到 n 编号。每一门课程有一定的持续上课时间（课程时间）t 以及关闭时间第 d 天。一门课要持续学习 t 天直到第 d 天时要完成，你将会从第 1 天开始。
//
//给出 n 个在线课程用 (t, d) 对表示。你的任务是找出最多可以修几门课。
//
//
//
//示例：
//
//输入: [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
//输出: 3
//解释:
//这里一共有 4 门课程, 但是你最多可以修 3 门:
//首先, 修第一门课时, 它要耗费 100 天，你会在第 100 天完成, 在第 101 天准备下门课。
//第二, 修第三门课时, 它会耗费 1000 天，所以你将在第 1100 天的时候完成它, 以及在第 1101 天开始准备下门课程。
//第三, 修第二门课时, 它会耗时 200 天，所以你将会在第 1300 天时完成它。
//第四门课现在不能修，因为你将会在第 3300 天完成它，这已经超出了关闭日期。
//提示:
//
//整数 1 <= d, t, n <= 10,000 。
//你不能同时修两门课程。

// 优先队列，贪心算法
// 根据结束时间来排序，选取课程时，t1+t2 <= d2 来选课，当 t1+t2 > d2超出时间时，
// 贪婪算法，剔除最长课程 增加可接受容量

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})
	prioQueue := &PQueue{}
	heap.Init(prioQueue)
	sum, temp := 0, 0
	for i := 0; i < len(courses); i++ {
		temp = sum + courses[i][0]
		if temp <= courses[i][1] {
			sum = temp
			heap.Push(prioQueue, courses[i][0])
		} else if prioQueue.Len() > 0 && prioQueue.top() > courses[i][0] {
			pop := heap.Pop(prioQueue)
			sum += courses[i][0] - pop.(int)
			heap.Push(prioQueue, courses[i][0])
		}
	}
	return prioQueue.Len()
}

type PQueue []int

func (pq PQueue) Less(i, j int) bool {
	return pq[i] > pq[j]

}
func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQueue) Pop() interface{} {
	if pq.Len() > 0 {
		item := (*pq)[pq.Len()-1]
		*pq = (*pq)[:pq.Len()-1]
		return item
	}
	return 0
}
func (pq PQueue) top() int {
	if pq.Len() > 0 {
		return pq[0]
	}
	return 0
}
func (pq PQueue) Len() int {
	return len(pq)
}

func main() {
	//[[100,200],[200,1300],[1000,1250],[2000,3200]]
	//scheduleCourse([][]int{{100,200},{200,1300},{1000,1250},{2000,3200}})
	//scheduleCourse([][]int{{5, 5}, {4, 6}, {2, 6}})
	//scheduleCourse([][]int{{5, 5}, {4, 6}, {2, 6}})
	//[[10,12],[6,15],[1,12],[3,20],[10,19]]
	//[[7,17],[3,12],[10,20],[9,10],[5,20],[10,19],[4,18]]
	//println(scheduleCourse([][]int{{10,12},{6,15},{1,12},{3,20},{10,19}}))
	println(scheduleCourse([][]int{{7, 17}, {3, 12}, {10, 20}, {9, 10}, {5, 20}, {10, 19}, {4, 18}}))

	prioQueue := &PQueue{}
	heap.Init(prioQueue)
	heap.Push(prioQueue, 9)
	heap.Push(prioQueue, 3)
	heap.Push(prioQueue, 10)
	heap.Push(prioQueue, 20)
	heap.Push(prioQueue, 5)
	x := heap.Pop(prioQueue)
	heap.Push(prioQueue, 30)
	heap.Push(prioQueue, 12)
	x = heap.Pop(prioQueue)
	println(prioQueue.top())
	println(x.(int))

}
