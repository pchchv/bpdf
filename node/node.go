package node

import "fmt"

// Node is the base of Tree construction.
type Node[T any] struct {
	id       int
	data     T
	previous *Node[T]
	nexts    []*Node[T]
}

// New creates a new node.
func New[T any](data T) *Node[T] {
	return &Node[T]{
		data: data,
	}
}

// WithID retrieves data from node.
func (n *Node[T]) WithID(id int) *Node[T] {
	n.id = id
	return n
}

// GetData retrieves data from node.
func (n *Node[T]) GetData() T {
	return n.data
}

// GetID retrieves id from node.
func (n *Node[T]) GetID() int {
	return n.id
}

// GetPrevious retrieves the next nodes.
func (n *Node[T]) GetPrevious() *Node[T] {
	return n.previous
}

// GetNexts retrieves the next nodes.
func (n *Node[T]) GetNexts() []*Node[T] {
	return n.nexts
}

// GetStructure retrieves the node structure.
func (n *Node[T]) GetStructure() (structure []string) {
	var current string
	if n.previous == nil {
		current = fmt.Sprintf("(NULL) -> (%d)", n.id)
	} else {
		current = fmt.Sprintf("(%d) -> (%d)", n.previous.id, n.id)
	}

	if n.nexts != nil {
		current += ", "
	}

	structure = append(structure, current)
	for _, next := range n.nexts {
		innerStructure := next.GetStructure()
		structure = append(structure, innerStructure...)
	}

	return
}

// IsRoot retrieves info if node is root.
func (n *Node[T]) IsRoot() bool {
	return n.previous == nil
}

// IsLeaf retrieves info if node is leaf.
func (n *Node[T]) IsLeaf() bool {
	return len(n.nexts) == 0
}

// Backtrack retrieves a path from node to root.
func (n *Node[T]) Backtrack() (nodes []*Node[T]) {
	for current := n; current != nil; {
		nodes = append(nodes, current)
		current = current.previous
	}
	return
}

// AddNext add node to current node.
func (n *Node[T]) AddNext(node *Node[T]) {
	node.previous = n
	n.nexts = append(n.nexts, node)
}
