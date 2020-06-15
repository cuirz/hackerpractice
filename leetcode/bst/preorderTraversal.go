package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := new([]int)
	search(root, result)
	return *result
}

func search(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	search(root.Left, result)
	search(root.Right, result)
}

//输入: [1,null,2,3]
