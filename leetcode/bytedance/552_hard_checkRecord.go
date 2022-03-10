package main

//552. 学生出勤记录 II
//可以用字符串表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//'A'：Absent，缺勤
//'L'：Late，迟到
//'P'：Present，到场
//如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
//按 总出勤 计，学生缺勤（'A'）严格 少于两天。
//学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
//给你一个整数 n ，表示出勤记录的长度（次数）。请你返回记录长度为 n 时，可能获得出勤奖励的记录情况 数量 。答案可能很大，所以返回对 10^9 + 7 取余 的结果。
//
//
//
//示例 1：
//
//输入：n = 2
//输出：8
//解释：
//有 8 种长度为 2 的记录将被视为可奖励：
//"PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
//只有"AA"不会被视为可奖励，因为缺勤次数为 2 次（需要少于 2 次）。
//示例 2：
//
//输入：n = 1
//输出：3
//示例 3：
//
//输入：n = 10101
//输出：183236316
//
//
//提示：
//
//1 <= n <= 10^5

//思路  动态规划
func checkRecord(n int) int {
	const (
		P = iota
		A
		L
		LL
		AL
		ALL
	)
	const mod = 1e9 + 7
	size := 6
	dp := make([]int, size)
	dp[P], dp[A], dp[L] = 1, 1, 1
	for i := 1; i < n; i++ {
		tmp := make([]int, size)
		for t, c := range dp {
			dp[t] = 0
			switch t {
			case P:
				tmp[P] = (tmp[P] + c) % mod
				tmp[A] = (tmp[A] + c) % mod
				tmp[L] = (tmp[L] + c) % mod
			case A:
				tmp[A] = (tmp[A] + c) % mod
				tmp[AL] = (tmp[AL] + c) % mod
			case L:
				tmp[P] = (tmp[P] + c) % mod
				tmp[A] = (tmp[A] + c) % mod
				tmp[LL] = (tmp[LL] + c) % mod
			case LL:
				tmp[P] = (tmp[P] + c) % mod
				tmp[A] = (tmp[A] + c) % mod
			case AL:
				tmp[A] = (tmp[A] + c) % mod
				tmp[ALL] = (tmp[ALL] + c) % mod
			case ALL:
				tmp[A] = (tmp[A] + c) % mod
			}
		}
		for index := range dp {
			dp[index] = tmp[index]
		}
	}
	var result int
	for _, c := range dp {
		result = (result + c) % mod
	}
	return result

}

func main() {
	//183236316
	println(checkRecord(10101))
}
