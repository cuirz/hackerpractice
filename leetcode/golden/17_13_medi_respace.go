package main

import "math"

//面试题 17.13. 恢复空格
//哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。
//
//注意：本题相对原题稍作改动，只需返回未识别的字符数
//
//
//示例：
//
//输入：
//dictionary = ["looked","just","like","her","brother"]
//sentence = "jesslookedjustliketimherbrother"
//输出： 7
//解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。
//提示：
//
//0 <= len(sentence) <= 1000
//dictionary中总字符数不超过 150000。
//你可以认为dictionary和sentence中只包含小写字母。

//思路 字典树 trie, 动态规划
// 制作单词后缀 字典树
// dp[i] 多少个未识别字数,默认赋值无穷大
// i作为下标i之前数组未识别出单词时，dp[i] = dp[i-1] + 1
// i之前下标j，从j....到i识别出单词时，dp[i]=min(dp[i],dp[j-1])

func respace(dictionary []string, sentence string) int {
	n := len(sentence)
	root := &Trie{}
	for _, word := range dictionary {
		root.put(word)
	}
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		dp[i+1] = dp[i] + 1
		next := root
		for j := i; j > -1; j-- {
			next = next.children[sentence[j]-'a']
			if next == nil {
				break
			}
			if next.isEnd {
				dp[i+1] = min(dp[i+1], dp[j])
			}
		}
	}
	return dp[n]
}

// 前缀树
type Trie struct {
	children [26]*Trie
	size     int
	isEnd    bool
}

func (t *Trie) put(word string) {
	if len(word) < 1 {
		return
	}
	root := t
	for i := len(word) - 1; i > -1; i-- {
		j := word[i] - 'a'
		if root.children[j] == nil {
			root.children[j] = &Trie{}
			root.size++
		}
		root = root.children[j]
	}
	root.isEnd = true

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(respace([]string{"looked", "just", "like", "her", "brother"}, "jesslookedjustliketimherbrother"))
}
