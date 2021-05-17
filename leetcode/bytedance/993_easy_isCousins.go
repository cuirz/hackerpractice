package main
//993. 二叉树的堂兄弟节点
//在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
//
//如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
//
//我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
//
//只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
//
//
//输入：root = [1,2,3,null,4], x = 2, y = 3
//输出：false
//提示：
//
//二叉树的节点数介于 2 到 100 之间。
//每个节点的值都是唯一的、范围为 1 到 100 的整数。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	var xFound,yFound bool
	var xParent,yParent *TreeNode
	var xDep,yDep int
	var search func(node,parent *TreeNode,dep int)
	search = func(node,parent *TreeNode,d int){
		if node == nil{
			return
		}
		if node.Val == x{
			xFound = true
			xParent = parent
			xDep = d
		}else if node.Val == y{
			yFound = true
			yParent = parent
			yDep = d
		}
		if xFound && yFound{
			return
		}
		search(node.Left,node,d+1)
		if xFound && yFound{
			return
		}
		search(node.Right,node,d+1)
	}
	search(root,nil,0)
	return xParent != yParent && xDep == yDep
}
