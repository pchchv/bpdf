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

// ToMap converts Config to a map[string]interface{} .
func (c *Config) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if c.ProviderType != "" {
		m["config_provider_type"] = c.ProviderType
	}

	if c.Dimensions != nil {
		m = c.Dimensions.AppendMap("bpdf", m)
	}

	if c.Margins != nil {
		m = c.Margins.AppendMap(m)
	}

	if c.DefaultFont != nil {
		m = c.DefaultFont.AppendMap(m)
	}

	m["generation_mode"] = c.GenerationMode
	m["chunk_workers"] = c.ChunkWorkers
	if c.Debug {
		m["config_debug"] = c.Debug
	}

	if c.MaxGridSize != 0 {
		m["config_max_grid_sum"] = c.MaxGridSize
	}

	if c.PageNumber != nil {
		m = c.PageNumber.AppendMap(m)
	}

	if c.Protection != nil {
		m = c.Protection.AppendMap(m)
	}

	if c.Compression {
		m["config_compression"] = c.Compression
	}

	if c.Metadata != nil {
		m = c.Metadata.AppendMap(m)
	}

	if c.BackgroundImage != nil {
		m = c.BackgroundImage.AppendMap(m)
	}

	if c.DisableAutoPageBreak {
		m["config_disable_auto_page_break"] = c.DisableAutoPageBreak
	}

	return m
}
