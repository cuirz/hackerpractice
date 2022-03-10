package main

//106. 从中序与后序遍历序列构造二叉树
//根据一棵树的中序遍历与后序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//中序遍历 inorder =[9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) < 1 || len(inorder) < 1 {
		return nil
	}
	var build func(in, post []int, treeNode *TreeNode)
	build = func(in, post []int, treeNode *TreeNode) {
		node := post[len(post)-1]
		post = post[:len(post)-1]
		treeNode.Val = node
		for i := 0; i < len(in); i++ {
			if node == in[i] {
				if i < len(in)-1 {
					treeNode.Right = &TreeNode{}
					build(in[i+1:], post[i:], treeNode.Right)
				}
				if i > 0 {
					treeNode.Left = &TreeNode{}
					build(in[:i], post[:i], treeNode.Left)
				}
				break

			}
		}

	}
	tree := &TreeNode{}
	build(inorder, postorder, tree)
	return tree
}

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	//[9,3,15,20,7]
	//[9,15,7,20,3]
}
