package main

/*
* 给你一个链表的头节点 head ，判断链表中是否有环。
*
* 思路及算法:
* 定义两个指针，一快一慢。慢指针每次只移动一步，而快指针每次移动两步。
* 如果在移动的过程中，快指针反过来追上慢指针，就说明该链表为环形链表。否则快指针将到达链表尾部，该链表不为环形链表。
* 时间复杂度：O(N)
* 空间复杂度：O(1)
 */
func hasCycle(head *ListNode) bool {
	step1, step2 := head, head
	for {
		if step2 == nil || step2.Next == nil {
			return false
		} else {
			step1 = step1.Next
			step2 = step2.Next.Next
		}
		if step1 == step2 {
			return true
		}
	}
}
