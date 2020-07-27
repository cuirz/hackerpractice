package main

import (
	"sort"
)

//406. 根据身高重建队列
//假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。
//
//注意：
//总人数少于1100人。
//
//示例
//
//输入:
//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
//输出:
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

//思路 贪婪算法
// 先排序，倒序，然后按顺次插入。
// 身高长的人无视矮的人，所以倒序插入按k序列就能完成
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i, v := range people {
		copy(people[v[1]+1:i+1], people[v[1]:i+1])
		people[v[1]] = v
	}
	return people
}
