package main

//139. 单词拆分
//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//    注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

func wordBreak(s string, wordDict []string) bool {
	//"cars"
	//["car","ca","rs"]
	dic := make(map[string]int)
	for _, v := range wordDict {
		dic[v] = 1
	}
	memo := make(map[string]bool)
	var check func(str string) bool
	check = func(str string) bool {
		n := len(str)
		if n < 1 {
			return true
		}
		if _, ok := memo[str]; ok {
			return memo[str]
		}
		for i := 1; i <= n; i++ {
			if dic[str[:i]] > 0 && check(str[i:]) {
				memo[str] = true
				return true
			}
		}
		memo[str] = false
		return false
	}
	return check(s)
}

func main() {
	println(wordBreak("cars", []string{"car", "ca", "rs"}))
}
