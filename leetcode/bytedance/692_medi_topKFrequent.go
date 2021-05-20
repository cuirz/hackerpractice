package main

import "container/heap"

//692. 前K个高频单词
//给一非空的单词列表，返回前 k 个出现次数最多的单词。
//
//返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
//
//示例 1：
//
//输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//注意，按字母顺序 "i" 在 "love" 之前。
// 
//
//示例 2：
//
//输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
//输出: ["the", "is", "sunny", "day"]
//解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//出现次数依次为 4, 3, 2 和 1 次。
// 
//
//注意：
//
//假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
//输入的单词均由小写字母组成。
// 
//
//扩展练习：
//
//尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决

//思路 哈希表+优先队列
func topKFrequent(words []string, k int) []string {
	dic := make(map[string]int)
	for _, v := range words {
		dic[v] ++
	}
	result := make([]string, k)
	stack := &hp{}
	for key, val := range dic {
		heap.Push(stack,pair{key,val})
		if stack.Len() > k{
			heap.Pop(stack)
		}
	}
	for i:=k-1;i>-1;i--{
		result[i] = heap.Pop(stack).(pair).word
	}
	return result
}

type pair struct {
	word string
	size int
}
type hp []pair
func (h hp) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *hp)Push(v interface{}){
	*h = append(*h,v.(pair))
}
func (h *hp)Pop()interface{}{
	res := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return res
}
func (h hp)Less(i,j int)bool{
	return h[i].size < h[j].size || h[i].size == h[j].size && h[i].word > h[j].word
}
func (h hp)Len()int{
	return len(h)
}
