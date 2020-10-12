package main

//107. 二叉树的层次遍历 II
//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其自底向上的层次遍历为：
//
//[
//[15,7],
//[9,20],
//[3]
//]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var search func(node *TreeNode, dep int)
	search = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		if dep == len(result) {
			result = append(result, []int{node.Val})
		} else {
			result[dep] = append(result[dep], node.Val)
		}
		search(node.Left, dep+1)
		search(node.Right, dep+1)

	}
	search(root,0)
	n := len(result)
	for i := 0; i < n>>1; i++ {
		result[i], result[n-i-1] = result[n-i-1], result[i]
	}
	return result
}
