package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	var search func(root *TreeNode) *TreeNode
	search = func(root *TreeNode) *TreeNode {
		if root != nil {
			search(root.Right)
			sum += root.Val
			root.Val = sum
			search(root.Left)
		}
		return root
	}
	return search(root)
}
