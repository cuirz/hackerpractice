package main

//210. 课程表 II
//现在你总共有 n 门课需要选，记为 0 到 n-1。
//
//在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
//
//给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
//
//可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
//说明:
//
//输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
//你可以假定输入的先决条件中没有重复的边。
//提示:
//
//这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
//通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
//拓扑排序也可以通过 BFS 完成。

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, numCourses)
	adjacency := make(map[int][]int)
	queue := make([]int, 0)

	for _, v := range prerequisites {
		cur := v[0]
		pre := v[1]
		indegree[cur]++
		adjacency[pre] = append(adjacency[pre], cur)
	}

	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	result := make([]int, 0)
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		result = append(result, pre)
		numCourses--
		for _, v := range adjacency[pre] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if numCourses == 0 {
		return result
	}
	return []int{}
}
