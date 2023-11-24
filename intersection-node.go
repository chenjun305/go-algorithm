package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 给你两个单链表的头节点 headA 和 headB ，
* 请你找出并返回两个单链表相交的起始节点。
* 如果两个链表不存在相交节点，返回 null 。
*
* 思路和算法
* 判断两个链表是否相交，可以使用哈希集合存储链表节点。
* 首先遍历链表 headA，并将链表 headA 中的每个节点加入哈希集合中。
* 然后遍历链表 headB，对于遍历到的每个节点，判断该节点是否在哈希集合中
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	for node := headA; node != nil; node = node.Next {
		m[node] = true
	}
	for node := headB; node != nil; node = node.Next {
		if m[node] {
			return node
		}
	}
	return nil
}

func main() {
	n := &ListNode{3, &ListNode{4, nil}}
	headA := &ListNode{1, n}
	headB := &ListNode{2, n}
	node := getIntersectionNode(headA, headB)
	fmt.Println("intersection node: ", node.Val)
}
