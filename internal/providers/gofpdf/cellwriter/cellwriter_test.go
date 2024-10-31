package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewCellCreator(t *testing.T) {
	sut := cellwriter.NewCellWriter(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.cellWriter", fmt.Sprintf("%T", sut))
}

func TestCellWriter_Apply(t *testing.T) {
	t.Run("when prop is nil without debug, should call cellformat correctly", func(t *testing.T) {
		config := &entity.Config{}
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "", 0, "C", false, 0, "")
		sut := cellwriter.NewCellWriter(fpdf)

		sut.Apply(width, height, config, nil)

		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})

	t.Run("when prop is nil with debug, should call cellformat correctly", func(t *testing.T) {
		config := &entity.Config{
			Debug: true,
		}
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "1", 0, "C", false, 0, "")
		sut := cellwriter.NewCellWriter(fpdf)

		sut.Apply(width, height, config, nil)

		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})

	t.Run("when has prop without debug, should call cellformat correctly", func(t *testing.T) {
		config := &entity.Config{}
		prop := fixture.CellProp()
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "L", 0, "C", true, 0, "")
		sut := cellwriter.NewCellWriter(fpdf)

		sut.Apply(width, height, config, &prop)

		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})

	t.Run("when has prop with debug, should call cellformat correctly", func(t *testing.T) {
		config := &entity.Config{
			Debug: true,
		}
		prop := fixture.CellProp()
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "1", 0, "C", true, 0, "")
		sut := cellwriter.NewCellWriter(fpdf)

		sut.Apply(width, height, config, &prop)

		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})
}
