package bpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/node"
	"github.com/stretchr/testify/assert"
)

func TestNewMetricsDecorator(t *testing.T) {
	sut := bpdf.NewMetricsDecorator(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*bpdf.MetricsDecorator", fmt.Sprintf("%T", sut))
}

func TestMetricsDecorator_AddPages(t *testing.T) {
	pg := page.New()
	docToReturn := mocks.NewDocument(t)
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := mocks.NewBPDF(t)
	inner.EXPECT().AddPages(pg)
	inner.EXPECT().Generate().Return(docToReturn, nil)
	sut := bpdf.NewMetricsDecorator(inner)

	sut.AddPages(pg)
	sut.AddPages(pg)

	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)
	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "add_page", report.TimeMetrics[1].Key)
	assert.Equal(t, 2, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddPages", 2)
}

func TestMetricsDecorator_AddRow(t *testing.T) {
	col := col.New(12)
	docToReturn := mocks.NewDocument(t)
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := mocks.NewBPDF(t)
	inner.EXPECT().AddRow(10.0, col).Return(nil)
	inner.EXPECT().Generate().Return(docToReturn, nil)
	sut := bpdf.NewMetricsDecorator(inner)

	sut.AddRow(10, col)
	sut.AddRow(10, col)

	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)
	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "add_row", report.TimeMetrics[1].Key)
	assert.Equal(t, 2, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddRow", 2)
}

func TestMetricsDecorator_AddRows(t *testing.T) {
	row := row.New(10).Add(col.New(12))
	docToReturn := mocks.NewDocument(t)
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := mocks.NewBPDF(t)
	inner.EXPECT().AddRows(row)
	inner.EXPECT().Generate().Return(docToReturn, nil)
	sut := bpdf.NewMetricsDecorator(inner)

	sut.AddRows(row)
	sut.AddRows(row)

	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)
	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "add_rows", report.TimeMetrics[1].Key)
	assert.Equal(t, 2, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddRows", 2)
}

func TestMetricsDecorator_RegisterHeader(t *testing.T) {
	row := text.NewRow(10, "text")
	inner := mocks.NewBPDF(t)
	inner.EXPECT().RegisterHeader(row).Return(nil)
	inner.EXPECT().Generate().Return(&core.Pdf{}, nil)
	sut := bpdf.NewMetricsDecorator(inner)

	err := sut.RegisterHeader(row)

	assert.Nil(t, err)
	doc, err := sut.Generate()
	assert.Nil(t, err)
	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "header", report.TimeMetrics[1].Key)
}

func TestMetricsDecorator_RegisterFooter(t *testing.T) {
	row := text.NewRow(10, "text")
	inner := mocks.NewBPDF(t)
	inner.EXPECT().RegisterFooter(row).Return(nil)
	inner.EXPECT().Generate().Return(&core.Pdf{}, nil)
	sut := bpdf.NewMetricsDecorator(inner)

	err := sut.RegisterFooter(row)

	assert.Nil(t, err)
	doc, err := sut.Generate()
	assert.Nil(t, err)
	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "footer", report.TimeMetrics[1].Key)
}

func TestMetricsDecorator_GetStructure(t *testing.T) {
	row := row.New(10).Add(col.New(12))
	docToReturn := mocks.NewDocument(t)
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := mocks.NewBPDF(t)
	inner.EXPECT().AddRows(row)
	inner.EXPECT().GetStructure().Return(&node.Node[core.Structure]{})
	inner.EXPECT().Generate().Return(docToReturn, nil)
	sut := bpdf.NewMetricsDecorator(inner)
	sut.AddRows(row)

	_ = sut.GetStructure()

	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)
	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 3, len(report.TimeMetrics))
	assert.Equal(t, "get_tree_structure", report.TimeMetrics[0].Key)
	assert.Equal(t, "generate", report.TimeMetrics[1].Key)
	assert.Equal(t, "add_rows", report.TimeMetrics[2].Key)
	assert.Equal(t, 1, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddRows", 1)
	inner.AssertNumberOfCalls(t, "GetStructure", 1)
}
