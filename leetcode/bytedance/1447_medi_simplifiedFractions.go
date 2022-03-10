package main

import "strconv"

//1447. 最简分数
//给你一个整数n，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于n的 最简分数。分数可以以 任意顺序返回。
//
//
//
//示例 1：
//
//输入：n = 2
//输出：["1/2"]
//解释："1/2" 是唯一一个分母小于等于 2 的最简分数。
//示例 2：
//
//输入：n = 3
//输出：["1/2","1/3","2/3"]
//示例 3：
//
//输入：n = 4
//输出：["1/2","1/3","1/4","2/3","3/4"]
//解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。
//示例 4：
//
//输入：n = 1
//输出：[]
//
//
//提示：
//
//1 <= n <= 100
//欧几里得算法，辗转相除法

func simplifiedFractions(n int) []string {
	var gcd func(i, j int) int
	gcd = func(i, j int) int {
		if j == 0 {
			return i
		}
		return gcd(j, i%j)
	}
	var result []string
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				result = append(result, strconv.Itoa(i)+"/"+strconv.Itoa(j))
			}
		}
	}
	return result
}
