package node

// Node is the base of Tree construction.
type Node[T any] struct {
	id       int
	data     T
	previous *Node[T]
	nexts    []*Node[T]
}
