package main

//318. 最大单词长度乘积
//给定一个字符串数组words，找到length(word[i]) * length(word[j])的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
//
//
//
//示例1:
//
//输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
//输出: 16
//解释: 这两个单词为 "abcw", "xtfn"。
//示例 2:
//
//输入: ["a","ab","abc","d","cd","bcd","abcd"]
//输出: 4
//解释: 这两个单词为 "ab", "cd"。
//示例 3:
//
//输入: ["a","aa","aaa","aaaa"]
//输出: 0
//解释: 不存在这样的两个单词。
//
//
//提示：
//
//2 <= words.length <= 1000
//1 <= words[i].length <= 1000
//words[i]仅包含小写字母

func maxProduct(words []string) int {
	masks := make(map[int]int)
	for _, word := range words {
		mask := 0
		for _, v := range word {
			mask |= 1 << (v - 'a')
		}
		if len(word) >= masks[mask] {
			masks[mask] = len(word)
		}
	}

	var maxlen int
	for k1, v1 := range masks {
		for k2, v2 := range masks {
			if k1&k2 == 0 && v1*v2 > maxlen {
				maxlen = v1 * v2
			}
		}
	}
	return maxlen
}
