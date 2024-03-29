package main

//1190. 反转每对括号间的子串
//给出一个字符串s（仅含有小写英文字母和括号）。
//
//请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
//
//注意，您的结果中 不应 包含任何括号。
//
//
//
//示例 1：
//
//输入：s = "(abcd)"
//输出："dcba"
//示例 2：
//
//输入：s = "(u(love)i)"
//输出："iloveu"
//示例 3：
//
//输入：s = "(ed(et(oc))el)"
//输出："leetcode"
//示例 4：
//
//输入：s = "a(bcdefghijkl(mno)p)q"
//输出："apmnolkjihgfedcbq"
//
//
//提示：
//
//0 <= s.length <= 2000
//s 中只有小写英文字母和括号
//我们确保所有括号都是成对出现的

func reverseParentheses(s string) string {
	stack := make([][]byte, 0)
	var str []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, str)
			str = []byte{}
		} else if s[i] == ')' {
			for j, n := 0, len(str); j < n/2; j++ {
				str[j], str[n-j-1] = str[n-j-1], str[j]
			}
			str = append(stack[len(stack)-1], str...)
			stack = stack[:len(stack)-1]
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}

func main() {
	println(reverseParentheses("ta()usw((((a))))"))
}
