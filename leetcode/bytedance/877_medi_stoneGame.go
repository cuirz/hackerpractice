package main

//亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。
//
//游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。
//
//亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。
//
//假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。
//
//
//
//示例：
//
//输入：[5,3,4,5]
//输出：true
//解释：
//亚历克斯先开始，只能拿前 5 颗或后 5 颗石子 。
//假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
//如果李拿走前 3 颗，那么剩下的是 [4,5]，亚历克斯拿走后 5 颗赢得 10 分。
//如果李拿走后 5 颗，那么剩下的是 [3,4]，亚历克斯拿走后 4 颗赢得 9 分。
//这表明，取前 5 颗石子对亚历克斯来说是一个胜利的举动，所以我们返回 true 。
//
//
//提示：
//
//2 <= piles.length <= 500
//piles.length 是偶数。
//1 <= piles[i] <= 500
//sum(piles) 是奇数。
func stoneGame(piles []int) bool {
	dp := make(map[[2]int][2]int)
	var search func(l, r int) [2]int
	search = func(l, r int) [2]int {
		if v, ok := dp[[2]int{l, r}]; ok {
			return v
		}
		if l == r {
			res := [2]int{l, 0}
			dp[[2]int{l, r}] = res
			return res
		} else if r-l == 1 {
			res := maxmin(piles[l], piles[r])
			dp[[2]int{l, r}] = res
			return res
		}
		left := search(l+1, r)
		right := search(l, r-1)
		first := max(piles[l]+left[1], piles[r]+right[1])
		res := [2]int{first, piles[l] + left[0] + left[1] - first}
		dp[[2]int{l, r}] = res
		return res
	}
	n := len(piles)
	result := search(0, n-1)
	return result[0] > result[1]
}

func maxmin(x, y int) [2]int {
	if x > y {
		return [2]int{x, y}
	}
	return [2]int{y, x}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
