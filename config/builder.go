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

// WithLeftMargin customize margin.
func (b *CfgBuilder) WithLeftMargin(left float64) Builder {
	if left < pagesize.MinLeftMargin {
		return b
	}

	b.margins.Left = left
	return b
}

// WithTopMargin customize margin.
func (b *CfgBuilder) WithTopMargin(top float64) Builder {
	if top < pagesize.MinTopMargin {
		return b
	}

	b.margins.Top = top
	return b
}

// WithRightMargin customize margin.
func (b *CfgBuilder) WithRightMargin(right float64) Builder {
	if right < pagesize.MinRightMargin {
		return b
	}

	b.margins.Right = right
	return b
}

// WithBottomMargin customize margin.
func (b *CfgBuilder) WithBottomMargin(bottom float64) Builder {
	if bottom < pagesize.MinBottomMargin {
		return b
	}

	b.margins.Bottom = bottom
	return b
}

// WithBackgroundImage defines the background image that will be applied in every page.
func (b *CfgBuilder) WithBackgroundImage(bytes []byte, ext extension.Extension) Builder {
	b.backgroundImage = &entity.Image{
		Bytes:     bytes,
		Extension: ext,
	}

	return b
}

// WithAuthor defines the author name metadata.
func (b *CfgBuilder) WithAuthor(author string, isUTF8 bool) Builder {
	if author == "" {
		return b
	}

	if b.metadata == nil {
		b.metadata = &entity.Metadata{}
	}

	b.metadata.Author = &entity.Utf8Text{
		Text: author,
		UTF8: isUTF8,
	}

	return b
}

// WithCreator defines the creator name metadata.
func (b *CfgBuilder) WithCreator(creator string, isUTF8 bool) Builder {
	if creator == "" {
		return b
	}

	if b.metadata == nil {
		b.metadata = &entity.Metadata{}
	}

	b.metadata.Creator = &entity.Utf8Text{
		Text: creator,
		UTF8: isUTF8,
	}

	return b
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
