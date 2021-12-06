package piscine

// NodeL 1
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List 2
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushBack yes
func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
