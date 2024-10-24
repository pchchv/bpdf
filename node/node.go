package node

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
