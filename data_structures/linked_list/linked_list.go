package main

import "fmt"

type LLNode struct {
	Data int
	Next *LLNode
}

type LinkedList struct {
	Head *LLNode
}

func print(head *LLNode) {
	if head == nil {
		return
	} else {
		fmt.Print(head.Data)
		print(head.Next)
	}
}

func insert(head, node *LLNode) *LLNode {
	if head == nil {
		return &LLNode{node.Data, nil}
	}

	return nil
}
