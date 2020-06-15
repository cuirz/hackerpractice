package main

import "sort"

//42 接雨水

//1-D的接雨水问题有一种解法是从左右两边的边界往中间不断进行收缩，收缩的过程中，对每个坐标（一维坐标）能接的雨水进行求解

func trap(height []int) int {
	left, right := 0, len(height)
	lmax, rmax := 0, 0
	sum := 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] > lmax {
				lmax = height[left]
			} else {
				sum += lmax - height[left]
			}
			left++
		} else {
			if height[right] > rmax {
				rmax = height[right]
			} else {
				sum += rmax - height[right]
			}
			right--
		}
	}
	return sum
}

//2-D的接雨水问题的边界不再是线段的两个端点，而是矩形的一周，所以我们用优先队列维护所有边界点，收缩时，也不仅仅只有左右两个方向，而是上下左右四个方向，并且维护一个visit的数组，记录哪些坐标已经被访问过，不然会造成重复求解。
type Rain struct {
	x int
	y int
	v int
}

type PriorityQueue []*Rain

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].v < q[j].v
}
func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(r *Rain) {
	*q = append(*q, r)
}

func (q *PriorityQueue) Pop() *Rain {
	if q.Len() > 0 {
		item := (*q)[0]
		*q = (*q)[:q.Len()-1]
		return item
	}
	return nil
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	n := len(heightMap)
	m := len(heightMap[0])
	sum := 0

	bucket := make(PriorityQueue, 0)
	var visited [][]bool
	for i := 0; i < n; i++ {
		visited = append(visited, make([]bool, m))
		for j := 0; j < m; j++ {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				visited[i][j] = true
				bucket.Push(&Rain{x: i, y: j, v: heightMap[i][j]})
			}
		}
	}
	sort.Sort(bucket)

	// 方向压缩成一维
	dir := []int{-1, 0, 1, 0, -1}
	for bucket.Len() > 0 {
		head := bucket.Pop()
		for k := 0; k < 4; k++ {
			x := head.x + dir[k]
			y := head.y + dir[k+1]
			if x >= 0 && y >= 0 && x < m && y < n && !visited[x][y] {
				if heightMap[x][y] < head.v {
					sum += head.v - heightMap[x][y]
				}
				bucket.Push(&Rain{x: x, y: y, v: heightMap[x][y]})
				sort.Sort(bucket)
				visited[x][y] = true
			}

		}
	}
	return sum
}

type User struct {
}

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
	//println(trap([]int{}))

	println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}))
}
