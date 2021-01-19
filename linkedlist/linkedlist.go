package linkedlist

import "fmt"

// LinkedList is the linked list implementation in go
type LinkedList struct {
	Head *Node
	Tail *Node
}

// Node is the basic element consisted of linkedList
type Node struct {
	Next *Node
	Data interface{}
}

// NewLinkedList returns a new empty linkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Insert d into the position, pushing the rest of list starting from pos one step forward
func (l *LinkedList) Insert(pos int, d interface{}) error {
	len := l.Length()
	if pos > len || pos < 0 {
		return fmt.Errorf("Position %v is out of range. Length of this list is %v", pos, len)
	}
	if pos == 0 {
		if l.Head == nil {
			newNode := Node{
				Data: d,
			}
			l.Head = &newNode
			l.Tail = &newNode
			return nil
		}
		newNode := Node{
			Next: l.Head,
			Data: d,
		}
		l.Head = &newNode
		return nil
	}
	curNode := l.Head
	for i := 1; i < pos; i++ {
		curNode = curNode.Next
	}
	if pos == len {
		newNode := Node{
			Data: d,
		}
		l.Tail = &newNode
		curNode.Next = &newNode
		return nil
	}
	postNewNodePtr := curNode.Next
	newNode := Node{
		Next: postNewNodePtr,
		Data: d,
	}
	curNode.Next = &newNode
	return nil
}

// Delete deletes the node at position pos
func (l *LinkedList) Delete(pos int) error {
	len := l.Length()
	if pos >= len || pos < 0 {
		return fmt.Errorf("Position %v is out of range. Length of this list is %v", pos, len)
	}
	if pos == 0 {
		if len == 1 {
			l.Head = nil
			l.Tail = nil
			return nil
		}
		newHeadPtr := l.Head.Next
		l.Head = newHeadPtr
		return nil
	}
	curNodePtr := l.Head
	for i := 1; i < pos; i++ {
		curNodePtr = curNodePtr.Next
	}
	if pos == len-1 {
		curNodePtr.Next = nil
		l.Tail = curNodePtr
		return nil
	}
	newCurNodePtr := curNodePtr.Next.Next
	curNodePtr.Next = newCurNodePtr
	return nil
}

// Append adds a new data after the tail
func (l *LinkedList) Append(d interface{}) {
	newNode := Node{
		Data: d,
	}
	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
		return
	}
	if l.Head == l.Tail {
		l.Head.Next = &newNode
		l.Tail = &newNode
		return
	}
	l.Tail.Next = &newNode
	l.Tail = &newNode
}

// Length returns the length of the linkedList
func (l *LinkedList) Length() int {
	if l.Head == nil {
		return 0
	}
	if l.Head == l.Tail {
		return 1
	}
	len := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		len++
	}
	return len
}
