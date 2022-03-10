package main

import "container/heap"

//752. 打开转盘锁
//你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为'0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
//锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
//列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
//字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
//
//
//
//示例 1:
//
//输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
//输出：6
//解释：
//可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
//注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
//因为当拨动到 "0102" 时这个锁就会被锁定。
//示例 2:
//
//输入: deadends = ["8888"], target = "0009"
//输出：1
//解释：
//把最后一位反向旋转一次即可 "0000" -> "0009"。
//示例 3:
//
//输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
//输出：-1
//解释：
//无法旋转到目标数字且不被锁定。
//示例 4:
//
//输入: deadends = ["0000"], target = "8888"
//输出：-1
//
//
//提示：
//
//死亡列表 deadends 的长度范围为 [1, 500]。
//目标数字 target 不会在 deadends 之中。
//每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。

//思路 BFS 哈希表
// 还可以用 A* 算法
func openLock(deadends []string, target string) int {
	dead := map[string]bool{}
	for _, d := range deadends {
		dead[d] = true
	}
	const start = "0000"
	if dead[start] {
		return -1
	}
	if start == target {
		return 0
	}
	getWeight := func(status, target string) (ret int) {
		for i := 0; i < 4; i++ {
			dist := abs(int(status[i]) - int(target[i]))
			ret += min(dist, 10-dist)
		}
		return
	}
	getNext := func(status string) (ret []string) {
		array := []byte(status)
		for i, b := range array {
			array[i] = b - 1
			if array[i] < '0' {
				array[i] = '9'
			}
			ret = append(ret, string(array))
			array[i] = b + 1
			if array[i] > '9' {
				array[i] = '0'
			}
			ret = append(ret, string(array))
			array[i] = b
		}
		return

	}
	h := hp{{step: 0, weight: getWeight(start, target), status: start}}
	for len(h) > 0 {
		node := heap.Pop(&h).(astar)
		for _, v := range getNext(node.status) {
			if !dead[v] {
				if v == target {
					return node.step + 1
				}
				dead[v] = true
				heap.Push(&h, astar{node.step + 1, getWeight(v, target), v})
			}

		}

	}
	return -1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type astar struct {
	step, weight int
	status       string
}
type hp []astar

func (h *hp) Push(v interface{}) { *h = append(*h, v.(astar)) }
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool { return h[i].step+h[i].weight < h[j].step+h[j].weight }
func (h hp) Len() int           { return len(h) }
