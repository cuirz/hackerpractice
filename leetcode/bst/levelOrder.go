package bst

func levelOrder(root *TreeNode) [][]int {
	result := new([][]int)
	depth := 0
	loop(root, result, depth)
	return *result
}

func loop(root *TreeNode, result *[][]int, depth int) {
	if root == nil {
		return
	}
	if len(*result) <= depth {
		*result = append(*result, []int{})
	}
	(*result)[depth] = append((*result)[depth], root.Val)
	loop(root.Left, result, depth+1)
	loop(root.Right, result, depth+1)
}

//二叉树：[3,9,20,null,null,15,7],
//[
//	[3],
//	[9,20],
//	[15,7]
//]
