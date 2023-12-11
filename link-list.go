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
