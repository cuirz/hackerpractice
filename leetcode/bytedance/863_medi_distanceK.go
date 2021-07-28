package main

//863. 二叉树中所有距离为 K 的结点
//给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。
//
//返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。
//
// 
//
//示例 1：
//
//输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
//输出：[7,4,1]
//解释：
//所求结点为与目标结点（值为 5）距离为 2 的结点，
//值分别为 7，4，以及 1
//
//
//
//注意，输入的 "root" 和 "target" 实际上是树上的结点。
//上面的输入仅仅是对这些对象进行了序列化描述。
// 
//
//提示：
//
//给定的树是非空的。
//树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
//目标结点 target 是树上的结点。
//0 <= K <= 1000

//思路 建图 哈希表

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := make(map[int]*TreeNode)
	var result []int
	var findParent func(node *TreeNode)
	var findK func(node, from *TreeNode, depth int)
	findParent = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParent(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParent(node.Right)
		}
	}

	findK = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k {
			result = append(result, node.Val)
			return
		}
		if node.Left != from {
			findK(node.Left, node, depth+1)
		}
		if node.Right != from {
			findK(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findK(parents[node.Val], node, depth+1)
		}
	}

	findParent(root)
	findK(target, nil, 0)
	return result

}
