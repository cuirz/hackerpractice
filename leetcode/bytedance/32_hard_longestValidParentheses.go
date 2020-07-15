package main

//32. 最长有效括号
//
//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
//示例 1:
//
//输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//示例 2:
//
//输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"

//思路  栈 ，动态规划
//"()(()"

func longestValidParentheses(s string) int {
	result := 0
	n := len(s)
	array := make([]int, n)
	stack := make([]int, 0)
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			size := len(stack)
			if size > 0 {
				array[i] = 1
				array[stack[size-1]] = 1
				stack = stack[:size-1]
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		if array[i] == 1 {
			sum += 1
			result = max(result, sum)
		} else {
			sum = 0
		}

	}
	return result
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	println(longestValidParentheses(")()())"))
}
