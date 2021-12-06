package piscine

// TreeNode yes
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData yes
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}

	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)

	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}

// PrintTree yes
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}

}
