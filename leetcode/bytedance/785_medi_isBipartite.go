package main

//785. 判断二分图
//给定一个无向图graph，当这个图为二分图时返回true。
//
//如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。
//
//graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。每个节点都是一个在0到graph.length-1之间的整数。这图中没有自环和平行边：graph[i]中不存在i，并且graph[i]中没有重复的值。
//
//
//示例 1:
//输入: [[1,3], [0,2], [1,3], [0,2]]
//输出: true
//解释:
//无向图如下:
//0----1
//|    |
//|    |
//3----2
//我们可以将节点分成两组: {0, 2} 和 {1, 3}。
//
//示例 2:
//输入: [[1,2,3], [0,2], [0,1,3], [0,2]]
//输出: false
//解释:
//无向图如下:
//0----1
//| \  |
//|  \ |
//3----2
//我们不能将节点分割成两个独立的子集。
//注意:
//
//graph 的长度范围为 [1, 100]。
//graph[i] 中的元素的范围为 [0, graph.length - 1]。
//graph[i] 不会包含 i 或者有重复的值。
//图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。

//二部图又叫二分图，是图论中的一种特殊模型。设G=(V,E)是一个无向图，如果顶点V可分割为两个互不相交的子集(A,B)，并且图中的每条边（i，j）所关联的两个
//顶点i和j分别属于这两个不同的顶点集(i in A,j in B)，则称图G为一个二分图。简单来说，如果图中点可以被分为两组，并且使得所有边都跨越组的边界，
//则这就是一个二分图。准确地说：把一个图的顶点划分为两个不相交子集，使得每一条边都分别连接两个集合中的顶点。如果存在这样的划分，则此图为一个二分图。

//思路 深度优先或者广度优先

func isBipartite(graph [][]int) bool {
	const RED, BLUE = 1, 2
	n := len(graph)
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		if nodes[i] == 0 {
			nodes[i] = RED
			queue := make([]int, 0)
			queue = append(queue, i)
			for j := 0; j < len(queue); j++ {
				against := RED
				if nodes[queue[j]] == RED {
					against = BLUE
				}
				for _, v := range graph[queue[j]] {
					if nodes[v] == 0 {
						nodes[v] = against
						queue = append(queue, v)
					} else if nodes[v] != against {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {
	println(isBipartite([][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}))
}
