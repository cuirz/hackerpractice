package main

//386.字典序排数
//给定一个整数 n, 返回从 1 到 n 的字典顺序。
//
//例如，
//
//给定 n=13，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。
//
//请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。

func lexicalOrder(n int) []int {
	result := make([]int, n)
	index := 0

	var back func(h int)
	back = func(h int) {
		for l := 0; l <= 9; l++ {
			v := h + l
			if v > n {
				break
			}
			result[index] = v
			index++
			back(v * 10)
		}
	}

	for i := 1; i <= 9; i++ {
		if i > n {
			break
		}
		result[index] = i
		index++
		back(i * 10)
	}

	return result

}

func main() {
	lexicalOrder(130)
}
