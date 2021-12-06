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

// ListLast yes
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
