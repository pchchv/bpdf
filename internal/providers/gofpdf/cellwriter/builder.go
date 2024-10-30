package cellwriter

import "github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"

type CellWriterBuilder struct{}

func NewBuilder() *CellWriterBuilder {
	return &CellWriterBuilder{}
}

func (c *CellWriterBuilder) Build(fpdf fpdfwrapper.Fpdf) CellWriter {
	cellCreator := NewCellWriter(fpdf)
	borderColorStyle := NewBorderColorStyler(fpdf)
	borderLineStyler := NewBorderLineStyler(fpdf)
	borderThicknessStyler := NewBorderThicknessStyler(fpdf)
	fillColorStyler := NewFillColorStyler(fpdf)

	borderThicknessStyler.SetNext(borderLineStyler)
	borderLineStyler.SetNext(borderColorStyle)
	borderColorStyle.SetNext(fillColorStyler)
	fillColorStyler.SetNext(cellCreator)

	return borderThicknessStyler
}
