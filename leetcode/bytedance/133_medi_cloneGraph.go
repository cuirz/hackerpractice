package main

//133. 克隆图
//给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
//
//图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
// * Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
	dic := make(map[*Node]*Node)
	var clone func(root *Node)*Node
	clone = func(root *Node)*Node{
		if dic[root] != nil{
			return dic[root]
		}
		head := &Node{
			Val:       root.Val,
			Neighbors: make([]*Node, 0),
		}
		dic[root] = head
		for _, v := range root.Neighbors {
			if dic[v] == nil {
				dic[v] = clone(v)
			}
			head.Neighbors = append(head.Neighbors, dic[v])
		}
		return head
	}

	return clone(node)

}
