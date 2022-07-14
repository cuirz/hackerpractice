package main

import "strings"

//745. 前缀和后缀搜索
//设计一个包含一些单词的特殊词典，并能够通过前缀和后缀来检索单词。
//
//实现 WordFilter 类：
//
//WordFilter(string[] words) 使用词典中的单词 words 初始化对象。
//f(string pref, string suff) 返回词典中具有前缀prefix和后缀 suff的单词的下标。如果存在不止一个满足要求的下标，返回其中 最大的下标 。如果不存在这样的单词，返回 -1 。
//
//
//示例：
//
//输入
//["WordFilter", "f"]
//[[["apple"]], ["a", "e"]]
//输出
//[null, 0]
//解释
//WordFilter wordFilter = new WordFilter(["apple"]);
//wordFilter.f("a", "e"); // 返回 0 ，因为下标为 0 的单词：前缀 prefix = "a" 且 后缀 suff = "e" 。
//
//提示：
//
//1 <= words.length <= 10^4
//1 <= words[i].length <= 7
//1 <= pref.length, suff.length <= 7
//words[i]、pref 和 suff 仅由小写英文字母组成
//最多对函数 f 执行 10^4 次调用

type WordFilter struct {
	children map[string]*WordFilter
	weight   int
}

func Constructor(words []string) WordFilter {
	root := &WordFilter{children: make(map[string]*WordFilter)}

	builder := strings.Builder{}
	for i, word := range words {
		node := root
		m := len(word)
		for j := 0; j < m; j++ {
			tmp := node
			for k := j; k < m; k++ {
				builder.Reset()
				builder.WriteByte(word[k])
				builder.WriteByte('#')
				key := builder.String()
				if _, ok := tmp.children[key]; !ok {
					tmp.children[key] = &WordFilter{children: make(map[string]*WordFilter)}
				}
				tmp = tmp.children[key]
				tmp.weight = i
			}
			tmp = node
			for k := j; k < m; k++ {
				builder.Reset()
				builder.WriteByte('#')
				builder.WriteByte(word[m-k-1])
				key := builder.String()
				if _, ok := tmp.children[key]; !ok {
					tmp.children[key] = &WordFilter{children: make(map[string]*WordFilter)}
				}
				tmp = tmp.children[key]
				tmp.weight = i
			}
			builder.Reset()
			builder.WriteByte(word[j])
			builder.WriteByte(word[m-j-1])
			key := builder.String()
			if _, ok := node.children[key]; !ok {
				node.children[key] = &WordFilter{children: make(map[string]*WordFilter)}
			}
			node = node.children[key]
			node.weight = i
		}
	}
	return *root
}

func (this *WordFilter) F(pref string, suff string) int {
	builder := strings.Builder{}
	m := max(len(pref), len(suff))
	node := this
	for i := 0; i < m; i++ {
		builder.Reset()
		if i < len(pref) {
			builder.WriteByte(pref[i])
		} else {
			builder.WriteByte('#')
		}
		if i < len(suff) {
			builder.WriteByte(suff[len(suff)-i-1])
		} else {
			builder.WriteByte('#')
		}
		key := builder.String()
		if _, ok := node.children[key]; !ok {
			return -1
		}
		node = node.children[key]
	}
	return node.weight
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	obj := Constructor([]string{"apple", "aeeee"})
	param := obj.F("app", "pple")
	println(param)
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
