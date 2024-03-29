package main

//1128. 等价多米诺骨牌对的数量
//给你一个由一些多米诺骨牌组成的列表dominoes。
//
//如果其中某一张多米诺骨牌可以通过旋转 0度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
//
//形式上，dominoes[i] = [a, b]和dominoes[j] = [c, d]等价的前提是a==c且b==d，或是a==d 且b==c。
//
//在0 <= i < j < dominoes.length的前提下，找出满足dominoes[i] 和dominoes[j]等价的骨牌对 (i, j) 的数量。
//
//
//
//示例：
//
//输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
//输出：1
//
//
//提示：
//
//1 <= dominoes.length <= 40000
//1 <= dominoes[i][j] <= 9

//思路
func numEquivDominoPairs(dominoes [][]int) int {
	count := make([]int, 100)
	result := 0
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		val := v[0]*10 + v[1]
		result += count[val]
		count[val]++
	}
	return result
}
