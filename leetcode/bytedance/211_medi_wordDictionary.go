package main

//211. 添加与搜索单词 - 数据结构设计
//请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
//
//实现词典类 WordDictionary ：
//
//WordDictionary() 初始化词典对象
//void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
//bool search(word) 如果数据结构中存在字符串与word 匹配，则返回 true ；否则，返回 false 。word 中可能包含一些 '.' ，每个. 都可以表示任何一个字母。
//
//
//示例：
//
//输入：
//["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//输出：
//[null,null,null,null,false,true,true,true]
//
//解释：
//WordDictionary wordDictionary = new WordDictionary();
//wordDictionary.addWord("bad");
//wordDictionary.addWord("dad");
//wordDictionary.addWord("mad");
//wordDictionary.search("pad"); // return False
//wordDictionary.search("bad"); // return True
//wordDictionary.search(".ad"); // return True
//wordDictionary.search("b.."); // return True
//
//
//提示：
//
//1 <= word.length <= 500
//addWord 中的 word 由小写英文字母组成
//search 中的 word 由 '.' 或小写英文字母组成
//最多调用 50000 次 addWord 和 search

type WordDictionary struct {
	IsEnd    bool
	Children [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) < 1 {
		return
	}
	root := this
	for _, v := range word {
		index := v - 'a'
		if root.Children[index] == nil {
			root.Children[index] = &WordDictionary{}
		}
		root = root.Children[index]
	}
	root.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	root := this
	for k, v := range word {
		if v == '.' {
			for i := 0; i < 26; i++ {
				if root.Children[i] != nil {
					if k == len(word)-1 {
						if root.Children[i].IsEnd {
							return true
						}
						continue
					}
					if root.Children[i].Search(word[k+1:]) {
						return true
					}
				}
			}
			return false
		} else {
			index := v - 'a'
			if root.Children[index] == nil {
				return false
			}
			root = root.Children[index]
		}
	}
	return root.IsEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
