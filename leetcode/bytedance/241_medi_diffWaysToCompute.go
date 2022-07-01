package main

import "unicode"

//241. 为运算表达式设计优先级
//给你一个由数字和运算符组成的字符串expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
//
//生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 10^4 。
//
//
//
//示例 1：
//
//输入：expression = "2-1-1"
//输出：[0,2]
//解释：
//((2-1)-1) = 0
//(2-(1-1)) = 2
//示例 2：
//
//输入：expression = "2*3-4*5"
//输出：[-34,-14,-10,-10,10]
//解释：
//(2*(3-(4*5))) = -34
//((2*3)-(4*5)) = -14
//((2*(3-4))*5) = -10
//(2*((3-4)*5)) = -10
//(((2*3)-4)*5) = 10
//
//
//提示：
//
//1 <= expression.length <= 20
//expression 由数字和算符 '+'、'-' 和 '*' 组成。
//输入表达式中的所有整数值在范围 [0, 99]

// 预处理  顶层自上而下，当选择最后一个算符时，集合left和集合right 就是最终可能结果的组合并集
func diffWaysToCompute(expression string) []int {
	const OpAdd, OpSub, OpMul = -1, -2, -3
	var ops []int
	for i := 0; i < len(expression); {
		if unicode.IsDigit(rune(expression[i])) {
			x := 0
			for ; i < len(expression) && unicode.IsDigit(rune(expression[i])); i++ {
				x = x*10 + int(expression[i]-'0')
			}
			ops = append(ops, x)
		} else {
			if expression[i] == '+' {
				ops = append(ops, OpAdd)
			} else if expression[i] == '-' {
				ops = append(ops, OpSub)
			} else {
				ops = append(ops, OpMul)
			}
			i++
		}
	}
	n := len(ops)
	var dp = make([][][]int, n)
	for i, _ := range dp {
		dp[i] = make([][]int, n)
	}
	var dfs func(int, int) []int
	dfs = func(l, r int) []int {
		res := dp[l][r]
		if res != nil {
			return res
		}
		if l == r {
			dp[l][r] = []int{ops[l]}
			return dp[l][r]
		}
		for i := l; i < r; i += 2 {
			left := dfs(l, i)
			right := dfs(i+2, r)
			for _, x := range left {
				for _, y := range right {
					if ops[i+1] == OpAdd {
						dp[l][r] = append(dp[l][r], x+y)
					} else if ops[i+1] == OpSub {
						dp[l][r] = append(dp[l][r], x-y)
					} else {
						dp[l][r] = append(dp[l][r], x*y)
					}
				}
			}
		}
		return dp[l][r]
	}
	return dfs(0, n-1)
}
