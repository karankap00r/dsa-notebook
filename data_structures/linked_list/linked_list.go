package main

import "fmt"

type LLNode struct {
	Data int
	Next *LLNode
}

type LinkedList struct {
	Head *LLNode
}

//Print will print out the linked list content
func (list LinkedList) Print() {
	print(list.Head)
}

func print(head *LLNode) {
	if head == nil {
		return
	} else {
		fmt.Print(head.Data)
		print(head.Next)
	}
}

func insertAtStart(head, node *LLNode) *LLNode {
	if head == nil {
		return &LLNode{node.Data, nil}
	}

	node.Next = head
	return node
}
