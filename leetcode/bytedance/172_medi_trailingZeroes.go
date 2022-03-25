package main

//172. 阶乘后的零
//给定一个整数 n ，返回 n! 结果中尾随零的数量。
//
//提示n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
//
//
//
//示例 1：
//
//输入：n = 3
//输出：0
//解释：3! = 6 ，不含尾随 0
//示例 2：
//
//输入：n = 5
//输出：1
//解释：5! = 120 ，有一个尾随 0
//示例 3：
//
//输入：n = 0
//输出：0
//
//
//提示：
//
//0 <= n <= 10^4
//
//
//进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？

//思路
// 每隔 5个 有质子5 1个
// 每隔 5*5个 有质子5 2个
// 每隔 5*5*5个 有质子 3个
func trailingZeroes(n int) int {
	var ans int
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}
func trailingZeroes2(n int) int {
	var ans int
	for i := 5; i < n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
	}
	return ans
}

func main() {
	println(trailingZeroes(300))
}
