package main

import "fmt"

type ListNode struct {
	Value int
	Next *ListNode
}

func (l *ListNode) printElements() {
	head := false
	current := l

	for !head {
		if current.Next == nil {
			fmt.Printf("%d", current.Value)
			head = true
		} else {
			fmt.Printf("%d -> ", current.Value)
			current = current.Next
		}
	}
}

func (l *ListNode) addElement(val int) *ListNode {
	newElement := &ListNode{Value: val, Next: nil}
	l.Next = newElement
	return newElement
}

func main() {
	list := &ListNode{Value: 0}
	newElement := list.addElement(2)
	newElement = newElement.addElement(3)

	list.printElements()
}
