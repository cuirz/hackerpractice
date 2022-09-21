package main

import (
	"sort"
)

//850. 矩形面积 II
//我们给出了一个（轴对齐的）二维矩形列表rectangles。 对于rectangle[i] = [x1, y1, x2, y2]，其中（x1，y1）是矩形i左下角的坐标，(xi1, yi1)是该矩形 左下角 的坐标，(xi2, yi2)是该矩形右上角 的坐标。
//
//计算平面中所有rectangles所覆盖的 总面积 。任何被两个或多个矩形覆盖的区域应只计算 一次 。
//
//返回 总面积 。因为答案可能太大，返回10^9+ 7 的模。
//
//
//
//示例 1：
//
//
//
//输入：rectangles = [[0,0,2,2],[1,0,2,3],[1,0,3,1]]
//输出：6
//解释：如图所示，三个矩形覆盖了总面积为6的区域。
//从(1,1)到(2,2)，绿色矩形和红色矩形重叠。
//从(1,0)到(2,3)，三个矩形都重叠。
//示例 2：
//
//输入：rectangles = [[0,0,1000000000,1000000000]]
//输出：49
//解释：答案是 1018 对 (109 + 7) 取模的结果， 即 49 。
//
//
//提示：
//
//1 <= rectangles.length <= 200
//rectanges[i].length = 4
//0 <= xi1, yi1, xi2, yi2<= 10^9
//矩形叠加覆盖后的总面积不会超越2^63 - 1，这意味着可以用一个64 位有符号整数来保存面积结果。

//思路 扫描线+离散化
//可以使用 线段树提高搜索速度
type seg []struct {
	cover, len, maxLen int
}

func (t seg) init(hBound []int, idx, l, r int) {
	if l == r {
		t[idx].maxLen = hBound[l] - hBound[l-1]
		return
	}
	mid := (l + r) / 2
	t.init(hBound, idx*2, l, mid)
	t.init(hBound, idx*2+1, mid+1, r)
	t[idx].maxLen = t[idx*2].maxLen + t[idx*2+1].maxLen
}

func (t seg) update(idx, l, r, ul, ur, diff int) {
	if l > ur || r < ul {
		return
	}
	if ul <= l && r <= ur {
		t[idx].cover += diff
		t.pushUp(idx, l, r)
		return
	}
	mid := (l + r) / 2
	t.update(idx*2, l, mid, ul, ur, diff)
	t.update(idx*2+1, mid+1, r, ul, ur, diff)
	t.pushUp(idx, l, r)
}

func (t seg) pushUp(idx, l, r int) {
	if t[idx].cover > 0 {
		t[idx].len = t[idx].maxLen
	} else if l == r {
		t[idx].len = 0
	} else {
		t[idx].len = t[idx*2].len + t[idx*2+1].len
	}
}

func rectangleArea(rectangles [][]int) int {
	n := len(rectangles) * 2
	hBound := make([]int, 0, n)
	for _, v := range rectangles {
		hBound = append(hBound, v[1], v[3])
	}
	sort.Ints(hBound)
	m := 0
	for _, v := range hBound[1:] {
		if hBound[m] != v {
			m++
			hBound[m] = v
		}
	}
	hBound = hBound[:m+1]
	type tuple struct {
		x, i, d int
	}
	sweep := make([]tuple, 0, n)
	for i, v := range rectangles {
		sweep = append(sweep, tuple{v[0], i, 1}, tuple{v[2], i, -1})
	}
	sort.Slice(sweep, func(i, j int) bool {
		return sweep[i].x < sweep[j].x
	})

	t := make(seg, m*4)
	t.init(hBound, 1, 1, m)

	var result int
	for i := 0; i < n; i++ {
		j := i
		for j+1 < n && sweep[j+1].x == sweep[i].x {
			j++
		}
		if j+1 == n {
			break
		}
		for k := i; k <= j; k++ {
			index, diff := sweep[k].i, sweep[k].d
			left := sort.SearchInts(hBound, rectangles[index][1]) + 1
			right := sort.SearchInts(hBound, rectangles[index][3])
			t.update(1, 1, m, left, right, diff)
		}
		result += t[1].len * (sweep[j+1].x - sweep[j].x)

		i = j
	}
	return result % (1e9 + 7)
}

func rectangleArea2(rectangles [][]int) int {
	n := len(rectangles) * 2
	hBound := make([]int, 0, n)
	for _, v := range rectangles {
		hBound = append(hBound, v[1], v[3])
	}
	sort.Ints(hBound)
	m := 0
	for _, v := range hBound[1:] {
		if hBound[m] != v {
			m++
			hBound[m] = v
		}
	}
	hBound = hBound[:m+1]
	type tuple struct {
		x, i, d int
	}
	sweep := make([]tuple, 0, n)
	for i, v := range rectangles {
		sweep = append(sweep, tuple{v[0], i, 1}, tuple{v[2], i, -1})
	}
	sort.Slice(sweep, func(i, j int) bool {
		return sweep[i].x < sweep[j].x
	})
	seg := make([]int, m)
	var result int
	for i := 0; i < n; i++ {
		j := i
		for j+1 < n && sweep[j+1].x == sweep[i].x {
			j++
		}
		if j+1 == n {
			break
		}
		for k := i; k <= j; k++ {
			index, diff := sweep[k].i, sweep[k].d
			left, right := rectangles[index][1], rectangles[index][3]
			for x := 0; x < m; x++ {
				if left <= hBound[x] && hBound[x+1] <= right {
					seg[x] += diff
				}
			}
		}
		cover := 0
		for k := 0; k < m; k++ {
			if seg[k] > 0 {
				cover += hBound[k+1] - hBound[k]
			}
		}
		result += cover * (sweep[j+1].x - sweep[j].x)

		i = j
	}
	return result % (1e9 + 7)
}
