package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	for i := 0; i < 3; i++ {
		fmt.Println(l1.Val, *l1.Next)
		fmt.Println(l2.Val, *l2.Next)
	}

	result := &ListNode{}
	current := result

	return current
}

func main() {
	listNode1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	listNode2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	result := addTwoNumbers(&listNode1, &listNode2)

	fmt.Println("result = ", result)
}
