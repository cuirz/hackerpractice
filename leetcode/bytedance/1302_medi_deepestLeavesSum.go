package main

//1302. 层数最深叶子节点的和
//给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
//
//
//
//示例 1：
//
//
//
//输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//输出：15
//示例 2：
//
//输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//输出：19
//
//
//提示：
//
//树中节点数目在范围 [1, 104]之间。
//1 <= Node.val <= 100

//思路 深度优先

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	var depth []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		if d == len(depth) {
			depth = append(depth, node.Val)
		} else {
			depth[d] += node.Val
		}
		dfs(node.Left, d+1)
		dfs(node.Right, d+1)
	}
	dfs(root, 0)
	return depth[len(depth)-1]
}
