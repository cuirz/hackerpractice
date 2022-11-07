package main

//816. 模糊坐标
//我们有一些二维坐标，如 "(1, 3)" 或 "(2, 0.5)"，然后我们移除所有逗号，小数点和空格，得到一个字符串S。返回所有可能的原始字符串到一个列表中。
//
//原始的坐标表示法不会存在多余的零，所以不会出现类似于"00", "0.0", "0.00", "1.0", "001", "00.01"或一些其他更小的数来表示坐标。此外，一个小数点前至少存在一个数，所以也不会出现“.1”形式的数字。
//
//最后返回的列表可以是任意顺序的。而且注意返回的两个数字中间（逗号之后）都有一个空格。
//
//
//
//示例 1:
//输入: "(123)"
//输出: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
//示例 2:
//输入: "(00011)"
//输出:  ["(0.001, 1)", "(0, 0.011)"]
//解释:
//0.0, 00, 0001 或 00.01 是不被允许的。
//示例 3:
//输入: "(0123)"
//输出: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]
//示例 4:
//输入: "(100)"
//输出: [(10, 0)]
//解释:
//1.0 是不被允许的。
//
//
//提示:
//
//4 <= S.length <= 12.
//S[0] = "(", S[S.length - 1] = ")", 且字符串 S 中的其他元素都是数字。
func ambiguousCoordinates(s string) []string {
	var makeNum func(string) []string
	makeNum = func(str string) (res []string) {
		if str[0] != '0' || str == "0" {
			res = append(res, str)
		}
		for i := 1; i < len(str); i++ {
			if i != 1 && str[0] == '0' || str[len(str)-1] == '0' {
				break
			}
			res = append(res, str[i:]+".", str[i:])
		}
		return
	}
	n := len(s) - 2
	s = s[1 : len(s)-1]
	var result []string

	for i := 1; i < n; i++ {
		left := makeNum(s[:i])
		if len(left) == 0 {
			continue
		}
		right := makeNum(s[i:])
		if len(right) == 0 {
			continue
		}
		for _, x := range left {
			for _, y := range right {
				result = append(result, "("+x+", "+y+")")
			}
		}
	}
	return result
}
