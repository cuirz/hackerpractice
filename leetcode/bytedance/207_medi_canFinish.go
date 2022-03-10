package main

//207. 课程表
//你这个学期必须选修 numCourse 门课程，记为0到numCourse-1 。
//
//在选修某些课程之前需要一些先修课程。例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
//
//给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

// 有向无环图变成线性的排序 ---叫拓扑排序

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	adjacency := make(map[int][]int)
	var queue []int
	//通过边缘列表,创建邻接列表，并计算节点的入度值做成表
	for _, v := range prerequisites {
		cur := v[0]
		pre := v[1]
		indegree[cur]++
		adjacency[pre] = append(adjacency[pre], cur)
	}

	// 入度是0的课程放入到队列中，入度0代表没有前置课程
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		// 寻找把入度0的课程当做前置课的节点
		pre := queue[0]
		queue = queue[1:]
		numCourses--
		for _, v := range adjacency[pre] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return numCourses == 0

}

func main() {
	canFinish(1, [][]int{{1, 0}, {0, 1}})
}
