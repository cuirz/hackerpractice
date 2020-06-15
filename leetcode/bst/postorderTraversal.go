package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := new([]int)
	search(root, result)
	return *result
}

func search(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	search(root.Left, result)
	search(root.Right, result)
	*result = append(*result, root.Val)
}

//输入: [1,null,2,3]
