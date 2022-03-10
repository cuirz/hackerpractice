package main

import "math"

//111. 二叉树的最小深度
//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明:叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树[3,9,20,null,null,15,7],

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	depth := math.MaxInt32
	if root.Left != nil {
		depth = min(minDepth(root.Left)+1, depth)
	}
	if root.Right != nil {
		depth = min(minDepth(root.Right)+1, depth)
	}
	return depth
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
