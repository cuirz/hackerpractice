package main

import (
	"container/heap"
	"math"
)

//632. 最小区间
//你有 k 个升序排列的整数数组。找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。
//
//我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。
//
//示例 1:
//
//输入:[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
//输出: [20,24]
//解释:
//列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
//列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
//列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
//注意:
//
//给定的列表可能包含重复元素，所以在这里升序表示 >= 。
//1 <= k <= 3500
//-105 <= 元素的值 <= 105
//对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。

//思路  优先堆排序
// 每个区间选取最小值 入堆--优先堆排序，
// 测量堆里区间大小完后出堆，
// 出去的值属于哪个区间就从那个区间选取下一个再入堆,然后循环

func smallestRange(nums [][]int) []int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	n := len(nums)
	heapMax := math.MinInt32
	for i := 0; i < n; i++ {
		heap.Push(pq, Range{v: nums[i][0], x: i, y: 0})
		heapMax = max(heapMax, nums[i][0])
	}
	var result []int
	best := math.MaxInt32
	for pq.Len() == n {
		top := heap.Pop(pq).(Range)
		if heapMax-top.v < best {
			best = heapMax - top.v
			result = []int{top.v, heapMax}
		}
		i, j := top.x, top.y+1
		if j < len(nums[i]) {
			heap.Push(pq, Range{v: nums[i][j], x: i, y: j})
			heapMax = max(heapMax, nums[i][j])
		}
	}
	return result
}

type Range struct {
	v int
	x int
	y int
}

type PriorityQueue []Range

func (p PriorityQueue) Len() int           { return len(p) }
func (p PriorityQueue) Less(i, j int) bool { return p[i].v < p[j].v }
func (p PriorityQueue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PriorityQueue) Push(one interface{}) {
	*p = append(*p, one.(Range))
}
func (p *PriorityQueue) Pop() interface{} {
	if p.Len() > 0 {
		// heap的pop时把最小值和尾部替换过，所以最小值在尾部
		res := (*p)[p.Len()-1]
		*p = (*p)[:p.Len()-1]
		return res
	}
	return nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(smallestRange([][]int{
		//{1},
		{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30},
		//{4,98},
		//{1,9,100},
		//{4,99},
		//{11, 38, 83, 84, 84, 85, 88, 89, 89, 92}, {28, 61, 89}, {52, 77, 79, 80, 81}, {21, 25, 26, 26, 26, 27}, {9, 83, 85, 90}, {84, 85, 87}, {26, 68, 70, 71}, {36, 40, 41, 42, 45}, {-34, 21}, {-28, -28, -23, 1, 13, 21, 28, 37, 37, 38}, {-74, 1, 2, 22, 33, 35, 43, 45}, {54, 96, 98, 98, 99}, {43, 54, 60, 65, 71, 75}, {43, 46}, {50, 50, 58, 67, 69}, {7, 14, 15}, {78, 80, 89, 89, 90}, {35, 47, 63, 69, 77, 92, 94},
	}))
}
