// Package list implements creation of lists (old tablelist).
package list

import (
	"errors"

	"github.com/pchchv/bpdf/core"
)

// Listable is the main abstraction of a listable item in a TableList.
// A collection of objects that implements this
// interface may be added in a list.
type Listable interface {
	GetHeader() core.Row
	GetContent(i int) core.Row
}

// Build is responsible to receive a collection of objects that
// implements Listable and build the rows of TableList.
// This method should be used in case of a collection of values.
func Build[T Listable](arr []T) (rows []core.Row, err error) {
	if len(arr) == 0 {
		return nil, errors.New("empty array")
	}

	rows = append(rows, arr[0].GetHeader())
	for i, element := range arr {
		rows = append(rows, element.GetContent(i))
	}

	return rows, nil
}
