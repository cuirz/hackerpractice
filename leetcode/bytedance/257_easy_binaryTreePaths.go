package main

import (
	"strconv"
	"strings"
)

//257. 二叉树的所有路径
//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	queue := make([]string, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			result = append(result, strings.Join(queue, "->"))
			return
		}
		if node.Left != nil {
			queue = append(queue, strconv.Itoa(node.Left.Val))
			dfs(node.Left)
			queue = queue[:len(queue)-1]
		}
		if node.Right != nil {
			queue = append(queue, strconv.Itoa(node.Right.Val))
			dfs(node.Right)
			queue = queue[:len(queue)-1]
		}
	}
	if root != nil {
		queue = append(queue, strconv.Itoa(root.Val))
		dfs(root)
	}
	return result
}
