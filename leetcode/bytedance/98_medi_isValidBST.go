package main

import "math"

//98. 验证二叉搜索树
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例1:
//
//输入:
//2
/// \
//1   3
//输出: true
//示例2:
//
//输入:
//5
/// \
//1   4
//    / \
//   3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//    根节点的值为 5 ，但是其右子节点值为 4 。

//思路 递归
// 每个节点圈定范围判断

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var isValid func(root *TreeNode, a, b int) bool
	isValid = func(root *TreeNode, a, b int) bool {
		if root == nil {
			return true
		}
		if root.Val <= a || root.Val >= b {
			return false
		}
		return isValid(root.Left, a, root.Val) && isValid(root.Right, root.Val, b)
	}
	return isValid(root, math.MinInt32, math.MaxInt32)

}
