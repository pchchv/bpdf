package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/components/list"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/signature"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/config"
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/metrics"
	"github.com/pchchv/bpdf/properties"
)

var (
	dummyText  = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."
	background = &properties.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
)

type Object struct {
	Key   string
	Value string
}

func (o Object) GetHeader() core.Row {
	return row.New(10).Add(
		text.NewCol(4, "Key", properties.Text{Style: fontstyle.Bold}),
		text.NewCol(8, "Bytes", properties.Text{Style: fontstyle.Bold}),
	)
}

func (o Object) GetContent(i int) core.Row {
	r := row.New(5).Add(
		text.NewCol(4, o.Key),
		text.NewCol(8, o.Value),
	)
	if i%2 == 0 {
		r.WithStyle(&properties.Cell{
			BackgroundColor: background,
		})
	}

	return r
}

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

func getObjects(max int) (objects []Object) {
	for i := 0; i < max; i++ {
		objects = append(objects, Object{
			Key:   fmt.Sprintf("Key: %d", i),
			Value: fmt.Sprintf("Bytes: %d", i),
		})
	}

	return
}

func run() *metrics.Time {
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

	objects := getObjects(1158)
	rows, err := list.Build[Object](objects)
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(rows...)
	for i := 0; i < 1158; i++ {
		m.AddRows(buildCodesRow()...)
		m.AddRows(buildImagesRow()...)
		m.AddRows(buildTextsRow()...)
	}

	m.AddRows(
		text.NewRow(15, "Dummy Data", properties.Text{Size: 12, Top: 5, Align: align.Center}),
	)

	for i := 0; i < 1158; i++ {
		m.AddRows(text.NewRow(20, dummyText+dummyText+dummyText+dummyText+dummyText))
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	return document.GetReport().TimeMetrics[0].Avg
}

func main() {
	var content string
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		content += fmt.Sprintf("%f", run().Value) + "\n"
	}

	if err := os.WriteFile("docs/assets/text/benchmark.txt", []byte(content), os.ModePerm); err != nil {
		log.Fatal(err.Error())
	}
}
