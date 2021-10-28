package main

//869. 重新排序得到 2 的幂
//给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。
//
//如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。
//
//
//
//示例 1：
//
//输入：1
//输出：true
//示例 2：
//
//输入：10
//输出：false
//示例 3：
//
//输入：16
//输出：true
//示例 4：
//
//输入：24
//输出：false
//示例 5：
//
//输入：46
//输出：true
//
//
//提示：
//
//1 <= N <= 10^9

func reorderedPowerOf2(n int) bool {
	var cntDigits func(x int) [10]int
	cntDigits = func(x int) (res [10]int) {
		for x > 0 {
			res[x%10]++
			x /= 10
		}
		return
	}
	mp := map[[10]int]bool{}
	for i := 1; i <= 1e9; i <<= 1 {
		mp[cntDigits(i)] = true
	}
	return mp[cntDigits(n)]
}

func main() {
	println(reorderedPowerOf2(10))
}
