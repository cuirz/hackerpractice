package main

import "bytes"

func longestPalindrome(s string) string {
	size := len(s)
	if size < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < size; i++ {
		l1 := expand(s, i, i)
		l2 := expand(s, i, i+1)
		m := max(l1, l2)
		if m > end-start {
			start = i - (m-1)/2
			end = i + m/2
		}
	}
	return s[start : end+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}

func main() {
	buff := bytes.Buffer{}
	buff.WriteString(s)
	buff.Bytes()

}
