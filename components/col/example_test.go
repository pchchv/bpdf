package col_test

import (
	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/row"
)

// ExampleNew demonstrates how to create a Col instance.
func ExampleNew() {
	// size is an optional parameters, if not provided, maroto
	// will apply the maximum size, even if custom size is applied.
	size := 12
	col := col.New(size)

	row := row.New(10).Add(col)

	m := bpdf.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}
