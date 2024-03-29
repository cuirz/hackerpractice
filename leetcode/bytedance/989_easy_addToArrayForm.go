package main

//989. 数组形式的整数加法
//对于非负整数X而言，X的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果X = 1231，那么其数组形式为[1,2,3,1]。
//
//给定非负整数 X 的数组形式A，返回整数X+K的数组形式。
//
//
//
//示例 1：
//
//输入：A = [1,2,0,0], K = 34
//输出：[1,2,3,4]
//解释：1200 + 34 = 1234
//示例 2：
//
//输入：A = [2,7,4], K = 181
//输出：[4,5,5]
//解释：274 + 181 = 455
//示例 3：
//
//输入：A = [2,1,5], K = 806
//输出：[1,0,2,1]
//解释：215 + 806 = 1021
//示例 4：
//
//输入：A = [9,9,9,9,9,9,9,9,9,9], K = 1
//输出：[1,0,0,0,0,0,0,0,0,0,0]
//解释：9999999999 + 1 = 10000000000
//
//
//提示：
//
//1 <= A.length <= 10000
//0 <= A[i] <= 9
//0 <= K <= 10000
//如果A.length > 1，那么A[0] != 0

func addToArrayForm(A []int, K int) []int {
	result := make([]int, 0)
	low, high := 0, 0
	for i := len(A) - 1; i > -1; i-- {
		low = high + A[i] + K%10
		high = low / 10
		low %= 10
		K /= 10
		result = append(result, low)
	}
	for K > 0 || high > 0 {
		low = high + K%10
		high = low / 10
		low %= 10
		K /= 10
		result = append(result, low)
	}
	n := len(result)
	for i := 0; i < n/2; i++ {
		result[i], result[n-1-i] = result[n-1-i], result[i]
	}
	return result
}
