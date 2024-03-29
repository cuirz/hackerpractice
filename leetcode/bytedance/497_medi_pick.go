package main

import (
	"math/rand"
	"sort"
)

//497. 非重叠矩形中的随机点
//给定一个由非重叠的轴对齐矩形的数组 rects ，其中 rects[i] = [ai, bi, xi, yi] 表示 (ai, bi) 是第 i 个矩形的左下角点，(xi, yi) 是第 i 个矩形的右上角角点。设计一个算法来随机挑选一个被某一矩形覆盖的整数点。矩形周长上的点也算做是被矩形覆盖。所有满足要求的点必须等概率被返回。
//
//在一个给定的矩形覆盖的空间内任何整数点都有可能被返回。
//
//请注意，整数点是具有整数坐标的点。
//
//实现Solution类:
//
//Solution(int[][] rects)用给定的矩形数组rects 初始化对象。
//int[] pick()返回一个随机的整数点 [u, v] 在给定的矩形所覆盖的空间内。
//
//
//示例 1：
//
//
//
//输入:
//["Solution","pick","pick","pick","pick","pick"]
//[[[[-2,-2,-1,-1],[1,0,3,0]]],[],[],[],[],[]]
//输出:
//[null,[-1,-2],[2,0],[-2,-1],[3,0],[-2,-2]
//
//解释：
//Solution solution = new Solution([[-2, -2, 1, 1], [2, 2, 4, 6]]);
//solution.pick(); // 返回 [1, -2]
//solution.pick(); // 返回 [1, -1]
//solution.pick(); // 返回 [-1, -2]
//solution.pick(); // 返回 [-2, -2]
//solution.pick(); // 返回 [0, 0]
//
//
//提示：
//
//1 <= rects.length <= 100
//rects[i].length == 4
//-10^9<= ai< xi<= 10^9
//-10^9<= bi< yi<= 10^9
//xi- ai<= 2000
//yi- bi<= 2000
//所有的矩形不重叠。
//pick 最多被调用10^4次。

type Solution struct {
	Rects  [][]int
	PreSum []int
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
	}
	return Solution{rects, sum}
}

func (this *Solution) Pick() []int {
	rate := rand.Intn(this.PreSum[len(this.PreSum)-1])
	index := sort.SearchInts(this.PreSum, rate+1) - 1
	r := this.Rects[index]
	a, b, y := r[0], r[1], r[3]
	da := (rate - this.PreSum[index]) / (y - b + 1)
	db := (rate - this.PreSum[index]) % (y - b + 1)
	return []int{a + da, b + db}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
