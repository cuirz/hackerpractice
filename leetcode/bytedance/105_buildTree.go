package main

//105. 从前序与中序遍历序列构造二叉树
//根据一棵树的前序遍历与中序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	var build func(in []int, treeNode *TreeNode)
	build = func(in []int, treeNode *TreeNode) {
		if len(preorder) < 1 {
			return
		}
		node := preorder[0]
		preorder = preorder[1:]
		treeNode.Val = node
		if len(in) == 1 {
			return
		}
		for i := 0; i < len(in); i++ {
			if node == in[i] {
				// 左树
				if i == 0 {
					treeNode.Right = &TreeNode{}
					build(in[i+1:], treeNode.Right)
				} else if i == len(in)-1 {
					treeNode.Left = &TreeNode{}
					build(in[:i], treeNode.Left)
				} else {
					treeNode.Left = &TreeNode{}
					treeNode.Right = &TreeNode{}

					build(in[:i], treeNode.Left)
					build(in[i+1:], treeNode.Right)
				}
				break
			}
		}

	}
	tree := &TreeNode{}
	build(inorder, tree)
	return tree
}
