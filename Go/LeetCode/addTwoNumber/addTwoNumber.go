package main

// https://leetcode.com/problems/add-two-numbers/description/

func main() {
	firstTree := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	secondTree := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	addTwoNumbers(firstTree, secondTree)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode { 
	var resultNode *ListNode
	overSize := 0
	for l1 != nil || l2 != nil || overSize != 0{
		sum := 0 + overSize
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
		if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
		if sum >= 10 {
			overSize = 1
            sum -= 10
		} else {
			overSize = 0
        }
		newNode := &ListNode{sum, resultNode}
		resultNode = newNode
	}

	return reverseTree(resultNode)
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
