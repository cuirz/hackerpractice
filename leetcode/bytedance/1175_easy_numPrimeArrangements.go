package main

//1175. 质数排列
//请你帮忙给从 1 到 n的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；你需要返回可能的方案总数。
//
//让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。
//
//由于答案可能会很大，所以请你返回答案 模 mod10^9 + 7之后的结果即可。
//
//
//
//示例 1：
//
//输入：n = 5
//输出：12
//解释：举个例子，[1,2,5,4,3] 是一个有效的排列，但 [5,2,3,4,1] 不是，因为在第二种情况里质数 5 被错误地放在索引为 1 的位置上。
//示例 2：
//
//输入：n = 100
//输出：682289015
//
//提示：
//
//1 <= n <= 100

//计算质数    判断 n是否质数  能否被整除 range [2,sqart(n))的元素
//排列总数量  = （质数数量）！x (非质数薯榔）！

func numPrimeArrangements(n int) int {
	const mod = 1e9 + 7
	isPrime := func(num int) bool {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	factorial := func(num int) int {
		total := 1
		for i := 1; i <= num; i++ {
			total = total * i % mod
		}
		return total
	}
	cnt := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return factorial(cnt) * factorial(n-cnt) % mod
}
