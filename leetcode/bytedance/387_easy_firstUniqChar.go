package main

//387. 字符串中的第一个唯一字符
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
// 
//
//示例：
//
//s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2

func firstUniqChar(s string) int {
	dic := make(map[byte]int)
	for _, v := range s {
		dic[byte(v)] ++
	}
	for i, v := range s {
		if dic[byte(v)] < 2 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	dic := make([]int, 26)
	for i, v := range s {
		index := v - 'a'
		if dic[index] == 0 {
			dic[index] = -(i + 1)
		} else if dic[index] < 0 {
			dic[index] = 2
		}
	}
	result := -1
	for _, v := range dic {
		if v < 0 {
			if result == -1 {
				result = -v - 1
			} else {
				result = min(result, -v-1)
			}
		}
	}
	return result
}
