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

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast := head
	slow := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.Val == value {
		l.head = l.head.Next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.Next.Val != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.length--
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
	removeNthFromEnd(node1, 2)
	myList.length--
	myList.printList()
	// result := middleNode(myList.head)
	// fmt.Println(result.Val)
}
