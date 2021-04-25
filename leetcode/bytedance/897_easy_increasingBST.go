package main

//897. 递增顺序搜索树
//给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
//输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
//输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
//提示：
//
//树中节点数的取值范围是 [1, 100]
//0 <= Node.val <= 1000
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	result := &TreeNode{}
	rootNode := result
	var bst func(node *TreeNode)
	bst = func(node *TreeNode) {
		if node == nil {
			return
		}
		bst(node.Left)
		result.Right = node
		node.Left = nil
		result = node
		bst(node.Right)
	}
	bst(root)
	return rootNode.Right
}
