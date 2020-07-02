package main

import "strconv"

//1028. 从先序遍历还原二叉树
//我们从二叉树的根节点 root 开始进行深度优先搜索。
//
//在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。
//
//如果节点只有一个子节点，那么保证该子节点为左子节点。
//
//给出遍历输出 S，还原树并返回其根节点 root。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	array := make([][]int, 0)
	index := 0
	num := ""
	dep := 0
	isFirst := true
	for i, _ := range S {
		switch S[i] {
		case '-':
			if isFirst {
				isFirst = false
				v, _ := strconv.Atoi(num)
				array = append(array, []int{dep, v})
				dep = 0
			}
			index = i + 1
			dep++
		default:
			num = S[index : i+1]
			isFirst = true
		}
	}
	v, _ := strconv.Atoi(num)
	array = append(array, []int{dep, v})

	root := &TreeNode{}
	root.Val = array[0][1]
	index = 1

	var dfs func(node *TreeNode, d int)
	dfs = func(node *TreeNode, d int) {
		if index >= len(array) {
			return
		}
		if array[index][0] > d {
			left := &TreeNode{}
			left.Val = array[index][1]
			node.Left = left
			index++
			dfs(node.Left, d+1)
		}
		if index >= len(array) {
			return
		}
		if array[index][0] > d {
			right := &TreeNode{}
			right.Val = array[index][1]
			node.Right = right
			index++
			dfs(node.Right, d+1)
		}
		return
	}
	dfs(root, 0)
	return root
}

func main() {
	recoverFromPreorder("1-2--3---4-5--6---7")
}
