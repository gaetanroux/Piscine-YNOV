package piscine

// NodeL yes
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List yes
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListSize yes
func ListSize(l *List) int {
	listSize := 0
	for l.Head != nil {
		listSize++
		l.Head = l.Head.Next
	}
	return listSize
}
