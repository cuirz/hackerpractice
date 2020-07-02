package main

import (
	"strconv"
	"strings"
)

//297. 二叉树的序列化与反序列化
//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
//
//请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
//
//示例:
//
//你可以将以下二叉树：
//
//1
/// \
//2   3
/// \
//4   5
//
//序列化为 "[1,2,3,null,null,4,5]"
//提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
//
//说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	var data []string
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			data = append(data, "null")
		} else {
			data = append(data, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return "[" + strings.Join(data, ",") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	size := len(data)
	sub := data[1 : size-1]
	if len(sub) < 1 {
		return nil
	}
	arrays := strings.Split(sub, ",")
	n := len(arrays)

	queue := make([]*TreeNode, 0)
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(arrays[0])
	queue = append(queue, root)

	index := 1
	for len(queue) > 0 {
		if index >= n {
			break
		}
		node := queue[0]
		queue = queue[1:]
		if arrays[index] != "null" {
			node.Left = &TreeNode{}
			node.Left.Val, _ = strconv.Atoi(arrays[index])
			queue = append(queue, node.Left)
		}
		index++
		if index >= n {
			break
		}
		if arrays[index] != "null" {
			node.Right = &TreeNode{}
			node.Right.Val, _ = strconv.Atoi(arrays[index])
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func main() {
	obj := Constructor()
	ans := obj.deserialize("[1,2,3,null,null,4,5,6,7,null,8,null,null,null,null,9,10]")
	println(ans)
	println(obj.serialize(ans))

}
