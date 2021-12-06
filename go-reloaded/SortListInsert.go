package main

import (
	"fmt"

	student ".."
)

// SortListInsert OK
func SortListInsert(l *NodeI, data-ref int) *NodeI {
	node := &NodeI{}
	node.Data = data-ref
	node.Next = nil

	if l == nil || l.Data > node.Data {
		node.Next = l
		return node
	} else {
		a := l
		for a.Next != nil && a.Next.Data <= node.Data {
			a = a.Next
		}

		node.Next = a.Next
		a.Next = node
	}
	return l
}

// PrintList OK
func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

// PrintList OK
func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *student.NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = student.SortListInsert(link, -2)
	link = student.SortListInsert(link, 2)
	PrintList(link)
}
