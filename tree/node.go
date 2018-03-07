package tree

import "fmt"

// Node in a tree
type Node struct {
	Value       int
	Left, Right *Node
}

// Print the node value
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// SetValue set value to node
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil" +
			"node. Ignored.")
		return
	}
	node.Value = value
}

// CreateNode from a value
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
