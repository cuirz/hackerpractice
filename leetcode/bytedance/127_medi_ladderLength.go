package main

import "math"

//27. 单词接龙
//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典中的单词。
//说明:
//
//如果不存在这样的转换序列，返回 0。
//所有单词具有相同的长度。
//所有单词只由小写字母组成。
//字典中不存在重复的单词。
//你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//示例 1:
//
//输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//返回它的长度 5。

//思路  图，广度优先

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dic := make(map[string]int)
	graph := make([][]int, 0)
	addWord := func(word string) int {
		if v, ok := dic[word]; !ok {
			id := len(dic)
			dic[word] = id
			graph = append(graph, []int{})
			return id
		} else {
			return v
		}
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		tmp := []byte(word)
		for i, v := range tmp {
			tmp[i] = '*'
			id2 := addWord(string(tmp))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			tmp[i] = v
		}
		return id1
	}
	for _, w := range wordList {
		addEdge(w)
	}
	start := addEdge(beginWord)
	end, ok := dic[endWord]
	if !ok {
		return 0
	}

	dist := make([]int, len(dic))
	for i:=0;i<len(dic);i++{
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	queue := []int{start}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == end {
			return dist[end]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == math.MaxInt32 {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0

}
