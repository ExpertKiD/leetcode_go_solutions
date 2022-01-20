package main

// Link To Problem: https://leetcode.com/problems/add-two-numbers/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var carry = 0
	var sum = 0

	temp1 := l1
	temp2 := l2

	var l3 *ListNode = &ListNode{
		Val:  0,
		Next: nil,
	}

	var temp3 = l3

	for temp1 != nil || temp2 != nil {
		var val1 = 0
		var val2 = 0

		if temp1 == nil {
			val1 = 0
		} else {
			val1 = temp1.Val
		}

		if temp2 == nil {
			val2 = 0
		} else {
			val2 = temp2.Val
		}

		sum = val1 + val2 + carry

		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}

		temp3.Val = sum

		if temp1.Next == nil && temp2.Next == nil {
			temp3.Next = nil
		} else {
			temp3.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
		}

		temp3 = temp3.Next

		temp1 = temp1.Next
		temp2 = temp2.Next
	}

	return l3
}
