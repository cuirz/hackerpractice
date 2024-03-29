package main

//28. 实现 strStr()
//实现strStr()函数。
//
//给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//
//
//
//说明：
//
//当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
//对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。
//
//
//
//示例 1：
//
//输入：haystack = "hello", needle = "ll"
//输出：2
//示例 2：
//
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1
//示例 3：
//
//输入：haystack = "", needle = ""
//输出：0
//
//
//提示：
//
//0 <= haystack.length, needle.length <= 5 * 10^4
//haystack 和 needle 仅由小写英文字符组成

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if n < m {
		return -1
	} else if m == 0 {
		return 0
	}
	for i := 0; i < n-m+1; i++ {
		if haystack[i] == needle[0] {
			index := i + 1
			find := true
			for j := 1; j < m; j++ {
				if haystack[index] != needle[j] {
					find = false
					break
				}
				index++
			}
			if find {
				return i
			}
		}
	}
	return -1
}

func main() {
	//println(strStr("abmissimississippi",
	//	"mississippi"))
	println(strStr("ababbbaabbba",
		"abb"))
}
