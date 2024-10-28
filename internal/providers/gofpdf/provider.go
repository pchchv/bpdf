package gofpdf

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type provider struct {
	fpdf       fpdfwrapper.Fpdf
	font       core.Font
	text       core.Text
	code       core.Code
	image      core.Image
	line       core.Line
	cache      cache.Cache
	cellWriter cellwriter.CellWriter
	cfg        *entity.Config
}

func (g *provider) AddText(text string, cell *entity.Cell, prop *properties.Text) {
	g.text.Add(text, cell, prop)
}

func (g *provider) GetLinesQuantity(text string, textProp *properties.Text, colWidth float64) int {
	return g.text.GetLinesQuantity(text, textProp, colWidth)
}
