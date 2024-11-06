package image_test

import (
	"os"

	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/consts/extension"
)

// ExampleNewFromBytes demonstrates how to create
// an image component reading bytes.
func ExampleNewFromBytes() {
	m := bpdf.New()

	bytes, _ := os.ReadFile("image.png")

	image := image.NewFromBytes(bytes, extension.Png)
	col := col.New(12).Add(image)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewFromBytesCol demonstrates how to create
// an image component wrapped into a column reading bytes.
func ExampleNewFromBytesCol() {
	m := bpdf.New()

	bytes, _ := os.ReadFile("image.png")

	imageCol := image.NewFromBytesCol(12, bytes, extension.Png)
	m.AddRow(10, imageCol)

	// generate document
}

// ExampleNewFromBytesRow demonstrates how to create
// an image component wrapped into a row reading bytes.
func ExampleNewFromBytesRow() {
	m := bpdf.New()

	bytes, _ := os.ReadFile("image.png")

	imageRow := image.NewFromBytesRow(10, bytes, extension.Png)
	m.AddRows(imageRow)

	// generate document
}
