package gofpdf

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"path/filepath"
// 	"strings"

// 	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"

// 	"github.com/johnfercher/maroto/v2/internal/cache"
// 	"github.com/johnfercher/maroto/v2/internal/merror"
// 	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"
// 	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/fpdfwrapper"
// 	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
// 	"github.com/johnfercher/maroto/v2/pkg/core"
// 	"github.com/johnfercher/maroto/v2/pkg/core/entity"
// 	"github.com/johnfercher/maroto/v2/pkg/props"
// )

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

func (g *provider) GetLinesQuantity(text string, textProp *properties.Text, colWidth float64) int {
	return g.text.GetLinesQuantity(text, textProp, colWidth)
}

func (g *provider) GetFontHeight(prop *properties.Font) float64 {
	return g.font.GetHeight(prop.Family, prop.Style, prop.Size)
}

func (g *provider) AddText(text string, cell *entity.Cell, prop *properties.Text) {
	g.text.Add(text, cell, prop)
}

func (g *provider) AddLine(cell *entity.Cell, prop *properties.Line) {
	g.line.Add(cell, prop)
}
