package code_test

import (
	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/properties"
)

// ExampleNewBar demonstrates how to generate a barcode and add it to bpdf.
func ExampleNewBar() {
	m := bpdf.New()

	barCode := code.NewBar("123456789", properties.Barcode{Percent: 70.5})
	col := col.New(6).Add(barCode)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewBarCol demonstrates how to generate a column with a barcode and add it to bpdf.
func ExampleNewBarCol() {
	m := bpdf.New()

	barCodeCol := code.NewBarCol(6, "123456", properties.Barcode{Percent: 70.5})
	m.AddRow(10, barCodeCol)

	// generate document
}

// ExampleNewBarRow demonstrates how to generate a row with a barcode and add it to bpdf.
func ExampleNewBarRow() {
	m := bpdf.New()

	barCodeRow := code.NewBarRow(10, "123456789", properties.Barcode{Percent: 70.5})
	m.AddRows(barCodeRow)

	// generate document
}

// ExampleNewQr demonstrates how to generate a qrcode and add it to bpdf.
func ExampleNewQr() {
	m := bpdf.New()

	qrCode := code.NewQr("123456789", properties.Rect{Percent: 70.5})
	col := col.New(6).Add(qrCode)
	m.AddRow(10, col)

	// generate document
}

// ExampleNewQrCol demonstrates how to generate a column with a qrcode and add it to bpdf.
func ExampleNewQrCol() {
	m := bpdf.New()

	qrCodeCol := code.NewQrCol(12, "123456789", properties.Rect{Percent: 70.5})
	m.AddRow(10, qrCodeCol)

	// generate document
}

// ExampleNewQrRow demonstrates how to generate a row with a qrcode and add it to bpdf.
func ExampleNewQrRow() {
	m := bpdf.New()

	qrCodeRow := code.NewQrRow(10, "123456789", properties.Rect{Percent: 70.5})
	m.AddRows(qrCodeRow)

	// generate document
}
