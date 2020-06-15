package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result *[][]int
	search(root, result, 0)
	return *result
}

func search(root *TreeNode, res *[][]int, depth int) {
	if root == nil {
		return
	}
	if len(*res) <= depth {
		*res = append(*res, []int{})
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	search(root.Left, res, depth+1)
	search(root.Right, res, depth+1)
}
