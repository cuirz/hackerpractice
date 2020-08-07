package main

//23. 合并K个排序链表
//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
//示例:
//
//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6

//思路 归并 双指针
// 两两合并排序，最终归并成一个

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	} else if n > 2 {
		mid := n >> 1
		left := mergeKLists(lists[:mid])
		right := mergeKLists(lists[mid:])
		return mergeKLists([]*ListNode{left, right})
	}
	result := &ListNode{}
	head := result
	l, r := lists[0], lists[1]
	for l != nil && r != nil {
		if l.Val < r.Val {
			result.Next = l
			result = result.Next
			l = l.Next
		} else {
			result.Next = r
			result = result.Next
			r = r.Next
		}
	}
	if l != nil {
		result.Next = l
	}
	if r != nil {
		result.Next = r
	}
	return head.Next
}
