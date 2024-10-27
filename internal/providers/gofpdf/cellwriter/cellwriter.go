package cellwriter

import (
	"github.com/pchchv/bpdf/consts/border"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type CellWriter interface {
	SetNext(next CellWriter)
	GetNext() CellWriter
	GetName() string
	Apply(width, height float64, config *entity.Config, prop *properties.Cell)
}

type cellWriter struct {
	stylerTemplate
	defaultColor *properties.Color
}

func NewCellWriter(fpdf fpdfwrapper.Fpdf) *cellWriter {
	return &cellWriter{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "cellWriter",
		},
		defaultColor: &properties.BlackColor,
	}
}

func (c *cellWriter) Apply(width, height float64, config *entity.Config, prop *properties.Cell) {
	bd := border.None
	fill := false
	if prop != nil {
		bd = prop.BorderType
		fill = prop.BackgroundColor != nil
	}

	if config.Debug {
		bd = border.Full
	}

	c.fpdf.CellFormat(width, height, "", string(bd), 0, "C", fill, 0, "")
}
