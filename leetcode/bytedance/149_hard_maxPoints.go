package main

//149. 直线上最多的点数
//给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
//输入：points = [[1,1],[2,2],[3,3]]
//输出：3
//输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
//输出：4
//提示：
//
//1 <= points.length <= 300
//points[i].length == 2
//-10^4 <= xi, yi <= 10^4
//points 中的所有点 互不相同

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2{
		return n
	}
	var result int

	for i, p := range points {
		dic := make(map[int]int)
		if result > n/2 || result > n-i {
			break
		}
		for j := i + 1; j < n; j++ {
			x, y := p[0]-points[j][0], p[1]-points[j][1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			dic[y+x*20001] ++
		}
		for _, c := range dic {
			result = max(result, c+1)
		}
	}
	return result
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
