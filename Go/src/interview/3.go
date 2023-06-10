/*
反转链表
*/
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func reverseList(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func main() {
	head := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: nil}}}

	fmt.Println("Original List: ")
	for curr := head; curr != nil; curr = curr.next {
		fmt.Printf("%d ", curr.data)
	}

	head = reverseList(head)

	fmt.Println("\nReversed List: ")
	for curr := head; curr != nil; curr = curr.next {
		fmt.Printf("%d ", curr.data)
	}
}
