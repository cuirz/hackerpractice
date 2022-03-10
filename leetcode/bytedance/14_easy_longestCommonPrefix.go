package main

//14. 最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串""。
//
//示例1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母a-z。
func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size < 1 {
		return ""
	}
	if size == 1 {
		return strs[0]
	}
	result := strs[0]
	n := len(strs[0])
	for j := 1; j < size; j++ {
		n = min(n, len(strs[j]))
		if n < 1 {
			return ""
		}
		for i := 0; i < n; i++ {
			if result[i] != strs[j][i] {
				n = i
				break
			}
		}
	}

	return result[:n]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(longestCommonPrefix([]string{"aca", "cba"}))
}
