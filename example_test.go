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

// ExampleBPDFGetStruct demonstrates how to get bpdf component tree
func ExampleBPDF_GetStructure() {
	m := bpdf.New()

	m.AddRow(40, text.NewCol(12, "text"))

	m.GetStructure()

	// Do things and generate
}

// ExampleBPDF_RegisterHeader demonstrates how to register a header to me added in every new page.
// An error is returned if the area occupied by the header is greater than the page area.
func ExampleBPDF_RegisterHeader() {
	m := bpdf.New()

	err := m.RegisterHeader(
		code.NewBarRow(12, "barcode"),
		text.NewRow(12, "text"))
	if err != nil {
		panic(err)
	}

	// Do things and generate
}

// ExampleBPDF_RegisterFooter demonstrates how to register a footer to me added in every new page.
// An error is returned if the area occupied by the footer is greater than the page area.
func ExampleBPDF_RegisterFooter() {
	m := bpdf.New()

	err := m.RegisterFooter(
		code.NewBarRow(12, "barcode"),
		text.NewRow(12, "text"))
	if err != nil {
		panic(err)
	}

	// Do things and generate
}

// ExampleBPDF_FitlnCurrentPage demonstrate how to check if the new line fits on the current page.
func ExampleBPDF_FitlnCurrentPage() {
	m := bpdf.New()

	m.FitlnCurrentPage(12)

	// Do things and generate
}

// ExampleBPDF_FitlnCurrentPage demonstrate how to check if the new line fits on the current page.
func ExampleBPDF_GetCurrentConfig() {
	m := bpdf.New()

	m.GetCurrentConfig()

	// Do things and generate
}
