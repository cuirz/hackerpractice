package main

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//示例 1:
//
//输入:
//matrix = [
//[1,   3,  5,  7],
//[10, 11, 16, 20],
//[23, 30, 34, 50]
//]
//target = 3
//输出: true
//示例 2:
//
//输入:
//matrix = [
//[1,   3,  5,  7],
//[10, 11, 16, 20],
//[23, 30, 34, 50]
//]
//target = 13
//输出: false

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n < 1 {
		return false
	}
	if len(matrix[0]) < 1 {
		return false
	}
	t := 0
	b := n
	pre, mid := -1, 0
	for t < b {
		mid = t + (b-t)>>1
		if pre == mid {
			break
		}
		pre = mid
		if target < matrix[mid][0] {
			b = mid
		} else if target > matrix[mid][0] {
			t = mid
		} else {
			return true
		}
	}
	l, r := 0, len(matrix[0])
	befor := -1
	for l < r {
		mid = l + (r-l)>>1
		if befor == mid {
			break
		}
		befor = mid
		if target == matrix[pre][mid] {
			return true
		} else if target < matrix[pre][mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return false

}

func main() {
	println(searchMatrix([][]int{{1, 2}, {3, 4}}, 2))
}
