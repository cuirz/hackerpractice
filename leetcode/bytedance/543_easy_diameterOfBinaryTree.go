package main

//543. 二叉树的直径
//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
//
//示例 :
//给定二叉树
//
//1
/// \
//2   3
/// \
//4   5
//返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。
//
//注意：两结点之间的路径长度是以它们之间边的数目表示。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	var depthCheck func(node *TreeNode, dep int) int
	depthCheck = func(node *TreeNode, dep int) int {
		if node == nil {
			return dep - 1
		}
		left := depthCheck(node.Left, dep+1)
		right := depthCheck(node.Right, dep+1)
		result = max(result, left+right-dep*2)
		return max(left, right)
	}
	depthCheck(root, 0)
	return result
}
