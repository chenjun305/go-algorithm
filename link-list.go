package main

// import "fmt"

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */
// 将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
// 如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	arr := make([]int, k)
	i := 0
	newHead := new(ListNode)
	h := newHead
	for head != nil {
		arr[i] = head.Val
		i++
		if i == k {
			for i > 0 {
				i--
				h.Next = &ListNode{Val: arr[i]}
				h = h.Next
			}
		}
		head = head.Next
	}
	if i > 0 {
		for j := 0; j < i; j++ {
			h.Next = &ListNode{Val: arr[j]}
			h = h.Next
		}
	}
	return newHead.Next
}

// 递归解法
func reverseKGroup2(head *ListNode, k int) *ListNode {
	// write code here
	tail := head
	i := k
	for i > 0 {
		if tail == nil {
			return head
		}
		tail = tail.Next
		i--
	}
	var pre *ListNode = nil
	cur := head
	for cur != tail {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	head.Next = reverseKGroup(tail, k)
	return pre
}

/*
* 将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)，空间复杂度 O(1)。
* 数据范围: 0<m≤n≤size
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	i := 0
	var pre, left, mid *ListNode
	h := head
	for h != nil {
		i++
		if i < m {
			pre = h
			h = h.Next
		} else if i == m {
			left = pre
			mid = h
			pre = h
			h = h.Next
		} else if i > m && i < n {
			tmp := h.Next
			h.Next = pre
			pre = h
			h = tmp
		} else if i == n {
			tmp := h.Next
			h.Next = pre
			mid.Next = tmp
			if left == nil {
				return h
			} else {
				left.Next = h
			}
		} else {
			break
		}
	}
	return head
}

/*
* 输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。
* 解法：双向奔赴？有点意思的题目，A+B = B+A，找最后相遇的点就是共同走的路
 */
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	cur1, cur2 := pHead1, pHead2
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = pHead2
		} else {
			cur1 = cur1.Next
		}
		if cur2 == nil {
			cur2 = pHead1
		} else {
			cur2 = cur2.Next
		}
	}

	return cur1
}

/*
* 删除给出链表中的重复元素（链表中元素从小到大有序）, 使链表中的所有元素都只出现一次
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	h := head
	var pre *ListNode
	for h != nil {
		if pre != nil && h.Val == pre.Val {
			pre.Next = h.Next
		} else {
			pre = h
		}
		h = h.Next
	}
	return head
}

/*
* 给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
* 解法： 递归，只处理前面的元素
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	} else {
		if head.Val != head.Next.Val {
			head.Next = deleteDuplicates(head.Next)
			return head
		} else {
			for head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			return deleteDuplicates(head.Next)
		}
	}
}
