package main

//739. 每日温度
//根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用0 来代替。
//
//例如，给定一个列表temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是[1, 1, 4, 2, 1, 1, 0, 0]。
//
//提示：气温 列表长度的范围是[1, 30000]。每个气温的值的均为华氏度，都是在[30, 100]范围内的整数。

//思路： 单调栈
// 区间，范围等 几乎是 单调栈算法
func dailyTemperatures(T []int) []int {
	n := len(T)
	result := make([]int, n)
	stack := make([]int, 0)
	stack = append(stack, 0)
	peek := 0
	for i := 1; i < n; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			peek = len(stack) - 1
			result[stack[peek]] = i - stack[peek]
			stack = stack[:peek]
		}
		stack = append(stack, i)
	}
	return result
}

func main() {
	println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
