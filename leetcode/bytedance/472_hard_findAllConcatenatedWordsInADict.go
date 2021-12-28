package main

import "sort"

//472. 连接词
//给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。
//
//连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。
//
//
//
//示例 1：
//
//输入：words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
//输出：["catsdogcats","dogcatsdog","ratcatdogcat"]
//解释："catsdogcats" 由 "cats", "dog" 和 "cats" 组成;
//"dogcatsdog" 由 "dog", "cats" 和 "dog" 组成;
//"ratcatdogcat" 由 "rat", "cat", "dog" 和 "cat" 组成。
//示例 2：
//
//输入：words = ["cat","dog","catdog"]
//输出：["catdog"]
//
//
//提示：
//
//1 <= words.length <= 10^4
//0 <= words[i].length <= 1000
//words[i] 仅由小写字母组成
//0 <= sum(words[i].length) <= 10^5

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	root := &trie{}
	var ans []string
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		if root.dfs(w) {
			ans = append(ans, w)
		} else {
			root.insert(w)
		}
	}
	return ans
}

type trie struct {
	c     [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	node := t
	for _, w := range word {
		w -= 'a'
		if node.c[w] == nil {
			node.c[w] = &trie{}
		}
		node = node.c[w]
	}
	node.isEnd = true
}
func (t *trie) dfs(word string) bool {
	if len(word) == 0 {
		return true
	}
	node := t
	for i, w := range word {
		node = node.c[w-'a']
		if node == nil {
			return false
		}
		if node.isEnd && t.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}
