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

// CompStr yes
func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	curr := l.Head
	for curr != nil {
		if comp(ref, curr.Data) {
			return &curr.Data
		}
		curr = curr.Next
	}
	a := &NodeL{Data: nil}
	return &a.Data
}
