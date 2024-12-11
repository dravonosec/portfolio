package main

import "fmt"

// https://leetcode.com/problems/merge-two-sorted-lists/description/


type ListNode struct { 
	Val int
	Next *ListNode
}

func main() { 
	var tail1 *ListNode = &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	var tail2 *ListNode = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	merged := mergeTwoLists(tail1, tail2)
	for merged != nil { 
        fmt.Print(merged.Val, " -> ")
        merged = merged.Next
    }
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var merged *ListNode
	for list1 != nil || list2 != nil { 
		if list1 == nil {
			newNode := &ListNode{list2.Val, merged}
			merged = newNode
			list2 = list2.Next
		} else if list2 == nil {
			newNode := &ListNode{list1.Val, merged}
            merged = newNode
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
                newNode := &ListNode{list1.Val, merged}
                merged = newNode
                list1 = list1.Next
            } else {
                newNode := &ListNode{list2.Val, merged}
                merged = newNode
                list2 = list2.Next
            }
		}
	}

	return reverseTree(merged)
}

func reverseTree(tree *ListNode) *ListNode {
	var resultTree *ListNode
	for tree != nil {
		newNode := &ListNode{tree.Val, resultTree}
        resultTree = newNode
        tree = tree.Next
	}

	return resultTree
}