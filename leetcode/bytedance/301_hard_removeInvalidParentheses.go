package main

import (
	"strings"
)

//301. 删除无效的括号
//删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
//
//说明: 输入可能包含了除(和)以外的字符。
//
//示例 1:
//
//输入: "()())()"
//输出: ["()()()", "(())()"]
//示例 2:
//
//输入: "(a)())()"
//输出: ["(a)()()", "(a())()"]
//示例 3:
//
//输入: ")("
//输出: [""]

//思路 广度优先 堆
func removeInvalidParentheses(s string) []string {
	isValid := func(str string) bool {
		count := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '(' {
				count++
			} else if str[i] == ')' {
				count--
			}
			if count < 0 {
				return false
			}
		}
		return count == 0
	}
	result := make([]string, 0)
	set := make(map[string]int)
	bd := strings.Builder{}
	queue := make([]string, 0)
	queue = append(queue, s)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if isValid(top) {
			result = append(result, top)
		} else if len(result) == 0 {
			for i := 0; i < len(top); i++ {
				if top[i] == '(' || top[i] == ')' {
					tmp := ""
					if i < len(top)-1 {
						bd.Reset()
						bd.WriteString(top[:i])
						bd.WriteString(top[i+1:])
						tmp = bd.String()
					} else {
						tmp = top[:i]
					}
					if set[tmp] == 0 {
						set[tmp] = 1
						queue = append(queue, tmp)
					}
				}
			}
		}
	}
	return result
}

func main() {
	println(removeInvalidParentheses("((()"))

}
