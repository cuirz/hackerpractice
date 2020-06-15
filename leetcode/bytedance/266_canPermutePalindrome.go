package main

func canPermutePalindrome(s string) bool {
	dic := make(map[rune]int, 0)
	for _, v := range s {
		dic[v]++
		if dic[v] > 1 {
			delete(dic, v)
		}
	}
	return len(dic) < 2
}
