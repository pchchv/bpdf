package col_test

import (
	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/signature"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/consts/border"
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/properties"
)

// ExampleNew demonstrates how to create a Col instance.
func ExampleNew() {
	// size is an optional parameters, if not provided, bpdf
	// will apply the maximum size, even if custom size is applied.
	size := 12
	col := col.New(size)

	row := row.New(10).Add(col)

	m := bpdf.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleCol_Add demonstrates how to add components to Col.
func ExampleCol_Add() {
	col := col.New()

	text := text.New("text content")
	qrCode := code.NewQr("qrcode")
	signature := signature.New("signature label")

	col.Add(text, qrCode, signature)

	row := row.New(10).Add(col)

	m := bpdf.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}

// ExampleCol_WithStyle demonstrates how to add style to Col.
func ExampleCol_WithStyle() {
	col := col.New()
	col.WithStyle(&properties.Cell{
		BackgroundColor: &properties.Color{
			Red:   10,
			Green: 100,
			Blue:  150,
		},
		BorderColor: &properties.Color{
			Red:   55,
			Green: 10,
			Blue:  60,
		},
		BorderType:      border.Full,
		BorderThickness: 0.1,
		LineStyle:       linestyle.Dashed,
	})

	row := row.New(10).Add(col)
	m := bpdf.New()
	m.AddRows(row)

	// Do things and generate
	_, _ = m.Generate()
}
