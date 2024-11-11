package main

import (
	"fmt"
	"os"

	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/signature"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/properties"
)

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
