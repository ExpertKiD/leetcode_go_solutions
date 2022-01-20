package main

import (
	"fmt"
)

// Link To Problem: https://leetcode.com/problems/add-two-numbers/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Arr []int
	var l2Arr []int

	var temp = l1

	// convert l1 and l2 to array
	for temp != nil {
		l1Arr = append(l1Arr, temp.Val)
		temp = temp.Next
	}

	temp = l2

	for temp != nil {
		l2Arr = append(l2Arr, temp.Val)
		temp = temp.Next
	}

	fmt.Printf("%v\n", l1Arr)
	fmt.Printf("%v\n", l2Arr)

	num1, num2 := 0, 0

	// convert arrays now to integer
	for i := len(l1Arr) - 1; i >= 0; i-- {
		num1 = num1*10 + l1Arr[i]
	}

	for i := len(l2Arr) - 1; i >= 0; i-- {
		num2 = num2*10 + l2Arr[i]
	}

	// get the sum
	var sum = num1 + num2

	print(sum, "\n")

	// convert sum to array, then to listnode and return
	var tempNumber = sum
	var sumArr []int

	for tempNumber != 0 {
		sumArr = append(sumArr, tempNumber%10)
		tempNumber /= 10
	}

	fmt.Printf("%v\n", sumArr)

	var l3 *ListNode = &ListNode{
		Val:  0,
		Next: nil,
	}

	tempNode := l3

	for i := 0; i < len(sumArr); i++ {
		tempNode.Val = sumArr[i]
		tempNode.Next = &ListNode{
			Val:  0,
			Next: nil,
		}

		if i == len(sumArr)-1 {
			tempNode.Next = nil
		}

		tempNode = tempNode.Next
	}

	return l3
}

func main() {

	var l1 *ListNode = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	var l2 *ListNode = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := addTwoNumbers(l1,
		l2)

	println(l3)
}
