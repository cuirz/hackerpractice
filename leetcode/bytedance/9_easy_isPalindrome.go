package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	l := 0
	r := len(s) - 1
	for l != r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true

}
