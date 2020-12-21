package main

//204. 计数质数
//统计所有小于非负整数 n 的质数的数量。
//
// 
//
//示例 1：
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//示例 2：
//
//输入：n = 0
//输出：0
//示例 3：
//
//输入：n = 1
//输出：0

// 埃氏筛
func countPrimes(n int) int {
	isprime := make([]bool, n)
	for i := 0; i < n; i++ {
		isprime[i] = true
	}
	count := 0
	for i := 2; i < n; i++ {
		if isprime[i] {
			count++
			for j := i * i; j < n; j += i {
				isprime[j] = false
			}
		}
	}
	return count
}

func main() {
	println(countPrimes(10))
}
