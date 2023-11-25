package main

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// Recursion递归解法
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists(list1, list2.Next)
			return list2
		}
	} else if list1 == nil {
		return list2
	} else {
		return list1
	}
}

// iteration迭代解法
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	list := new(ListNode)
	head := list
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return list.Next
}
