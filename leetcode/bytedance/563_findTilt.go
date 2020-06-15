package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	result := 0
	var search func(root *TreeNode, val int) int
	search = func(root *TreeNode, val int) int {
		if root == nil {
			return val
		}
		l, r := val, val
		if root.Left != nil {
			l += root.Left.Val
		}
		if root.Right != nil {
			r += root.Right.Val
		}

		l = search(root.Left, root.Val)
		r = search(root.Right, root.Val)
		result += abs(l - r)
		return l + r
	}
	search(root, 0)
	return result

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
