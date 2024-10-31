// Package config implements custom configuration builder.
package config

import (
	"time"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/consts/generation"
	"github.com/pchchv/bpdf/consts/orientation"
	"github.com/pchchv/bpdf/consts/pagesize"
	"github.com/pchchv/bpdf/consts/protection"
	"github.com/pchchv/bpdf/consts/provider"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

// Builder is the abstraction responsible for global customizations on the document.
type Builder interface {
	WithPageSize(size pagesize.Size) Builder
	WithDimensions(width float64, height float64) Builder
	WithLeftMargin(left float64) Builder
	WithTopMargin(top float64) Builder
	WithRightMargin(right float64) Builder
	WithBottomMargin(bottom float64) Builder
	WithConcurrentMode(chunkWorkers int) Builder
	WithSequentialMode() Builder
	WithSequentialLowMemoryMode(chunkWorkers int) Builder
	WithDebug(on bool) Builder
	WithMaxGridSize(maxGridSize int) Builder
	WithDefaultFont(font *properties.Font) Builder
	WithPageNumber(pageNumber ...properties.PageNumber) Builder
	WithProtection(protectionType protection.Protection, userPassword, ownerPassword string) Builder
	WithCompression(compression bool) Builder
	WithOrientation(orientation orientation.Orient) Builder
	WithAuthor(author string, isUTF8 bool) Builder
	WithCreator(creator string, isUTF8 bool) Builder
	WithSubject(subject string, isUTF8 bool) Builder
	WithTitle(title string, isUTF8 bool) Builder
	WithCreationDate(time time.Time) Builder
	WithCustomFonts([]*entity.CustomFont) Builder
	WithBackgroundImage([]byte, extension.Extension) Builder
	WithDisableAutoPageBreak(disabled bool) Builder
	WithKeywords(keywordsStr string, isUTF8 bool) Builder
	Build() *entity.Config
}

type CfgBuilder struct {
	providerType         provider.Provider
	dimensions           *entity.Dimensions
	margins              *entity.Margins
	chunkWorkers         int
	debug                bool
	maxGridSize          int
	defaultFont          *properties.Font
	customFonts          []*entity.CustomFont
	pageNumber           *properties.PageNumber
	protection           *entity.Protection
	compression          bool
	pageSize             *pagesize.Size
	orientation          orientation.Orient
	metadata             *entity.Metadata
	backgroundImage      *entity.Image
	disableAutoPageBreak bool
	generationMode       generation.Mode
}

// Build finalizes the customization returning the entity.Config.
func (b *CfgBuilder) Build() *entity.Config {
	if b.pageNumber != nil {
		b.pageNumber.WithFont(b.defaultFont)
	}

	return &entity.Config{
		ProviderType:         b.providerType,
		Dimensions:           b.getDimensions(),
		Margins:              b.margins,
		GenerationMode:       b.generationMode,
		ChunkWorkers:         b.chunkWorkers,
		Debug:                b.debug,
		MaxGridSize:          b.maxGridSize,
		DefaultFont:          b.defaultFont,
		PageNumber:           b.pageNumber,
		Protection:           b.protection,
		Compression:          b.compression,
		Metadata:             b.metadata,
		CustomFonts:          b.customFonts,
		BackgroundImage:      b.backgroundImage,
		DisableAutoPageBreak: b.disableAutoPageBreak,
	}
}

func (b *CfgBuilder) getDimensions() *entity.Dimensions {
	if b.dimensions != nil {
		return b.dimensions
	}

	pageSize := pagesize.A4
	if b.pageSize != nil {
		pageSize = *b.pageSize
	}

	width, height := pagesize.GetDimensions(pageSize)
	dimensions := &entity.Dimensions{
		Width:  width,
		Height: height,
	}

	if b.orientation == orientation.Horizontal && height > width {
		dimensions.Width, dimensions.Height = dimensions.Height, dimensions.Width
	}

	return dimensions
}
