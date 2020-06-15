package main

//351. 安卓系统手势解锁
// 中等
//我们都知道安卓有个手势解锁的界面，是一个 3 x 3 的点所绘制出来的网格。
//
//给你两个整数，分别为 ​​m 和 n，其中 1 ≤ m ≤ n ≤ 9，那么请你统计一下有多少种解锁手势，是至少需要经过 m 个点，但是最多经过不超过 n 个点的。
//先来了解下什么是一个有效的安卓解锁手势:
//
//每一个解锁手势必须至少经过 m 个点、最多经过 n 个点。
//解锁手势里不能设置经过重复的点。
//假如手势中有两个点是顺序经过的，那么这两个点的手势轨迹之间是绝对不能跨过任何未被经过的点。
//经过点的顺序不同则表示为不同的解锁手势。
//
//解释:
//
//| 1 | 2 | 3 |
//| 4 | 5 | 6 |
//| 7 | 8 | 9 |
//无效手势：4 - 1 - 3 - 6
//连接点 1 和点 3 时经过了未被连接过的 2 号点。
//
//无效手势：4 - 1 - 9 - 2
//连接点 1 和点 9 时经过了未被连接过的 5 号点。
//
//有效手势：2 - 4 - 1 - 3 - 6
//连接点 1 和点 3 是有效的，因为虽然它经过了点 2 ，但是点 2 在该手势中之前已经被连过了。
//
//有效手势：6 - 5 - 4 - 1 - 9 - 2
//连接点 1 和点 9 是有效的，因为虽然它经过了按键 5 ，但是点 5 在该手势中之前已经被连过了。
//
//
//
//示例:
//
//输入: m = 1，n = 1
//输出: 9
//

func numberOfPatterns(m int, n int) int {
	var disable [10]bool
	sum := 0
	nums := map[int]map[int]int{
		1: {3: 2, 9: 5, 7: 4},
		2: {8: 5},
		3: {1: 2, 7: 5, 9: 6},
		4: {6: 5},
		6: {4: 5},
		7: {1: 4, 3: 5, 9: 8},
		8: {2: 5},
		9: {3: 6, 1: 5, 7: 8},
	}
	var search func(index, cnt int)
	search = func(index, cnt int) {
		if cnt >= m && cnt <= n {
			sum++
		}
		if cnt == n {
			return
		}
		//获取阻挡路径
		dic, has := nums[index]
		for i := 1; i <= 9; i++ {
			// 除了自己以外
			if index != i {
				if disable[i] {
					continue
				}
				// 有必须经过的，就检查
				if has && dic[i] > 0 && !disable[dic[i]] {
					continue
				}
				disable[i] = true
				search(i, cnt+1)
				disable[i] = false
			}
		}
	}
	for i := 1; i <= 9; i++ {
		disable[i] = true
		search(i, 1)
		disable[i] = false
	}
	return sum
}

func main() {
	println(numberOfPatterns(1, 3))

	dic := map[int]int{
		1: 1,
	}
	println(dic[1])
}
