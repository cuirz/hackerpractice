package main

//208. 实现 Trie (前缀树)
//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
//示例:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//说明:
//
//你可以假设所有的输入都是由小写字母 a-z 构成的。
//保证所有输入均为非空字符串。

type Trie struct {
	W     [26]*Trie
	IsEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	next := this
	for _, v := range word {
		if next.W[v-'a'] == nil{
			next.W[v-'a'] = &Trie{}
		}
		next = next.W[v-'a']
	}
	next.IsEnd = true

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	next := this
	for _, v := range word {
		if next.W[v-'a'] == nil{
			return false
		}
		next = next.W[v-'a']
	}
	return next != nil && next.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	next := this
	for _, v := range prefix {
		if next.W[v-'a'] == nil{
			return false
		}
		next = next.W[v-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
["Trie","insert","insert","insert","insert","insert","insert","search","search","search","search","search","search","search","search","search","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith","startsWith"]
[[],["app"],["apple"],["beer"],["add"],["jam"],["rental"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"],["jam"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"],["jam"]]

[null,null,null,null,null,null,null,false,true,false,false,false,true,true,true,true,false,true,true,false,false,true,true,true,true]
[null,null,null,null,null,null,null,false,true,false,false,false,false,false,true,true,false,true,true,false,false,false,true,true,true]