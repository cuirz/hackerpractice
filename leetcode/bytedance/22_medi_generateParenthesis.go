package main

//22. 括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例：
//
//输入：n = 3
//输出：[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]

//思路  递归
// 树状分支结构，去掉不成立的分支，保持可行的分支发展下去

func generateParenthesis(n int) []string {

	result := make([]string, 0)
	var generate func(s string, c1, c2 int)
	generate = func(s string, c1, c2 int) {
		if c1 == 0 && c2 == 0 {
			result = append(result, s)
			return
		}
		if c1 > 0 {
			generate(s+"(", c1-1, c2)
		}
		if c1 < c2 && c2 > 0 {
			generate(s+")", c1, c2-1)
		}
	}
	generate("(", n-1, n)
	return result
}

func main() {
	generateParenthesis(3)
}
