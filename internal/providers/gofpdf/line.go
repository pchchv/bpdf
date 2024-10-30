package gofpdf

import (
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/consts/orientation"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type line struct {
	pdf              fpdfwrapper.Fpdf
	defaultColor     *properties.Color
	defaultThickness float64
}

func NewLine(pdf fpdfwrapper.Fpdf) *line {
	return &line{
		pdf:              pdf,
		defaultColor:     &properties.BlackColor,
		defaultThickness: linestyle.DefaultLineThickness,
	}
}

func (l *line) Add(cell *entity.Cell, prop *properties.Line) {
	if prop.Orientation == orientation.Vertical {
		l.renderVertical(cell, prop)
	} else {
		l.renderHorizontal(cell, prop)
	}
}

func (l *line) renderVertical(cell *entity.Cell, prop *properties.Line) {
	size := cell.Height * (prop.SizePercent / 100.0)
	position := cell.Width * (prop.OffsetPercent / 100.0)
	space := (cell.Height - size) / 2.0
	left, top, _, _ := l.pdf.GetMargins()
	if prop.Color != nil {
		l.pdf.SetDrawColor(prop.Color.Red, prop.Color.Green, prop.Color.Blue)
	}

	l.pdf.SetLineWidth(prop.Thickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 1}, 0)
	}

	l.pdf.Line(left+cell.X+position, top+cell.Y+space, left+cell.X+position, top+cell.Y+cell.Height-space)

	if prop.Color != nil {
		l.pdf.SetDrawColor(l.defaultColor.Red, l.defaultColor.Green, l.defaultColor.Blue)
	}

	l.pdf.SetLineWidth(l.defaultThickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 0}, 0)
	}
}

func (l *line) renderHorizontal(cell *entity.Cell, prop *properties.Line) {
	size := cell.Width * (prop.SizePercent / 100.0)
	position := cell.Height * (prop.OffsetPercent / 100.0)
	space := (cell.Width - size) / 2.0
	left, top, _, _ := l.pdf.GetMargins()

	if prop.Color != nil {
		l.pdf.SetDrawColor(prop.Color.Red, prop.Color.Green, prop.Color.Blue)
	}

	l.pdf.SetLineWidth(prop.Thickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 1}, 0)
	}

	l.pdf.Line(left+cell.X+space, top+cell.Y+position, left+cell.X+cell.Width-space, top+cell.Y+position)

	if prop.Color != nil {
		l.pdf.SetDrawColor(l.defaultColor.Red, l.defaultColor.Green, l.defaultColor.Blue)
	}

	l.pdf.SetLineWidth(l.defaultThickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 0}, 0)
	}
}
