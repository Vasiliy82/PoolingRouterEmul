package model

type ListNode struct {
	Value any
	next  *ListNode
}

type CircularLinkedList struct {
	head *ListNode
	tail *ListNode
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{head: nil, tail: nil}
}

func (cll *CircularLinkedList) IsEmpty() bool {
	return cll.head == nil
}

func (cll *CircularLinkedList) AddNode(value any) {
	newNode := &ListNode{Value: value, next: nil}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = cll.head
	} else {
		cll.tail.next = newNode
		cll.tail = newNode
		newNode.next = cll.head
	}
}
