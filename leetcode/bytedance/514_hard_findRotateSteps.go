package main

import "math"

//514. 自由之路
//视频游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。
//
//给定一个字符串ring，表示刻在外环上的编码；给定另一个字符串key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
//
//最初，ring的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使key的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完key中的所有字符。
//
//旋转ring拼出 key 字符key[i]的阶段中：
//
//您可以将ring顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串ring的一个字符与 12:00 方向对齐，并且这个字符必须等于字符key[i] 。
//如果字符key[i]已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作1 步。按完之后，您可以开始拼写key的下一个字符（下一阶段）, 直至完成所有拼写。

func findRotateSteps(ring string, key string) int {
	n := len(ring)
	dic := make(map[byte][]int)
	dis := func(s, d int) int {
		distance := abs(s - d)
		return min(distance, abs(n-distance))
	}
	for i, _ := range ring {
		dic[ring[i]] = append(dic[ring[i]], i)
	}

	dp := make([][]int, len(key))
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for _, v := range dic[key[0]] {
		dp[0][v] = min(v, n-v) + 1
	}
	for i := 1; i < len(key); i++ {
		for _, j := range dic[key[i]] {
			for _, k := range dic[key[i-1]] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+dis(j, k)+1)
			}
		}
	}
	return min(dp[len(key)-1]...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x ...int) int {
	y := x[0]
	for _, v := range x[1:] {
		if v < y {
			y = v
		}
	}
	return y

}

func main() {
	println(findRotateSteps("godding", "godding"))
}
