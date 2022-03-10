package main

//在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为days的数组给出。每一项是一个从1到365的整数。
//
//火车票有三种不同的销售方式：
//
//一张为期一天的通行证售价为costs[0] 美元；
//一张为期七天的通行证售价为costs[1] 美元；
//一张为期三十天的通行证售价为costs[2] 美元。
//通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张为期 7 天的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。
//
//返回你想要完成在给定的列表days中列出的每一天的旅行所需要的最低消费。

func mincostTickets(days []int, costs []int) int {
	result := [366]int{}
	match := map[int]bool{}
	for _, d := range days {
		match[d] = true
	}

	var dp func(day int) int
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if result[day] > 0 {
			return result[day]
		}
		if match[day] {
			result[day] = min(min(dp(day+1)+costs[0], dp(day+7)+costs[1]), dp(day+30)+costs[2])
		} else {
			result[day] = dp(day + 1)
		}
		return result[day]
	}
	return dp(1)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {

}
