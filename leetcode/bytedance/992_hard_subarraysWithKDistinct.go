package main

//992. K 个不同整数的子数组
//给定一个正整数数组 A，如果 A的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。
//
//（例如，[1,2,3,1,2] 中有3个不同的整数：1，2，以及3。）
//
//返回A中好子数组的数目。
//
//
//
//示例 1：
//
//输入：A = [1,2,1,2,3], K = 2
//输出：7
//解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
//示例 2：
//
//输入：A = [1,2,1,3,4], K = 3
//输出：3
//解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].
//
//
//提示：
//
//1 <= A.length <= 20000
//1 <= A[i] <= A.length
//1 <= K <= A.length

//思路 滑动窗口
//K个长度子集 = 最多存在 K 个子集 - 最多存在 (K-1)个子集
func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	mostK := func(arr []int, k int) int {
		result := 0
		dic := make([]int, n+1)
		size := 0
		l, r := 0, 0
		for r < n {
			if dic[arr[r]] == 0 {
				size++
			}
			dic[arr[r]]++
			r++
			for size > k {
				dic[arr[l]]--
				if dic[arr[l]] == 0 {
					size--
				}
				l++
			}
			result += r - l
		}
		return result
	}

	return mostK(A, K) - mostK(A, K-1)
}

func main() {
	print(subarraysWithKDistinct([]int{1, 1, 1, 1, 1}, 1))
}
