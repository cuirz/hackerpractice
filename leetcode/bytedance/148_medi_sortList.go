package main

//148. 排序链表
//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//示例 1:
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2:
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5

//思路 用数组排序，然后重组链表
//思路 使用O(1)空间复杂度， 归并排序
// 分割链表，然后再合并链表
// 分割使用 快慢指针找到中心点进行分割
// 合并使用 双指针排序
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)

	result := &ListNode{}
	top := result
	for left != nil && right != nil {
		if left.Val < right.Val {
			top.Next, left = left, left.Next
		} else {
			top.Next, right = right, right.Next
		}
		top = top.Next
	}
	if left == nil {
		top.Next = right
	} else {
		top.Next = left
	}
	return result.Next

}
