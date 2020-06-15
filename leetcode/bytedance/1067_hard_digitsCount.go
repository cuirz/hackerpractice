package main

//1067. 范围内的数字计数
//给定一个在 0 到 9 之间的整数 d，和两个正整数 low 和 high 分别作为上下界。返回 d 在 low 和 high 之间的整数中出现的次数，包括边界 low 和 high。
//
//
//
//示例 1：
//
//输入：d = 1, low = 1, high = 13
//输出：6
//解释：
//数字 d=1 在 1,10,11,12,13 中出现 6 次。注意 d=1 在数字 11 中出现两次。
//示例 2：
//
//输入：d = 3, low = 100, high = 250
//输出：35
//解释：
//数字 d=3 在 103,113,123,130,131,...,238,239,243 出现 35 次。
//
//
//提示：
//
//0 <= d <= 9
//1 <= low <= high <= 2×10^8

func digitsCount(d int, low int, high int) int {
	var count func(d, n int) int
	count = func(d, n int) int {
		result := 0
		k := n
		for i := 1; i <= n; i *= 10 {
			k = n / i
			hi := k / 10
			if d == 0 {
				if hi != 0 {
					hi--
				} else {
					break
				}
			}
			result += hi * i
			println(hi * i)
			// 当前位的数字。
			cur := k % 10
			//div := i * 10
			//result += hi*i min(max(n - k*i + 1, 0), i)
			//result += hi*i + min(max(n%div-i+1, 0), i)

			if cur > d {
				result += i
			} else if cur == d {
				// n - k * i 为低位的数字。
				println(i, n-k*i+1)
				result += n - k*i + 1
			}
		}
		return result
	}

	return count(d, high) - count(d, low-1)

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//println(digitsCount(5,1,10000))
	println(digitsCount(1, 1, 111))
}
