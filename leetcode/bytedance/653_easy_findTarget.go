package main

//653. 两数之和 IV - 输入 BST
//给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
//
//
//
//示例 1：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 9
//输出: true
//示例 2：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 28
//输出: false
//
//
//提示:
//
//二叉树的节点个数的范围是[1, 10^4].
//-10^4<= Node.val <= 10^4
//root为二叉搜索树
//-10^5<= k <= 10^5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	dic := make(map[int]struct{})
	var dfs func(treeNode *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := dic[k-node.Val]; ok {
			return true
		}
		dic[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}

func findTarget2(root *TreeNode, k int) bool {
	array := make([]int, 0)
	var search func(*TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		search(node.Left)
		array = append(array, node.Val)
		search(node.Right)
	}
	search(root)
	l, r := 0, len(array)-1
	for l < r {
		sum := array[l] + array[r]
		if sum == k {
			return true
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return false
}
