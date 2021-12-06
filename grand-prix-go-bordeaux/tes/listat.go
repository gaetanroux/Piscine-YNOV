package piscine

// NodeL yes
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// ListAt yes
func ListAt(l *NodeL, pos int) *NodeL {
	current := 0
	for l != nil {
		if current == pos {
			return l
		}
		current++
		l = l.Next
	}
	return nil
}
