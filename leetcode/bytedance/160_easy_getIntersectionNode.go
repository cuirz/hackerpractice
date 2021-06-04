package main

//160. 相交链表

//思路 哈希表
//思路 双指针循环遍历 他们之间有交点时，迟早会在交点上碰撞。 A 指针循环到完 接到 B头开始，B指针循环完接到 A头开始
//思路 双指针
//从一个头开始循环链表倒转链表后，另一个头开始循环到末尾，然后末尾回来时倒转链表到末尾
//
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var pre *ListNode
	next := headA

	loop := func() int {
		count := 0
		for next != nil {
			count++
			tmp := next
			next = next.Next
			tmp.Next = pre
			pre = tmp
		}
		return count
	}

	x := loop()
	tailA := pre
	pre, next = nil, headB
	y := loop()
	tailB := pre

	pre, next = nil, tailA
	z := loop()

	//判断是否有交点
	if tailB != headA {
		//没有交点===
		pre, next = nil, tailB
		loop()
		return nil
	}
	mid := x - (x+z-y+1)>>1
	next = headA
	index := 0
	for next != nil {
		if mid == index {
			return next
		}
		next = next.Next
		index++
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil{
		return nil
	}
	tailA,tailB := headA,headB
	for tailA != tailB{
		if tailA == nil{
			tailA = headB
		}else{
			tailA = tailA.Next
		}
		if tailB == nil{
			tailB = headA
		}else{
			tailB = tailB.Next
		}
	}
	return tailA
}
