package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/signature"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/config"
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/properties"
)

var dummyText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."

func buildCodesRow() []core.Row {
	return []core.Row{
		row.New(20).Add(
			text.NewCol(4, "Barcode:", properties.Text{Size: 15, Top: 6, Align: align.Center}),
			code.NewBarCol(8, "barcode", properties.Barcode{Center: true, Percent: 70}),
		),
		row.New(20).Add(
			text.NewCol(4, "QrCode:", properties.Text{Size: 15, Top: 6, Align: align.Center}),
			code.NewQrCol(8, "qrcode", properties.Rect{Center: true, Percent: 70}),
		),
		row.New(20).Add(
			text.NewCol(4, "MatrixCode:", properties.Text{Size: 15, Top: 6, Align: align.Center}),
			code.NewMatrixCol(8, "matrixcode", properties.Rect{Center: true, Percent: 70}),
		),
	}
}

func buildImagesRow() []core.Row {
	bytes, err := os.ReadFile("docs/assets/images/frontpage.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	return []core.Row{
		row.New(20).Add(
			text.NewCol(4, "Image From File:", properties.Text{Size: 15, Top: 6, Align: align.Center}),
			image.NewFromFileCol(8, "docs/assets/images/biplane.jpg", properties.Rect{Center: true, Percent: 90}),
		),
		row.New(20).Add(
			text.NewCol(4, "Image From Bytes:", properties.Text{Size: 15, Top: 6, Align: align.Center}),
			image.NewFromBytesCol(8, bytes, extension.Png, properties.Rect{Center: true, Percent: 90}),
		),
	}
}

func buildHeader() []core.Row {
	r1 := row.New(30).Add(
		col.New(12).Add(
			text.New("Config V2", properties.Text{
				Top:   5,
				Size:  15,
				Align: align.Center,
			}),
			text.New("Grid system, fast generation, embedded metrics and testable.", properties.Text{
				Top:   13,
				Size:  13,
				Align: align.Center,
			}),
		),
	)

	return []core.Row{r1}
}

func buildFooter() []core.Row {
	return []core.Row{
		row.New(10).Add(
			text.NewCol(2, "GitHub: https://github.com/pchchv/bpdf/"),
		),
	}
}

func buildTextsRow() []core.Row {
	colText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."
	return []core.Row{
		row.New(20).Add(
			text.NewCol(4, "Text:", properties.Text{Size: 15, Top: 6, Align: align.Center}),
			text.NewCol(8, colText, properties.Text{Size: 12, Top: 5, Align: align.Center}),
		),
		row.New(40).Add(
			text.NewCol(4, "Signature:", properties.Text{Size: 15, Top: 17, Align: align.Center}),
			signature.NewCol(8, "Name", properties.Signature{FontSize: 10}),
		),
	}
}

func main() {
	var err error
	cfg := config.NewBuilder().
		WithPageNumber().
		Build()
	mrt := bpdf.New(cfg)
	m := bpdf.NewMetricsDecorator(mrt)
	if err = m.RegisterHeader(buildHeader()...); err != nil {
		log.Fatal(err.Error())
	}

	if err = m.RegisterFooter(buildFooter()...); err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(
		text.NewRow(20, "Main features", properties.Text{Size: 15, Top: 6.5}),
	)
	m.AddRows(buildCodesRow()...)
	m.AddRows(buildImagesRow()...)
	m.AddRows(buildTextsRow()...)
	m.AddRows(
		text.NewRow(15, "Dummy Data", properties.Text{Size: 12, Top: 5, Align: align.Center}),
	)

	for i := 0; i < 50; i++ {
		m.AddRows(text.NewRow(20, dummyText+dummyText+dummyText+dummyText+dummyText))
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = document.Save("docs/assets/pdf/v2.pdf"); err != nil {
		log.Fatal(err.Error())
	}

	if err = document.GetReport().Save("docs/assets/text/v2.txt"); err != nil {
		log.Fatal(err.Error())
	}
}
