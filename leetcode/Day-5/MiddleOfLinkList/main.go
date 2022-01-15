package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type linkedList struct {
	head   *ListNode
	length int
}

func (l *linkedList) prepend(n *ListNode) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}

func (l linkedList) printList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Val)
		toPrint = toPrint.Next
		l.length--
	}
	fmt.Print("\n")
}

func middleNode(head *ListNode) *ListNode {
	// var A *ListNode = head
	// var B *ListNode = head
	A := head
	B := head
	for B != nil {
		B = B.Next
		if B == nil {
			return A
		}
		A = A.Next
		B = B.Next
	}
	return A
}

func main() {
	myList := linkedList{}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	// node6 := &ListNode{Val: 6}
	// myList.prepend(node6)
	myList.prepend(node5)
	myList.prepend(node4)
	myList.prepend(node3)
	myList.prepend(node2)
	myList.prepend(node1)
	fmt.Println(myList)
	myList.printList()
	result := middleNode(myList.head)
	fmt.Println(result.Val)
}
