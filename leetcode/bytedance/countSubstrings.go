package main

//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。
//
//示例 1:
//
//输入: "abc"
//输出: 3
//解释: 三个回文子串: "a", "b", "c".
//示例 2:
//
//输入: "aaa"
//输出: 6
//说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".

func countSubstrings(s string) int {
	count := 0
	find := func(l, r int) {
		for l > -1 && r < len(s){
			if s[l] != s[r]{
				break
			}
			l --
			r ++
			count ++
		}
	}
	for i := 0; i < len(s); i++ {
		find(i,i)
		find(i,i+1)
	}
	return count
}
