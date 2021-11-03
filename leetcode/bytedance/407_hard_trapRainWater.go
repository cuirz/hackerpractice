package main

import "container/heap"

//407. 接雨水 II
//给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
//
//
//
//示例 1:
//
//
//
//输入: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
//输出: 4
//解释: 下雨后，雨水将会被上图蓝色的方块中。总的接雨水量为1+2+1=4。
//示例 2:
//
//
//
//输入: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
//输出: 10
//
//
//提示:
//
//m == heightMap.length
//n == heightMap[i].length
//1 <= m, n <= 200
//0 <= heightMap[i][j] <= 2 * 10^4

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	n := len(heightMap)
	m := len(heightMap[0])
	sum := 0

	bucket := &PriorityQueue{}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				visited[i][j] = true
				heap.Push(bucket, Rain{x: i, y: j, v: heightMap[i][j]})
			}
		}
	}

	dir := []int{-1, 0, 1, 0, -1}
	for bucket.Len() > 0 {
		head := heap.Pop(bucket).(Rain)
		for k := 0; k < 4; k++ {
			x := head.x + dir[k]
			y := head.y + dir[k+1]
			if x >= 0 && y >= 0 && x < n && y < m && !visited[x][y] {
				if heightMap[x][y] < head.v {
					sum += head.v - heightMap[x][y]
				}
				heap.Push(bucket, Rain{x: x, y: y, v: max(head.v, heightMap[x][y])})
				visited[x][y] = true
			}

		}
	}
	return sum
}

type Rain struct {
	x int
	y int
	v int
}

type PriorityQueue []Rain

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].v < q[j].v
}
func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(r interface{}) {
	*q = append(*q, r.(Rain))
}

func (q *PriorityQueue) Pop() interface{} {
	a := *q
	v := a[len(a)-1]
	*q = a[:len(a)-1]
	return v
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
