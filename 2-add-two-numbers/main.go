package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	i := addTwoNumbers(&ListNode{}, &ListNode{})
	fmt.Println(i)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{}
}
