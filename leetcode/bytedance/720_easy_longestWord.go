package main

import "sort"

//720. 词典中最长的单词
//给出一个字符串数组words 组成的一本英语词典。返回words 中最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。
//
//若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。
//
//
//
//示例 1：
//
//输入：words = ["w","wo","wor","worl", "world"]
//输出："world"
//解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。
//示例 2：
//
//输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
//输出："apple"
//解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"
//
//
//提示：
//
//1 <= words.length <= 1000
//1 <= words[i].length <= 30
//所有输入的字符串words[i]都只包含小写字母。

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] > words[j]
		}
		return len(words[i]) < len(words[j])
	})
	var result string
	dic := map[string]struct{}{"": {}}
	for _, w := range words {
		if _, ok := dic[w[:len(w)-1]]; ok {
			result = w
			dic[w] = struct{}{}
		}
	}
	return result

}

func main() {
	println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
}
