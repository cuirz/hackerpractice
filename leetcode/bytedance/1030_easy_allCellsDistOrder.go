package main

import "sort"

//1030. 距离顺序排列矩阵单元格
//给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。
//
//另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。
//
//返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int,0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			result = append(result,[]int{i,j})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		x,y := result[i],result[j]
		return (abs(x[0]-r0)+abs(x[1]-c0)) < (abs(y[0]-r0)+abs(y[1]-c0))
	})

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main(){
	println(allCellsDistOrder(2,3,1,2))
}
