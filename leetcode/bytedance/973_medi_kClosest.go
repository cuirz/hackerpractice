package main

import (
	"sort"
)

//973. 最接近原点的 K 个点
//我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
//
//（这里，平面上两点之间的距离是欧几里德距离。）
//
//你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

func kClosest(points [][]int, K int) [][]int {

	sort.Slice(points, func(i, j int) bool {
		return (points[i][0]*points[i][0]+points[i][1]*points[i][1]) <
			(points[j][0]*points[j][0]+points[j][1]*points[j][1])
	})

	return points[:K]
}
