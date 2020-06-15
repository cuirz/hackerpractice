package main

import (
	"sort"
	"strings"
)

//最小覆盖子串
//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
//
//示例：
//
//输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//说明：
//
//如果 S 中不存这样的子串，则返回空字符串 ""。
//如果 S 中存在这样的子串，我们保证它是唯一的答案。

//func minWindow(s string, t string) string {
//	l,r := 0,0
//
//}

func trailingZeroes(n int) int {
	p := perm(n)
	if p < 10 {
		return 0
	}
	count := 0
	for p%10 == 0 {
		p /= 10
		count++
	}
	return count

}

func perm(n int) int {
	if n <= 1 {
		return n
	}
	return n * perm(n-1)
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	var l, r byte
	s = strings.ToLower(s)
	for i < j {
		if !check(s[i]) {
			i++
			continue
		}
		l = s[i]
		if !check(s[j]) {
			j--
			continue
		}
		r = s[j]

		if l != r {
			return false
		}
		i++
		j--
	}
	return true

}

func check(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')

}

type svalue [][]int

func (p svalue) Len() int {
	return len(p)
}
func (p svalue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]

}
func (p svalue) Less(i, j int) bool {
	if p[i][1] < p[j][1] {
		return true
	}
	return false
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}
	var array svalue = intervals
	sort.Sort(array)
	max := 1
	count := 1
	for i, v := range array {
		count = 1
		for k := i + 1; k < len(intervals); k++ {
			if v[1] > array[k][0] {
				count++
			}
		}
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	//println(trailingZeroes(30))
	println(isPalindrome("race a car"))
	println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}
