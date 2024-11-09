package bpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/mocks"
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
