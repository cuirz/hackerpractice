package main

//699. 掉落的方块
//在无限长的数轴（即 x 轴）上，我们根据给定的顺序放置对应的正方形方块。
//
//第 i 个掉落的方块（positions[i] = (left, side_length)）是正方形，其中left 表示该方块最左边的点位置(positions[i][0])，side_length 表示该方块的边长(positions[i][1])。
//
//每个方块的底部边缘平行于数轴（即 x 轴），并且从一个比目前所有的落地方块更高的高度掉落而下。在上一个方块结束掉落，并保持静止后，才开始掉落新方块。
//
//方块的底边具有非常大的粘性，并将保持固定在它们所接触的任何长度表面上（无论是数轴还是其他方块）。邻接掉落的边不会过早地粘合在一起，因为只有底边才具有粘性。
//
//
//
//返回一个堆叠高度列表ans 。每一个堆叠高度ans[i]表示在通过positions[0], positions[1], ..., positions[i]表示的方块掉落结束后，目前所有已经落稳的方块堆叠的最高高度。
//
//
//
//
//
//示例 1:
//
//输入: [[1, 2], [2, 3], [6, 1]]
//输出: [2, 5, 5]
//解释:
//
//第一个方块 positions[0] = [1, 2] 掉落：
//_aa
//_aa
//-------
//方块最大高度为 2 。
//
//第二个方块 positions[1] = [2, 3] 掉落：
//__aaa
//__aaa
//__aaa
//_aa__
//_aa__
//--------------
//方块最大高度为5。
//大的方块保持在较小的方块的顶部，不论它的重心在哪里，因为方块的底部边缘有非常大的粘性。
//
//第三个方块 positions[1] = [6, 1] 掉落：
//__aaa
//__aaa
//__aaa
//_aa
//_aa___a
//--------------
//方块最大高度为5。
//
//因此，我们返回结果[2, 5, 5]。
//
//示例 2:
//
//输入: [[100, 100], [200, 100]]
//输出: [100, 100]
//解释: 相邻的方块不会过早地卡住，只有它们的底部边缘才能粘在表面上。
//
//
//注意:
//
//1 <= positions.length <= 1000.
//1 <= positions[i][0] <= 10^8.
//1 <= positions[i][1] <= 10^6.

//思路  线段树
//参考 https://www.youtube.com/watch?v=xuoQdt5pHj0&t=622s
// 二叉树
// 动态开点   懒标记
func fallingSquares(positions [][]int) []int {
	var result []int
	tree := &SegmentTree{}
	root := &Node{L: 0, R: 1e9}
	//root := &Node{L: 0, R: 7}
	for _, pos := range positions {
		l, r, h := pos[0], pos[0]+pos[1]-1, pos[1]
		h += tree.Query(root, l, r)
		tree.Update(root, l, r, h)
		result = append(result, root.Height)
	}
	return result
}

type Node struct {
	Height      int
	Add         int
	L, R        int
	Left, Right *Node
}
type SegmentTree struct {
	Root *Node
}

func (seg *SegmentTree) Update(node *Node, l, r, height int) {
	if node.L >= l && node.R <= r {
		node.Add = height
		node.Height = height
		return
	}
	seg.PushDown(node)
	mid := (node.L + node.R) >> 1
	if l <= mid {
		seg.Update(node.Left, l, r, height)
	}
	if r > mid {
		seg.Update(node.Right, l, r, height)
	}
	seg.PushUp(node)
}
func (seg *SegmentTree) Query(node *Node, l, r int) int {
	if l <= node.L && r >= node.R {
		return node.Height
	}
	seg.PushDown(node)
	mid, ret := (node.L+node.R)>>1, 0
	if l <= mid {
		ret = seg.Query(node.Left, l, r)
	}
	if r > mid {
		ret = max(ret, seg.Query(node.Right, l, r))
	}
	return ret
}

func (seg *SegmentTree) PushDown(node *Node) {
	if node.Left == nil {
		node.Left = &Node{L: node.L, R: (node.L + node.R) >> 1}
	}
	if node.Right == nil {
		node.Right = &Node{L: (node.L+node.R)>>1 + 1, R: node.R}
	}
	if node.Add == 0 {
		return
	}
	node.Left.Add = node.Add
	node.Right.Add = node.Add
	node.Left.Height = node.Add
	node.Right.Height = node.Add
	node.Add = 0
}
func (seg *SegmentTree) PushUp(node *Node) {
	node.Height = max(node.Left.Height, node.Right.Height)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}})
}
