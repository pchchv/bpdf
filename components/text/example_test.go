package text_test

import (
	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/text"
)

// ExampleNew demonstrates how to create a text component.
func ExampleNew() {
	m := bpdf.New()

	text := text.New("text")
	col := col.New(12).Add(text)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewCol demonstrates how to create a text component wrapped into a column.
func ExampleNewCol() {
	m := bpdf.New()

	textCol := text.NewCol(12, "text")
	m.AddRow(10, textCol)

	// generate document
}

// ExampleNewRow demonstrates how to create a text component wrapped into a row.
func ExampleNewRow() {
	m := bpdf.New()

	textRow := text.NewRow(10, "text")
	m.AddRows(textRow)

	// generate document
}
