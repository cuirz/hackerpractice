package main

//116. 填充每个节点的下一个右侧节点指针
//给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
//struct Node {
//int val;
//Node *left;
//Node *right;
//Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
//初始状态下，所有next 指针都被设置为 NULL。

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	deep := root
	for deep != nil {
		var start, last *Node
		next := deep
		help := func(n *Node) {
			if n == nil {
				return
			}
			if start == nil {
				start = n
			}
			if last != nil {
				last.Next = n
			}
			last = n
		}
		for next != nil {
			help(next.Left)
			help(next.Right)
			next = next.Next
		}
		deep = start
	}
	return root

}
