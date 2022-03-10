package main

//240. 搜索二维矩阵 II
//编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//给定 target=5，返回true。
//
//给定target=20，返回false。

//思路 二分法
// 先判断哪个行

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n < 1 {
		return false
	}
	m := len(matrix[0])
	if m < 1 {
		return false
	}
	binary := func(l, r, index int) int {
		mid := 0
		for l < r {
			mid = l + (r-l)>>1
			if target < matrix[index][l] {
				return l - 1
			} else if target >= matrix[index][r] {
				return r
			}
			mv := matrix[index][mid]
			if target < mv {
				r = mid
			} else if target > mv {
				l = mid + 1
			} else {
				return mid
			}
		}
		return mid
	}

	if target > matrix[n-1][m-1] || target < matrix[0][0] {
		return false
	}

	index := m - 1
	for i := 0; i < n; i++ {
		if target <= matrix[i][m-1] {
			index = binary(0, index, i)
			if index < 0 {
				return false
			}
			if target == matrix[i][index] {
				return true
			}
		}
		if target < matrix[i][0] {
			return false
		}
	}
	return false
}

//思路 指针 斜着走
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n < 1 {
		return false
	}
	m := len(matrix[0])
	if m < 1 {
		return false
	}
	col, row := 0, n-1
	for row >= 0 && col < m {
		if target > matrix[row][col] {
			col++
		} else if target < matrix[row][col] {
			row--
		} else {
			return true
		}
	}
	return false
}

func main() {
	println(searchMatrix([][]int{
		{2, 3, 6, 6, 10}, {5, 9, 12, 17, 19}, {10, 14, 17, 20, 20}, {15, 17, 20, 24, 24}, {20, 20, 25, 26, 29},
	}, 4))
}
