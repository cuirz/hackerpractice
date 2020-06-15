package main

import (
	"bytes"
	"sort"
)

// 单词搜索II
// 使用 前缀树  或者 散列表
func findWords(board [][]byte, words []string) []string {
	if len(board) < 1 {
		return []string{}
	}
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	result := make([]string, 0)
	dic := make(map[string]bool)
	wordmap := make(map[string]bool)
	// 先做前缀字典表
	for _, word := range words {
		wordmap[word] = true
		for i, _ := range word {
			dic[word[:i+1]] = true
		}
	}

	var search func(i, j int, subword string)
	search = func(i, j int, subword string) {
		if visited[i][j] || len(wordmap) < 1 {
			return
		}

		var buffer bytes.Buffer
		buffer.WriteString(subword)
		buffer.WriteByte(board[i][j])
		key := buffer.String()
		if _, ok := dic[key]; !ok {
			return
		}
		if _, ok := wordmap[key]; ok {
			delete(wordmap, key)
			result = append(result, key)
		}
		visited[i][j] = true
		if i > 0 {
			search(i-1, j, key)
		}
		if j > 0 {
			search(i, j-1, key)
		}
		if i < rows-1 {
			search(i+1, j, key)
		}
		if j < cols-1 {
			search(i, j+1, key)
		}
		visited[i][j] = false

	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			search(i, j, "")
		}
	}
	sort.Strings(result)
	return result
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
	for _, v := range word {
		i := v - 'a'
		if root.children[i] == nil {
			root.children[i] = &Trie{}
			root.size++
		}
		root = root.children[i]
	}
	root.isEnd = true
}

func findWordsByTrie(board [][]byte, words []string) []string {
	if len(board) < 1 {
		return []string{}
	}
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	result := make([]string, 0)
	trie := &Trie{}
	// 先做前缀树
	for _, word := range words {
		trie.put(word)
	}

	var search func(i, j int, root *Trie, subword bytes.Buffer)
	search = func(i, j int, root *Trie, subword bytes.Buffer) {
		index := board[i][j] - 'a'
		if visited[i][j] || root.children[index] == nil {
			return
		}
		subword.WriteByte(board[i][j])

		if root.children[index].isEnd {
			root.children[index].isEnd = false
			result = append(result, subword.String())
			if root.children[index].size == 0 {
				root.children[index] = nil
				root.size--
				return
			}
		}

		visited[i][j] = true
		if i > 0 {
			search(i-1, j, root.children[index], subword)
		}
		if j > 0 {
			search(i, j-1, root.children[index], subword)
		}
		if i < rows-1 {
			search(i+1, j, root.children[index], subword)
		}
		if j < cols-1 {
			search(i, j+1, root.children[index], subword)
		}
		visited[i][j] = false

	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			search(i, j, trie, bytes.Buffer{})
		}
	}
	sort.Strings(result)
	return result
}

func main() {

	println(findWordsByTrie([][]byte{{'a', 'b'}, {'a', 'a'}}, []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}))

}
