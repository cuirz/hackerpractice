package main

import "math"

//126. 单词接龙 II
//给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：
//
//每次转换只能改变一个字母。
//转换后得到的单词必须是字典中的单词。
//说明:
//
//如果不存在这样的转换序列，返回一个空列表。
//所有单词具有相同的长度。
//所有单词只由小写字母组成。
//字典中不存在重复的单词。
//你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//示例 1:
//
//输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出:
//[
//["hit","hot","dot","dog","cog"],
// ["hit","hot","lot","log","cog"]
//]
//示例 2:
//
//输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释:endWord "cog" 不在字典中，所以不存在符合要求的转换序列。

//思路 广度优先 ， 图

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	dic := make(map[string]int)
	for i := 0; i < len(wordList); i++ {
		dic[wordList[i]] = i
	}
	if _, ok := dic[endWord]; !ok {
		return [][]string{}
	}
	if _, ok := dic[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		dic[beginWord] = len(wordList) - 1
	}

	cost := make([]int, len(wordList))
	for i := 0; i < len(wordList); i++ {
		cost[i] = math.MaxInt32
	}
	check := func(a, b string) bool {
		count := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count == 1
	}

	edge := make(map[int][]int)
	for i := 0; i < len(dic); i++ {
		for j := i + 1; j < len(dic); j++ {
			if check(wordList[i], wordList[j]) {
				edge[i] = append(edge[i], j)
				edge[j] = append(edge[j], i)
			}
		}
	}

	result := make([][]string, 0)
	dest := dic[endWord]
	path := make([][]int, 0)
	// 加入 begin id
	path = append(path, []int{dic[beginWord]})
	cost[dic[beginWord]] = 0
	for len(path) > 0 {
		node := path[0]
		last := node[len(node)-1]
		path = path[1:]
		if last == dest {
			p := make([]string, 0)
			for _, v := range node {
				p = append(p, wordList[v])
			}
			result = append(result, p)
		} else {
			for _, to := range edge[last] {
				if cost[to] >= cost[last]+1 {
					cost[to] = cost[last] + 1
					p := make([]int, len(node))
					copy(p, node)
					p = append(p, to)
					path = append(path, p)
				}
			}
		}
	}
	return result
}

func main() {
	findLadders("a", "c", []string{"a", "b", "c"})
}
