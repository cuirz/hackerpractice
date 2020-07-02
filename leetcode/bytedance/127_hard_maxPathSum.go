package main

//124. 二叉树中的最大路径和
//给定一个非空二叉树，返回其最大路径和。
//
//本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//示例 1:
//
//输入: [1,2,3]
//
//1
/// \
//2   3
//
//输出: 6
//示例 2:
//
//输入: [-10,9,20,null,null,15,7]
//
//  -10
// / \
// 9 20
//  / \
// 15  7
//
//输出: 42

//[5,4,8,11,null,13,4,7,2,null,null,null,1]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	result := root.Val
	var backtrace func(root *TreeNode) int
	backtrace = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := backtrace(root.Left)
		right := backtrace(root.Right)
		res := max(root.Val, max(root.Val+left, root.Val+right))
		result = max(result, max(res, root.Val+left+right))
		return res
	}
	backtrace(root)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
