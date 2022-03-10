package main

//797. 所有可能的路径
//给你一个有n个节点的 有向无环图（DAG），请你找出所有从节点 0到节点 n-1的路径并输出（不要求按特定顺序）
//
//二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。
//
//译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。
//
//
//
//示例 1：
//
//
//
//输入：graph = [[1,2],[3],[3],[]]
//输出：[[0,1,3],[0,2,3]]
//解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
//示例 2：
//
//
//
//输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
//输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
//示例 3：
//
//输入：graph = [[1],[]]
//输出：[[0,1]]
//示例 4：
//
//输入：graph = [[1,2,3],[2],[3],[]]
//输出：[[0,1,2,3],[0,2,3],[0,3]]
//示例 5：
//
//输入：graph = [[1,3],[2],[3],[]]
//输出：[[0,1,2,3],[0,3]]
//
//
//提示：
//
//n == graph.length
//2 <= n <= 15
//0 <= graph[i][j] < n
//graph[i][j] != i（即，不存在自环）
//graph[i] 中的所有元素 互不相同
//保证输入为 有向无环图（DAG

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	result := make([][]int, 0)
	visited := []int{0}
	var search func(index int)
	search = func(index int) {
		if index == n-1 {
			clone := make([]int, len(visited))
			copy(clone, visited)
			result = append(result, clone)
			return
		}
		for _, v := range graph[index] {
			visited = append(visited, v)
			search(v)
			visited = visited[:len(visited)-1]
		}
	}
	search(0)
	return result

}

func main() {
	println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
