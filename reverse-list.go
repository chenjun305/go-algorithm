package main

/*
* 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 */

// recursion递归解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		revert := reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return revert
	}
}

// iteration迭代解法
// 在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。
// 在更改引用之前，还需要存储后一个节点。最后返回新的头引用。
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
