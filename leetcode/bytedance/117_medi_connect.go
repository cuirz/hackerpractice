package main

//117. 填充每个节点的下一个右侧节点指针 II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	NEXT := root
	for NEXT != nil {
		next := NEXT
		var last,start *Node
		var help func(n *Node)
		help = func(n *Node) {
			if n == nil {
				return
			}
			if start == nil{
				start = n
			}
			if last != nil{
				last.Next = n
			}
			last = n
		}
		for next != nil{
			help(next.Left)
			help(next.Right)
			next = next.Next
		}
		NEXT = start
	}
	return root
}
