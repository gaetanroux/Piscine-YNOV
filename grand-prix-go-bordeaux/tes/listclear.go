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

// ListClear yes
func ListClear(l *List) {

	for l.Head != nil {

		l.Head.Data = nil
		l.Head = l.Head.Next
	}

}
