package main

import (
	"math"
	"strconv"
	"strings"
)

//449. 序列化和反序列化二叉搜索树
//序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
//
//设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
//
//编码的字符串应尽可能紧凑。
//
//
//示例 1：
//
//输入：root = [2,1,3]
//输出：[2,1,3]
//示例 2：
//
//输入：root = []
//输出：[]
//
//
//提示：
//
//树中节点数范围是 [0, 10^4]
//0 <= Node.val <= 10^4
//题目数据 保证 输入的树是一棵二叉搜索树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var array []string
	var loop func(node *TreeNode)
	loop = func(node *TreeNode) {
		if node == nil {
			return
		}
		loop(node.Left)
		loop(node.Right)
		array = append(array, strconv.Itoa(node.Val))
	}
	loop(root)

	return strings.Join(array, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	array := strings.Split(data, ",")
	var construct func(low, high int) *TreeNode
	construct = func(low, high int) *TreeNode {
		if len(array) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(array[len(array)-1])
		if val < low || val > high {
			return nil
		}
		array = array[:len(array)-1]
		// 先构造root，在构造Right，最后构造Left
		return &TreeNode{Val: val, Right: construct(val, high), Left: construct(low, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
