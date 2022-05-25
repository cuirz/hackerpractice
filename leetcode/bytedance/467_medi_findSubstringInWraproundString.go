package main

//467. 环绕字符串中唯一的子字符串
//把字符串 s 看作是“abcdefghijklmnopqrstuvwxyz”的无限环绕字符串，所以s 看起来是这样的：
//
//"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".
//现在给定另一个字符串 p 。返回s 中唯一 的 p 的 非空子串的数量。
//
//
//
//示例1:
//
//输入: p = "a"
//输出: 1
//解释: 字符串 s 中只有一个"a"子字符。
//示例 2:
//
//输入: p = "cac"
//输出: 2
//解释: 字符串 s 中的字符串“cac”只有两个子串“a”、“c”。.
//示例 3:
//
//输入: p = "zab"
//输出: 6
//解释: 在字符串 s 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。
//
//
//提示:
//
//1 <= p.length <= 10^5
//p由小写英文字母构成

func findSubstringInWraproundString(p string) int {
	var dp [26]int
	j := 0
	for i, v := range p {
		if i > 0 && (p[i]-p[i-1]+26)%26 == 1 {
			j += 1
		} else {
			j = 1
		}
		dp[v-'a'] = max(dp[v-'a'], j)
	}
	var result int
	for _, v := range dp {
		result += v
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
	println(findSubstringInWraproundString("zabhijk"))
}
