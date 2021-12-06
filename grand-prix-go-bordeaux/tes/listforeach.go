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

// ListForEach yes
func ListForEach(l *List, f func(*NodeL)) {
	curr := l.Head
	for curr != nil {
		f(curr)
		curr = curr.Next
	}
}

// Add2_node yes
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

// Subtract3_node yes
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
