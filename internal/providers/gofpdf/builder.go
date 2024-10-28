package gofpdf

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
)

// Dependencies is the dependencies provider for gofpdf
type Dependencies struct {
	Fpdf       fpdfwrapper.Fpdf
	Font       core.Font
	Text       core.Text
	Code       core.Code
	Image      core.Image
	Line       core.Line
	Cache      cache.Cache
	CellWriter cellwriter.CellWriter
	Cfg        *entity.Config
}

// Builder is the dependencies builder for gofpdf
type Builder interface {
	Build(cfg *entity.Config, cache cache.Cache) *Dependencies
}

type builder struct{}
