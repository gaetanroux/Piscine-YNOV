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

// ListRemoveIf yes
func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	var previous *NodeL = nil
	curr := l.Head
	for curr != nil {
		if curr.Data != data_ref {
			previous = curr
			curr = curr.Next
			continue
		}
		curr = curr.Next
		if previous == nil {
			l.Head = curr
		} else {
			previous.Next = curr
		}
	}
}
