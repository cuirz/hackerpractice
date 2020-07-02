package main

import "strings"

//面试题 16.18. 模式匹配
//你有两个字符串，即pattern和value。 pattern字符串由字母"a"和"b"组成，用于描述字符串中的模式。例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），该字符串也匹配像"a"、"ab"和"b"这样的模式。但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。
//
//示例 1：
//
//输入： pattern = "abba", value = "dogcatcatdog"
//输出： true
//示例 2：
//
//输入： pattern = "abba", value = "dogcatcatfish"
//输出： false
//示例 3：
//
//输入： pattern = "aaaa", value = "dogcatcatdog"
//输出： false
//示例 4：
//
//输入： pattern = "abba", value = "dogdogdogdog"
//输出： true
//解释： "a"="dogdog",b=""，反之也符合规则
//提示：
//
//0 <= len(pattern) <= 1000
//0 <= len(value) <= 1000
//你可以假设pattern只包含字母"a"和"b"，value仅包含小写字母。

func patternMatching(pattern string, value string) bool {
	n := len(value)
	an := strings.Count(pattern, "a")
	bn := strings.Count(pattern, "b")

	var a, b, tmp string
	matched := true
	x, y := 0, 0

	if an == 0 {
		an, bn = bn, an
		pattern = strings.ReplaceAll(pattern, "b", "a")
	}
	if len(value) == 0 {
		return bn == 0
	}
	if len(pattern) == 0 {
		return false
	}

	for x = n / an; x > -1; x-- {
		if bn > 0 {
			y = (n - an*x) / bn
		} else {
			y = 0
		}
		s := value
		a, b = "", ""
		matched = true
		for _, v := range pattern {
			switch v {
			case 'a':
				if x == 0 {
					continue
				}
				if x > len(s) {
					matched = false
					break
				}
				tmp = s[:x]
				if len(a) == 0 {
					a = tmp
				} else if a != tmp {
					matched = false
					break
				}
				s = s[x:]
			case 'b':
				if y == 0 {
					continue
				}
				if y > len(s) {
					matched = false
					break
				}
				tmp = s[:y]
				if len(b) == 0 {
					b = tmp
				} else if b != tmp {
					matched = false
					break
				}
				s = s[y:]
			}
		}
		if matched && len(s) == 0 {
			return true
		}

	}
	return false
}

func main() {
	println(patternMatching("bbba", "xxxxxx"))
}
