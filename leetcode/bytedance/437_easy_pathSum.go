package main

//437. 路径总和 III
//给定一个二叉树，它的每个结点都存放着一个整数值。
//
//找出路径和等于给定数值的路径总数。
//
//路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

//思路  哈希表+深度优先

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {

	dic := make(map[int]int)
	dic[0] = 1
	var dfs func(node *TreeNode, value int) int
	dfs = func(node *TreeNode, value int) int {
		if node == nil {
			return 0
		}
		value += node.Val
		count := dic[value-sum]
		dic[value] += 1
		lCount := dfs(node.Left, value)
		rCount := dfs(node.Right, value)
		dic[value] -= 1
		return count + lCount + rCount
	}
	return dfs(root, 0)
}
