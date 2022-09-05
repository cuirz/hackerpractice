package main

//652. 寻找重复的子树
//给定一棵二叉树 root，返回所有重复的子树。
//
//对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
//如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
//
//
//示例 1：
//
//
//
//输入：root = [1,2,3,4,null,2,4,null,null,4]
//输出：[[2,4],[4]]
//示例 2：
//
//
//
//输入：root = [2,1,1]
//输出：[[1]]
//示例 3：
//
//
//
//输入：root = [2,2,2,3,null,3,null]
//输出：[[2,3],[3]]
//
//
//提示：
//
//树中的结点数在[1,10^4]范围内。
//-200 <= Node.val <= 200

//方法：子树序列化

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	type pair struct {
		node *TreeNode
		idx  int
	}
	seen := make(map[[3]int]pair)
	repeat := make(map[*TreeNode]struct{})
	var dfs func(treeNode *TreeNode) int
	idx := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
		if p, ok := seen[tri]; ok {
			repeat[p.node] = struct{}{}
			return p.idx
		}
		idx++
		seen[tri] = pair{node, idx}
		return idx
	}
	dfs(root)
	var result []*TreeNode
	for node := range repeat {
		result = append(result, node)
	}
	return result
}
