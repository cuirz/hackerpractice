package main

//593. 有效的正方形
//给定2D空间中四个点的坐标p1,p2,p3和p4，如果这四个点构成一个正方形，则返回 true 。
//
//点的坐标pi 表示为 [xi, yi] 。输入 不是 按任何顺序给出的。
//
//一个 有效的正方形 有四条等边和四个等角(90度角)。
//
//
//
//示例 1:
//
//输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
//输出: True
//示例 2:
//
//输入：p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
//输出：false
//示例 3:
//
//输入：p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
//输出：true
//
//
//提示:
//
//p1.length == p2.length == p3.length == p4.length == 2
//-10^4<= xi, yi<= 10^4

//数学   方法1 任意3点都是等腰三角形  方法2 4个边 2个对角线 共2种长度


func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	sqrMagnitude := func(a, b []int) int {
		x := a[0] - b[0]
		y := a[1] - b[1]
		return x*x + y*y
	}
	hMap := map[int]int{}
	hMap[sqrMagnitude(p1, p2)]++
	hMap[sqrMagnitude(p1, p3)]++
	hMap[sqrMagnitude(p1, p4)]++
	hMap[sqrMagnitude(p2, p3)]++
	hMap[sqrMagnitude(p3, p4)]++
	hMap[sqrMagnitude(p2, p4)]++
	if len(hMap) != 2 || hMap[0] > 0 {
		return false
	}
	var array []int
	for _, v := range hMap {
		array = append(array, v)
	}
	if array[0] > array[1] {
		return array[0] == 2*array[1]
	} else {
		return array[1] == 2*array[0]
	}
}
