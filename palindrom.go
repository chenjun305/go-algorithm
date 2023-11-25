package main

/*
* 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
*
* 解法：将值复制到数组中后用双指针法
* 易于理解，时间复杂度：O(n), 空间复杂度：O(n)O(n)O(n)
 */
func isPalindrome(head *ListNode) bool {
	slice := copyToSlice(head)
	size := len(slice)
	for i := 0; i < size/2; i++ {
		if slice[i] != slice[size-1-i] {
			return false
		}
	}
	return true
}

func copyToSlice(head *ListNode) []int {
	slice := []int{}
	for {
		if head == nil {
			break
		} else {
			slice = append(slice, head.Val)
		}
		head = head.Next
	}
	return slice
}
