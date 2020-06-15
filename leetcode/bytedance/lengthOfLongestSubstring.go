package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	max := 0
	m := 0
	index := 0
	for i := 0; i < len(s); i++ {
		c := strings.Index(s[index:i], string(s[i]))
		if c > -1 {
			index += c + 1
		}
		m = i - index + 1
		if max < m {
			max = m
		}
	}
	return max
}

func main() {
	print(lengthOfLongestSubstring("bbbbb"))
}
