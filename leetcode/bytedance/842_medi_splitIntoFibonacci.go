package main

import "math"

//842. 将数组拆分成斐波那契序列
//给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。
//
//形式上，斐波那契式序列是一个非负整数列表 F，且满足：
//
//0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
//F.length >= 3；
//对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。
//另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。
//
//返回从 S 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []。
//
//
//
//示例 1：
//
//输入："123456579"
//输出：[123,456,579]

//回溯，剪支

func splitIntoFibonacci(S string) []int {
	result := make([]int, 0)
	n := len(S)
	var backtrack func(index, sum, pre int) bool
	backtrack = func(index, sum, pre int) bool {
		if index == n {
			return len(result) >= 3
		}
		cur := 0

		for i := index; i < n; i++ {
			if i > index && S[index] == '0' {
				break
			}
			cur = cur*10 + int(S[i]-'0')
			if cur > math.MaxInt32 {
				break
			}
			if len(result) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}
			result = append(result, cur)
			if backtrack(i+1, pre+cur, cur) {
				return true
			}
			result = result[:len(result)-1]
		}

		return false
	}
	backtrack(0, 0, 0)
	return result
}
