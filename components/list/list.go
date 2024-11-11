// Package list implements creation of lists (old tablelist).
package list

import "github.com/pchchv/bpdf/core"

// Listable is the main abstraction of a listable item in a TableList.
// A collection of objects that implements this
// interface may be added in a list.
type Listable interface {
	GetHeader() core.Row
	GetContent(i int) core.Row
}
