package entity

import (
	"github.com/pchchv/bpdf/consts/generation"
	"github.com/pchchv/bpdf/consts/provider"
	"github.com/pchchv/bpdf/properties"
)

// Config is the configuration of a bpdf instance.
type Config struct {
	ProviderType         provider.Provider
	Dimensions           *Dimensions
	Margins              *Margins
	DefaultFont          *properties.Font
	CustomFonts          []*CustomFont
	GenerationMode       generation.Mode
	ChunkWorkers         int
	Debug                bool
	MaxGridSize          int
	PageNumber           *properties.PageNumber
	Protection           *Protection
	Compression          bool
	Metadata             *Metadata
	BackgroundImage      *Image
	DisableAutoPageBreak bool
}
