package main

//234. 回文链表
//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

//思路 链表存于数组，然后对数组进行回文检查
//思路 O(1)空间，定位 n/2位置 反转后半段链表，前和尾比较判断回文

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	n := 0
	next := head
	for next != nil {
		next = next.Next
		n++
	}
	next = head
	pre := head

	for i := 0; i < n; i++ {
		if i > n/2 {
			tmp := next
			next = next.Next
			tmp.Next = pre
			pre = tmp
		} else {
			pre = next
			next = next.Next
		}
	}
	next = pre
	for i := 0; i < n/2; i++ {
		if head.Val != next.Val {
			return false
		}
		head = head.Next
		next = next.Next
	}
	return true
}
