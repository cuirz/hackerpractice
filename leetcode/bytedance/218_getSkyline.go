package main

// 分治法
func getSkyline(buildings [][]int) [][]int {
	if len(buildings) < 1 {
		return [][]int{}
	}
	if len(buildings) == 1 {
		// 左上角 和 右下角
		return [][]int{{buildings[0][0], buildings[0][2]}, {buildings[0][1], 0}}
	}

	mid := len(buildings) / 2
	left := getSkyline(buildings[:mid])
	right := getSkyline(buildings[mid:])
	return merge(left, right)

}

func merge(left, right [][]int) [][]int {
	// 遍历到的 当前左边建筑高度和右边建筑高度
	lh, rh := 0, 0
	l, r := 0, 0
	var res [][]int
	var point []int
	for l < len(left) && r < len(right) {
		// 三种情况
		if left[l][0] < right[r][0] {
			lh = left[l][1]
			point = []int{left[l][0], max(lh, rh)}
			l++
		} else if left[l][0] > right[r][0] {
			rh = right[r][1]
			point = []int{right[r][0], max(lh, rh)}
			r++
		} else {
			lh = left[l][1]
			rh = right[r][1]
			point = []int{left[l][0], max(lh, rh)}
			l++
			r++
		}
		// 比较之前的高度，只获取不同时的 点
		if len(res) == 0 || point[1] != res[len(res)-1][1] {
			res = append(res, point)
		}
	}
	// 把剩余的也加进去
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 扫描法

func main() {

}
