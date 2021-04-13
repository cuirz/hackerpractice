package main

import "math"

//783. 二叉搜索树节点最小距离
//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//输入：root = [4,2,6,1,3]
//输出：1
//输入：root = [1,0,48,null,null,12,49]
//输出：1
//提示：
//
//树中节点数目在范围 [2, 100] 内
//0 <= Node.val <= 10^5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	result, pre := math.MaxInt32, -1
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		search(node.Left)
		if pre != -1 && node.Val-pre < result {
			result = node.Val - pre
		}
		pre = node.Val
		search(node.Right)
	}
	search(root)
	return result
}
