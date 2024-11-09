package bpdf_test

import (
	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/config"
)

// ExampleNew demonstrates how to create a bpdf instance.
func ExampleNew() {
	// optional
	b := config.NewBuilder()
	cfg := b.Build()

	m := bpdf.New(cfg) // cfg is an optional

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleNewMetricsDecorator demonstrates how to create a bpdf metrics decorator instance.
func ExampleNewMetricsDecorator() {
	// optional
	b := config.NewBuilder()
	cfg := b.Build()

	mrt := bpdf.New(cfg)               // cfg is an optional
	m := bpdf.NewMetricsDecorator(mrt) // decorator of bpdf

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleBPDF_AddPages demonstrates how to add a new page in bpdf.
func ExampleBPDF_AddPages() {
	m := bpdf.New()

	p := page.New()
	p.Add(code.NewBarRow(10, "barcode"))

	m.AddPages(p)

	// Do things and generate
}

// ExampleBPDF_AddRows demonstrates how to add new rows in bpdf.
func ExampleBPDF_AddRows() {
	m := bpdf.New()

	m.AddRows(
		code.NewBarRow(12, "barcode"),
		text.NewRow(12, "text"),
	)

	// Do things and generate
}

// ExampleBPDF_AddRow demonstrates how to add a new row in bpdf.
func ExampleBPDF_AddRow() {
	m := bpdf.New()

	m.AddRow(10, text.NewCol(12, "text"))

	// Do things and generate
}
