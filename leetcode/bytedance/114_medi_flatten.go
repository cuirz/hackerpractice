package main

//114. 二叉树展开为链表

//思路 树的右支移到左支
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil{
		return
	}
	var flat func(node *TreeNode) *TreeNode
	flat = func(node *TreeNode) *TreeNode {
		if node.Left == nil{
			if node.Right == nil{
				return node
			}
			return flat(node.Right)
		}
		if node.Right == nil{
			node.Right = node.Left
			node.Left = nil
			return flat(node.Right)
		}else{
			tmp := node.Right
			node.Right = node.Left
			node.Left = nil
			flat(node.Right).Right = tmp
			return flat(tmp)
		}
	}
	flat(root)
}
