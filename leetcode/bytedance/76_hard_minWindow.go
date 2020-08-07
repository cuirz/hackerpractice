package main

//76. 最小覆盖子串
//给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。
//
// 
//
//示例：
//
//输入：S = "ADOBECODEBANC", T = "ABC"
//输出："BANC"
// 
//
//提示：
//
//如果 S 中不存这样的子串，则返回空字符串 ""。
//如果 S 中存在这样的子串，我们保证它是唯一的答案。

//思路 字符统计,滑动窗口
//
func minWindow(s string, t string) string {
	n := len(s)

	dic := make(map[byte]int)
	win := make(map[byte]int)
	for _, v := range t {
		dic[byte(v)] ++
	}
	isEqual := func() bool {
		for k, v := range dic {
			if win[k] < v {
				return false
			}
		}
		return true
	}

	size, l, r := n, 0, 0
	var result string

	for r < n {
		win[s[r]]++
		if dic[s[r]] > 0 {
			for isEqual() {
				if size >= r-l+1 {
					size = r - l + 1
					result = s[l:l+size]
				}
				for {
					win[s[l]] --
					l++
					if l == n || dic[s[l]] != 0 {
						break
					}
				}
			}
		}
		r++
	}
	return result

}

func main(){
	println(minWindow("a","aa"))
}