package main

import (
	"bytes"
	"sort"
)

//1202. 交换字符串中的元素
//给你一个字符串s，以及该字符串中的一些「索引对」数组pairs，其中pairs[i] =[a, b]表示字符串中的两个索引（编号从 0 开始）。
//
//你可以 任意多次交换 在pairs中任意一对索引处的字符。
//
//返回在经过若干次交换后，s可以变成的按字典序最小的字符串。
//
//
//
//示例 1:
//
//输入：s = "dcab", pairs = [[0,3],[1,2]]
//输出："bacd"
//解释：
//交换 s[0] 和 s[3], s = "bcad"
//交换 s[1] 和 s[2], s = "bacd"
//示例 2：
//
//输入：s = "dcab", pairs = [[0,3],[1,2],[0,2]]
//输出："abcd"
//解释：
//交换 s[0] 和 s[3], s = "bcad"
//交换 s[0] 和 s[2], s = "acbd"
//交换 s[1] 和 s[2], s = "abcd"
//示例 3：
//
//输入：s = "cba", pairs = [[0,1],[1,2]]
//输出："abc"
//解释：
//交换 s[0] 和 s[1], s = "bca"
//交换 s[1] 和 s[2], s = "bac"
//交换 s[0] 和 s[1], s = "abc"
//提示：
//
//1 <= s.length <= 10^5
//0 <= pairs.length <= 10^5
//0 <= pairs[i][0], pairs[i][1] <s.length
//s中只含有小写英文字母

// 并查集 的问题
// 分集合
// //"udyyek"
//	//[[3, 3], [3, 0], [5,1], [3, 1], [3, 4], [3, 5]]
// 0,1,3,4,5 一个组    2 一个组
// 一个分组的 都可以排序  udyek  -> dekuy  ,  y - > y
// 回溯到自己位置    deykuy
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	array := []byte(s)
	group := make(map[int][]byte)
	var result bytes.Buffer

	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	var find func(index int) int
	var union func(v []int)
	union = func(v []int) {
		parents[find(v[0])] = find(v[1])
	}
	find = func(index int) int {
		for parents[index] != index {
			parents[index] = parents[parents[index]]
			index = parents[index]
		}
		return index
	}

	for _, v := range pairs {
		union(v)
	}
	for i := 0; i < n; i++ {
		index := find(parents[i])
		group[index] = append(group[index], array[i])
	}
	for _, v := range group {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	for i := 0; i < n; i++ {
		index := find(parents[i])
		if _, ok := group[index]; ok {
			result.WriteByte(group[index][0])
			group[index] = group[index][1:]
		}
	}
	return result.String()
}

func main() {
	//"udyyek"
	//[[3, 3], [3, 0], [5,1], [3, 1], [3, 4], [3, 5]]
	println(smallestStringWithSwaps("udyyek", [][]int{{3, 3}, {3, 0}, {5, 1}, {3, 1}, {3, 4}, {3, 5}}))
}
